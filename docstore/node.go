package docstore

import (
	"net/http"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net"
	"time"
	"bytes"
	"github.com/v2pro/quokka/kvstore"
	"strings"
	"github.com/v2pro/plz/countlog"
	"github.com/v2pro/quokka/docstore/runtime"
)

// node.go is the process running endpoint
// it will forward to other nodes, or itself (command_processor)
// node.go will delegate the state to topo.go/cluster.go

func StartNode(addr string) {
	var mux = &http.ServeMux{}
	mux.HandleFunc("/docstore/ping", ping)
	mux.HandleFunc("/docstore/exec", exec)
	mux.HandleFunc("/docstore/", exec)
	// will promote servers to master if needed
	go http.ListenAndServe(addr, mux)
	if !isAlive(addr) {
		countlog.Error("event!node.failed to listen http", "addr", addr)
		return
	}
	err := publishNode(addr)
	if err != nil {
		countlog.Error("event!node.failed to publish", "addr", addr)
		return
	}
}

func isAlive(addr string) bool {
	for i := 0; i < 3; i++ {
		_, err := http.Get("http://" + addr + "/docstore/ping")
		if err != nil {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		return true
	}
	return false
}

func ping(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Write([]byte("pong"))
}

func exec(respWriter http.ResponseWriter, req *http.Request) {
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
	partitionId := kvstore.HashToPartition(cmd.EntityId)
	var commandResp []byte
	target, isLocal := chooseCommandTarget(partitionId, &cmd, req)
	if isLocal {
		commandResp = commandProcessors[partitionId].delegatedExec(&cmd, time.Second)
	} else {
		commandResp = forwardCommand(target, &cmd)
	}
	// TODO: clear master if master hint is different from our setting
	respWriter.Write(commandResp)
}

func chooseCommandTarget(partitionId uint64, cmd *command, req *http.Request) (string, bool) {
	if cmd.IsPromoting {
		return "", false
	}
	ctx := req.Context()
	localAddr, _ := ctx.Value(http.LocalAddrContextKey).(*net.TCPAddr)
	partition := topo.get(partitionId)
	cmd.IsPromoting = partition.isPromoting
	master := partition.master
	// TODO: when master is empty, forward to any other node
	// TODO: prevent forward cycle, fail after 3 hops
	if localAddr == nil {
		return master, false
	}
	if master != localAddr.String() {
		return master, false
	}
	return "", false
}

func forwardCommand(target string, cmd *command) []byte {
	req, err := runtime.Json.Marshal(cmd)
	if err != nil {
		return replyError(err)
	}
	url := "http://" + target + "/docstore/exec"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	if err != nil {
		return replyError(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return replyError(err)
	}
	return body
}
