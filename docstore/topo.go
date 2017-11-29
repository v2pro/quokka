package docstore

import (
	"github.com/v2pro/quokka/kvstore"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/v2pro/plz/countlog"
	"sync"
)

// node.go use topo.go to know the master/slaves of partition
// topo.go use node.go to know the nodes and update topo to optimal state
// node.go will trigger topo.go to update topo if nodes changed

type partitionNodes struct {
	isPromoting bool
	master string
	slaves []string
}

type topology []partitionNodes

var topo = make(topology, kvstore.PartitionsCount)
var topoMutex = &sync.Mutex{}

func (topo topology) get(partitionId uint64) partitionNodes {
	topoMutex.Lock()
	defer topoMutex.Unlock()
	return topo[partitionId]
}

func (topo topology) setPromotingMaster(partitionId uint64, master string) {
	topoMutex.Lock()
	defer topoMutex.Unlock()
	topo[partitionId].master = master
	topo[partitionId].isPromoting = true
	countlog.Info("event!topo.promoting master",
		"partitionId", partitionId,
		"master", master)
}

func (topo topology) setPromotedMaster(partitionId uint64, master string) {
	topoMutex.Lock()
	defer topoMutex.Unlock()
	topo[partitionId].master = master
	topo[partitionId].isPromoting = false
	countlog.Info("event!topo.promoted master",
		"partitionId", partitionId,
			"master", master)
}

// if the master changed, clear the master
// it will be filled when rebalance
func (topo topology) clearMaster(partitionId uint64) {
	topoMutex.Lock()
	defer topoMutex.Unlock()
	oldMaster := topo[partitionId].master
	topo[partitionId].master = ""
	topo[partitionId].isPromoting = false
	countlog.Info("event!topo.clear master",
		"partitionId", partitionId,
		"oldMaster", oldMaster)
}

func (topo topology) refresh(partitionId uint64) (string, error) {
	metadataKey := fmt.Sprintf("partition_%v", partitionId)
	encodedPartition, err := kvstore.GetMetadata(metadataKey)
	if err != nil {
		countlog.Error("event!topo.failed to get partition topo from kvstore",
			"err", err,
			"metadataKey", metadataKey,
			"partitionId", partitionId)
		return "", err
	}
	if len(encodedPartition) == 0 {
		return "", nil
	}
	master := jsoniter.Get(encodedPartition, "Master").ToString()
	if master != "" {
		topo.setPromotedMaster(partitionId, master)
	}
	return master, nil
}

// rebalance is only triggered from node.go
// given this cluster and your current topo
// update the topo to a new optimal state
func (topo topology) rebalance(cluster map[string]*nodeStatus) error {
	var err error
	myMasterPartitionsCount := 0
	for partitionId := uint64(0); partitionId < kvstore.PartitionsCount; partitionId++ {
		master := topo.get(partitionId).master
		if master == "" {
			master, err = topo.refresh(partitionId)
			if err != nil {
				return err
			}
		}
		if master == "" {
			node := chooseMaster(cluster)
			// will be reset when refreshing from kvstore
			// only used in this rebalancing cycle
			node.MasterPartitionsCount += 1
			topo.setPromotingMaster(partitionId, node.Addr)
		}
		if master == thisNodeAddr {
			myMasterPartitionsCount += 1
		}
	}
	cluster[thisNodeAddr].MasterPartitionsCount = myMasterPartitionsCount
	return nil
}

func chooseMaster(cluster map[string]*nodeStatus) *nodeStatus {
	var chosen *nodeStatus
	for _, node := range cluster {
		if chosen == nil {
			chosen = node
		} else {
			if node.MasterPartitionsCount < chosen.MasterPartitionsCount {
				chosen = node
			}
		}
	}
	return chosen
}

func (topo topology) savePromotedMasterInBackground(partitionId uint64) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!topo.savePromotedMasterInBackground.panic",
				"err", recovered,
				"stacktrace", countlog.ProvideStacktrace)
		}
	}()
	master := thisNodeAddr
	topo.setPromotedMaster(partitionId, master)
	encodedPartition, err := jsoniter.Marshal(map[string]string{"Master":master})
	if err != nil {
		countlog.Error("event!topo.failed to marshal partition topo",
			"partitionId", partitionId, "err", err)
		return
	}
	metadataKey := fmt.Sprintf("partition_%v", partitionId)
	err = kvstore.SetMetadata(metadataKey, encodedPartition)
	if err != nil {
		countlog.Error("event!topo.failed to save partition topo",
			"partitionId", partitionId, "err", err)
		return
	}
	countlog.Info("event!topo.promoted master",
		"partitionId", partitionId, "master", master)
	// TODO: broadcast new master to the cluster
}
