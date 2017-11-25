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
)

var Mux = &http.ServeMux{}

func init() {
	Mux.HandleFunc("/docstore/exec", exec)
}

func exec(respWriter http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respWriter.Write(replyError(err))
		return
	}
	var reqObj command
	err = jsoniter.Unmarshal(body, &reqObj)
	if err != nil {
		respWriter.Write(replyError(err))
		return
	}
	ctx := req.Context()
	localAddr, _ := ctx.Value(http.LocalAddrContextKey).(*net.TCPAddr)
	partition := HashToPartition(reqObj.EntityId)
	reqObj.partition = partition
	master, err := getMaster(partition)
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
	commandResp := commandProcessors[partition].delegatedExec(&reqObj, time.Second)
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