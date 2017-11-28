package docstore

import (
	"github.com/v2pro/quokka/kvstore"
	"github.com/json-iterator/go"
	"time"
	"github.com/v2pro/plz/countlog"
	"errors"
)

// node.go use cluster.go to track the cluster nodes, including
// let other node know me
// know other nodes
// track node health, and detect faulty nodes

// stored in metadata server_[addr]
type nodeStatus struct {
	Addr      string
	Heartbeat time.Time
	// we use the master partitions count saved by heartbeat
	// instead of inferring the load from partition topo
	// because the partition topo might out of sync
	// this MasterPartitionsCount is saved with heartbeat, should be more reliable
	MasterPartitionsCount int
	isDead                bool // according to heartbeat
	isBlacklisted         *time.Time // according to error rate
}

var thisNodeAddr string

func publishNode(addr string) error {
	thisNodeAddr = addr
	cluster := map[string]*nodeStatus{
		addr: {Addr: addr},
	}
	_, err := refreshMasterNodesStatus(cluster)
	if err != nil {
		return err
	}
	err = topo.rebalance(cluster)
	if err != nil {
		return err
	}
	err = joinCluster(cluster[thisNodeAddr], false)
	if err != nil {
		return err
	}
	go refreshAndRebalanceInBackground(cluster)
	return nil
}

// if forward to master and master replied it failed to commit event log
// which means it has lost the master position, trigger rebalance
// if failed to http post master/slave for a period of time, trigger rebalance
var topoChanged = make(chan bool, 1)

func triggerTopoChanged() {
	// trigger immediate rebalance
	select {
	case topoChanged <- true:
	default:
		// if queue is not empty, which means it is already triggered
		return
	}
}

// rebalance will only be triggered from this goroutine
func refreshAndRebalanceInBackground(cluster map[string]*nodeStatus) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!cluster.trackClusterInBackground.panic",
				"err", recovered,
				"stacktrace", countlog.ProvideStacktrace)
		}
	}()
	for {
		refreshAndRebalanceOnce(cluster)
	}
}

func refreshAndRebalanceOnce(cluster map[string]*nodeStatus) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!cluster.refreshAndRebalanceOnce.panic",
				"err", recovered,
				"stacktrace", countlog.ProvideStacktrace)
		}
	}()
	isTopoChanged := sleepBetweenRefreshing()
	isClusterChanged, err := refreshMasterNodesStatus(cluster)
	if err != nil {
		countlog.Warn("event!cluster.failed to refresh cluster", "err", err)
	}
	if isTopoChanged || isClusterChanged {
		topo.rebalance(cluster)
	}
	err = joinCluster(cluster[thisNodeAddr], false)
	if err != nil {
		countlog.Warn("event!cluster.failed to update heartbeat", "err", err)
	}
}

func sleepBetweenRefreshing() bool {
	timer := time.NewTimer(time.Minute)
	select {
	case <-timer.C:
		return false
	case <-topoChanged:
		timer.Stop()
		return true
	}
}

// tell other nodes about myself
func joinCluster(node *nodeStatus, isSlave bool) error {
	node.Heartbeat = time.Now()
	statusJson, err := jsoniter.ConfigFastest.Marshal(node)
	if err != nil {
		return err
	}
	if isSlave {
		return kvstore.SetMetadata("slave_node_"+node.Addr, statusJson)
	} else {
		return kvstore.SetMetadata("master_node_"+node.Addr, statusJson)
	}
}

// get other master candidates
func refreshMasterNodesStatus(cluster map[string]*nodeStatus) (bool, error) {
	iter, err := kvstore.ScanMetadata("node_", "node"+string([]byte{'_' + 1}))
	if err != nil {
		return false, err
	}
	isClusterChanged := false
	for {
		batch, err := iter()
		if err != nil {
			return isClusterChanged, err
		}
		if batch == nil {
			return isClusterChanged, nil
		}
		for _, row := range batch {
			var node nodeStatus
			err = jsoniter.ConfigFastest.Unmarshal(row.Value, &node)
			if err != nil {
				return isClusterChanged, err
			}
			existingNode := cluster[node.Addr]
			if existingNode != nil {
				isClusterChanged = existingNode.update(&node)
			} else {
				cluster[node.Addr] = &node
				isClusterChanged = true
			}
		}
	}
}

func (node *nodeStatus) update(newNode *nodeStatus) bool {
	node.Heartbeat = newNode.Heartbeat
	isClusterChanged := false
	if time.Now().Sub(node.Heartbeat) > time.Hour*24 {
		if !node.isDead {
			node.isDead = true
			isClusterChanged = true
		}
	} else if node.isDead {
		node.isDead = false
		isClusterChanged = true
	}
	// TODO: update isBlacklisted according to error rate
	return isClusterChanged
}

func queryWholeCluster() ([]string, error) {
	return nil, errors.New("not implemented")
}
