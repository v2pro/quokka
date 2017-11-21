package agent

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/v2pro/plz/countlog"
	"sync"
	"time"
	"strings"
)

var Mux = &http.ServeMux{}

func init() {
	Mux.HandleFunc("/active-processes", listActiveProcesses)
	Mux.HandleFunc("/ping", ping)
	Mux.HandleFunc("/", homepage)
}

func homepage(respWriter http.ResponseWriter, request *http.Request) {
	respWriter.Write([]byte("OK"))
}

type process struct {
	ProcessId   int
	ProcessInfo map[string]interface{} // more info about what the process is about
	LastHeartbeat  time.Time
}

var activeProcesses = map[int]process{}
var activeProcessesMutex = &sync.Mutex{}

func ping(respWriter http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		countlog.Error("event!agent.failed to read ping request", "err", err)
		respWriter.Write([]byte(`{"errno": 1, "errmsg": "failed to read ping request"}`))
		return
	}
	var proc process
	err = json.Unmarshal(body, &proc)
	if err != nil {
		countlog.Error("event!agent.failed to ping poll request", "err", err, "body", body)
		respWriter.Write([]byte(`{"errno": 1, "errmsg": "failed to unmarshal ping request"}`))
		return
	}
	hijacker, _ := respWriter.(http.Hijacker)
	if hijacker == nil {
		countlog.Error("event!agent.failed to take hijacker", "err", err)
		respWriter.Write([]byte(`{"errno": 1, "errmsg": "failed to take hijacker"}`))
		return
	}
	conn, newWriter, err := hijacker.Hijack()
	if err != nil {
		countlog.Error("event!agent.failed to hijack", "err", err)
		respWriter.Write([]byte(`{"errno": 1, "errmsg": "failed to hijack"}`))
		return
	}
	done := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Second)
			updateActiveProcess(proc)
			select {
			case <-done:
				return
			default:
				continue
			}
		}
	}()
	conn.SetReadDeadline(time.Now().Add(time.Minute * 5))
	conn.Read(make([]byte, 1))
	close(done)
	resp := http.Response{}
	resp.Body = ioutil.NopCloser(strings.NewReader(`{"errno": 0, "errmsg": "ok, now ping again"}`))
	resp.Write(newWriter)
	return
}

func listActiveProcesses(respWriter http.ResponseWriter, request *http.Request) {
	activeProcessesMutex.Lock()
	defer activeProcessesMutex.Unlock()
	resp, err := json.Marshal(map[string]interface{}{
		"errno": 0,
		"data":  activeProcesses,
	})
	if err != nil {
		countlog.Error("event!agent.failed to active processes", "err", err)
		respWriter.Write([]byte(`{"errno": 1, "errmsg": "can not marshal active processes"}`))
		return
	}
	respWriter.Write(resp)
}

func updateActiveProcess(process process) {
	activeProcessesMutex.Lock()
	defer activeProcessesMutex.Unlock()
	process.LastHeartbeat = time.Now()
	activeProcesses[process.ProcessId] = process
}
