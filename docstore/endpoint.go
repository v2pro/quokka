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
)

func StartNode(addr string) {
	node := newNode(addr)
	var mux = &http.ServeMux{}
	mux.HandleFunc("/docstore/ping", node.ping)
	mux.HandleFunc("/docstore/exec", node.exec)
	mux.HandleFunc("/docstore/", node.exec)
	// will promote servers to master if needed
	go http.ListenAndServe(addr, mux)
	if !node.isAlive() {
		countlog.Error("event!endpoint.failed to start node", "addr", addr)
		return
	}
	// tell other servers about myself
	joinCluster(node.status)
	// continue update myself to let other servers know I am alive
	// will step down from master if there are more servers available
	go node.RebalanceInBackground()
}

func (node *node) isAlive() bool {
	for i := 0; i < 3; i++ {
		_, err := http.Get("http://" + node.status.Addr + "/docstore/ping")
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
			cmd.CommandType = segments[len(segments) - 1]
			cmd.EntityType = segments[len(segments) - 2]
		}
	}
	ctx := req.Context()
	localAddr, _ := ctx.Value(http.LocalAddrContextKey).(*net.TCPAddr)
	partition := kvstore.HashToPartition(cmd.EntityId)
	master, err := node.getMaster(partition)
	if err != nil {
		respWriter.Write(replyError(err))
		return
	}
	if localAddr == nil {
		respWriter.Write(replyError(errors.New("missing local addr")))
		return
	}
	if !master.IP.Equal(localAddr.IP) || master.Port != localAddr.Port {
		respWriter.Write(forwardToMaster(master, body))
		return
	}
	commandResp := commandProcessors[partition].delegatedExec(&cmd, time.Second)
	respWriter.Write(commandResp)
	return
}

func forwardToMaster(master *net.TCPAddr, execReq []byte) []byte {
	url := fmt.Sprintf("http://%s:%d/docstore/exec", master.IP.String(), master.Port)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(execReq))
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