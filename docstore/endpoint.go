package docstore

import (
	"net/http"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net"
	"errors"
	"time"
	"fmt"
	"bytes"
	"github.com/v2pro/quokka/kvstore"
	"strings"
	"github.com/v2pro/plz/countlog"
	"github.com/v2pro/quokka/docstore/runtime"
)

func StartNode(addrStr string) {
	addr, err := net.ResolveTCPAddr("tcp", addrStr)
	if err != nil {
		countlog.Error("event!endpoint.failed to resolve tcp addr", "err", err, "addrStr", addrStr)
		return
	}
	theNode = newNode(addr)
	var mux = &http.ServeMux{}
	mux.HandleFunc("/docstore/ping", theNode.ping)
	mux.HandleFunc("/docstore/exec", theNode.exec)
	mux.HandleFunc("/docstore/", theNode.exec)
	// will promote servers to master if needed
	go http.ListenAndServe(addrStr, mux)
	if !theNode.isAlive() {
		countlog.Error("event!endpoint.failed to start node", "addr", addr)
		return
	}
	// tell other servers about myself
	joinCluster(theNode.status)
	// continue update myself to let other servers know I am alive
	// will step down from master if there are more servers available
	go theNode.RebalanceInBackground()
}

func (node *node) isAlive() bool {
	for i := 0; i < 3; i++ {
		_, err := http.Get("http://" + node.status.Addr.String() + "/docstore/ping")
		if err != nil {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		return true
	}
	return false
}

func (node *node) ping(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Write([]byte("pong"))
}

func (node *node) exec(respWriter http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respWriter.Write(replyError(err))
		return
	}
	var cmd command
	err = jsoniter.Unmarshal(body, &cmd)
	if err != nil {
		respWriter.Write(replyError(err))
		return
	}
	if !strings.HasSuffix(req.URL.Path, "/exec") {
		path := strings.Trim(req.URL.Path, "/")
		segments := strings.Split(path, "/")
		if len(segments) >= 2 {
			cmd.CommandType = segments[len(segments)-1]
			cmd.EntityType = segments[len(segments)-2]
		}
	}
	partition := kvstore.HashToPartition(cmd.EntityId)
	if !cmd.IsPromoting {
		ctx := req.Context()
		localAddr, _ := ctx.Value(http.LocalAddrContextKey).(*net.TCPAddr)
		master, isPromoting, err := node.getMaster(partition)
		if err != nil {
			respWriter.Write(replyError(err))
			return
		}
		cmd.IsPromoting = isPromoting
		if localAddr == nil {
			respWriter.Write(replyError(errors.New("missing local addr")))
			return
		}
		if !master.IP.Equal(localAddr.IP) || master.Port != localAddr.Port {
			respWriter.Write(forwardToMaster(master, &cmd))
			return
		}
	}
	commandResp := commandProcessors[partition].delegatedExec(&cmd, time.Second)
	respWriter.Write(commandResp)
	return
}

func forwardToMaster(master *net.TCPAddr, cmd *command) []byte {
	req, err := runtime.Json.Marshal(cmd)
	if err != nil {
		return replyError(err)
	}
	url := fmt.Sprintf("http://%s:%d/docstore/exec", master.IP.String(), master.Port)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	if err != nil {
		return replyError(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return replyError(err)
	}
	iter := jsoniter.ConfigFastest.BorrowIterator(body)
	defer jsoniter.ConfigFastest.ReturnIterator(iter)
	stream := jsoniter.ConfigFastest.BorrowStream(nil)
	defer jsoniter.ConfigFastest.ReturnStream(stream)
	stream.WriteObjectStart()
	iter.ReadObjectCB(func(iter *jsoniter.Iterator, field string) bool {
		stream.WriteObjectField(field)
		stream.Write(iter.SkipAndReturnBytes())
		stream.WriteMore()
		return true
	})
	stream.WriteObjectField("hint_master")
	stream.WriteVal(master)
	stream.WriteObjectEnd()
	return stream.Buffer()
}
