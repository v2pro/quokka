package agent

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/v2pro/plz/countlog"
	"sync"
	"time"
	"strings"
	"sort"
)

var Mux = &http.ServeMux{}

func init() {
	Mux.HandleFunc("/active-processes", listActiveProcesses)
	Mux.HandleFunc("/ping", ping)
	Mux.HandleFunc("/", homepage)
	go gcActiveProcessesInBackground()
}

func homepage(respWriter http.ResponseWriter, request *http.Request) {
	respWriter.Write([]byte(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <!-- 引入样式 -->
    <link rel="stylesheet" href="http://cdn.jsdeliver.net/npm/element-ui/lib/theme-chalk/index.css">
    <style>
        body {
            font-family: Helvetica
        }
    </style>
</head>
<body>
<div id="app">
	<el-table
    :data="processes"
    stripe
	@row-click="onRowClick"
    style="width: 100%">
    <el-table-column
      prop="Service"
      label="Service"
      width="180">
    </el-table-column>
    <el-table-column
      prop="Cluster"
      label="Cluster"
      width="180">
    </el-table-column>
    <el-table-column
      prop="ProcessId"
      label="Process Id">
    </el-table-column>
    <el-table-column
      prop="HeartbeatDelay"
      label="Heartbeat Delay">
    </el-table-column>
  </el-table>
</div>
</body>
<!-- 先引入 Vue -->
<script src="http://cdn.jsdeliver.net/npm/vue/dist/vue.js"></script>
<!-- 引入组件库 -->
<script src="http://cdn.jsdeliver.net/npm/element-ui/lib/index.js"></script>
<script src="http://cdn.jsdeliver.net/npm/axios/dist/axios.min.js"></script>
<script>
    var $vue = new Vue({
        el: '#app',
        data: function() {
            return { processes: [] }
        },
		created: function() {
			var me = this;
			setInterval(function() {
                axios.get('/active-processes?ts=' + Date.now())
                    .then(function (resp) {
						if (resp.data) {
	                        me.processes = resp.data.data;
						}
                    });
			}, 3000);
		},
		methods: {
			onRowClick: function(row) {
				window.location.href = "/processes/" + row.ProcessId
			}
		}
    });
    axios.interceptors.response.use(function (response) {
        return response;
    }, function (error) {
        $vue.$notify.error({
            position: 'bottom-right',
            title: error.message,
            message: error.response
        });
        return error;
    });
</script>
</html>
	`))
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
	if proc.ProcessInfo == nil {
		proc.ProcessInfo = map[string]interface{}{}
	}
	proc.ProcessInfo["ProcessId"] = proc.ProcessId
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
	sortedPids := sort.IntSlice{}
	for pid := range activeProcesses {
		sortedPids = append(sortedPids, pid)
	}
	sort.Sort(sortedPids)
	sortedProcesses := []map[string]interface{}{}
	for _, pid := range sortedPids {
		procInfo := activeProcesses[pid].ProcessInfo
		procInfo["HeartbeatDelay"] = time.Now().Sub(procInfo["LastHeartbeat"].(time.Time)).String()
		sortedProcesses = append(sortedProcesses, procInfo)
	}
	resp, err := json.Marshal(map[string]interface{}{
		"errno": 0,
		"data":  sortedProcesses,
	})
	if err != nil {
		countlog.Error("event!agent.failed to marshal active processes", "err", err)
		respWriter.Write([]byte(`{"errno": 1, "errmsg": "can not marshal active processes"}`))
		return
	}
	respWriter.Write(resp)
}

func updateActiveProcess(process process) {
	activeProcessesMutex.Lock()
	defer activeProcessesMutex.Unlock()
	process.LastHeartbeat = time.Now()
	process.ProcessInfo["LastHeartbeat"] = process.LastHeartbeat
	activeProcesses[process.ProcessId] = process
}

func gcActiveProcessesInBackground() {
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!agent.gcActiveProcessesInBackground.panic", "err", recovered,
				"stacktrace", countlog.ProvideStacktrace)
		}
	}()
	for {
		time.Sleep(time.Second * 10)
		gcActiveProcessesOneRound()
	}
}
func gcActiveProcessesOneRound() {
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!agent.gcActiveProcessesOneRound.panic", "err", recovered,
				"stacktrace", countlog.ProvideStacktrace)
		}
	}()
	activeProcessesMutex.Lock()
	defer activeProcessesMutex.Unlock()
	now := time.Now()
	newMap := map[int]process{}
	expiredProcessesCount := 0
	for processId, process := range activeProcesses {
		if now.Sub(process.LastHeartbeat) < time.Second * 10 {
			newMap[processId] = process
		} else {
			expiredProcessesCount++
		}
	}
	activeProcesses = newMap
	countlog.Trace("event!agent.gcActiveProcessesOneRound",
		"expiredProcessesCount", expiredProcessesCount)
}