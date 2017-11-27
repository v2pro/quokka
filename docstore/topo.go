package docstore

import (
	"net"
	"github.com/v2pro/quokka/kvstore"
	"fmt"
	"errors"
	"github.com/json-iterator/go"
	"github.com/v2pro/plz/countlog"
	"sync/atomic"
	"unsafe"
)

type node struct {
	status nodeStatus
	topo   unsafe.Pointer
}

type partitionNodes struct {
	Master *net.TCPAddr
	Slaves []*net.TCPAddr
}

type topo struct {
	partitions []partitionNodes
}

func (servers partitionNodes) Clone() partitionNodes {
	slaves := make([]*net.TCPAddr, len(servers.Slaves))
	for i := 0; i < len(servers.Slaves); i++ {
		slaves[i] = cloneTCPAddr(servers.Slaves[i])
	}
	return partitionNodes{
		Master: cloneTCPAddr(servers.Master),
		Slaves: slaves,
	}
}

func cloneTCPAddr(addr *net.TCPAddr) *net.TCPAddr {
	copied := *addr
	return &copied
}

func newNode(addr string) *node {
	return &node{status: nodeStatus{Addr: addr}}
}

func (node *node) partition(partition uint64) partitionNodes {
	topo := (*topo)(atomic.LoadPointer(&node.topo))
	if topo == nil {
		return partitionNodes{}
	}
	if topo.partitions == nil {
		return partitionNodes{}
	}
	return topo.partitions[partition]
}

func (node *node) cloneTopo() topo {
	topoPtr := (*topo)(atomic.LoadPointer(&node.topo))
	copied := make([]partitionNodes, kvstore.PartitionsCount)
	if topoPtr != nil && topoPtr.partitions != nil {
		for i := 0; i < kvstore.PartitionsCount; i++ {
			copied[i] = topoPtr.partitions[i].Clone()
		}
	}
	return topo{copied}
}

func (node *node) setTopo(topo topo) {
	atomic.StorePointer(&node.topo, (unsafe.Pointer)(&topo))
}

func (node *node) getMaster(partition uint64) (*net.TCPAddr, error) {
	if partition < 0 || partition >= kvstore.PartitionsCount {
		return nil, errors.New("partition out of range")
	}
	var err error
	partitionNodes := node.partition(partition)
	if partitionNodes.Master == nil {
		partitionNodes, err = node.refreshPartition(partition)
		if err != nil {
			countlog.Error("event!cluster.failed to refresh partition servers",
				"partition", partition,
				"err", err)
		}
	}
	if partitionNodes.Master != nil {
		return partitionNodes.Master, nil
	}
	return node.promotePartitionMaster(partition)
}

func (node *node) refreshPartition(partition uint64) (partitionNodes, error) {
	metadataKey := fmt.Sprintf("partition_%v_servers", partition)
	encodedServers, err := kvstore.GetMetadata(metadataKey)
	if err != nil {
		countlog.Error("event!topo.failed to get servers",
			"err", err,
			"metadataKey", metadataKey,
			"partition", partition)
		return partitionNodes{}, err
	}
	if len(encodedServers) == 0 {
		countlog.Error("event!topo.received empty server list",
			"partition", partition,
			"metadataKey", metadataKey,
			"encodedServers", encodedServers)
		return partitionNodes{}, errors.New("received empty server list")
	}
	var servers partitionNodes
	err = jsoniter.Unmarshal(encodedServers, &servers)
	if err != nil {
		countlog.Error("event!topo.failed to unmarshal partition servers metadata",
			"err", err,
			"partition", partition,
			"metadataKey", metadataKey,
			"encodedServers", encodedServers)
		return partitionNodes{}, fmt.Errorf("failed to unmarshal partition servers metadata: %s", err.Error())
	}
	newTopo := node.cloneTopo()
	newTopo.partitions[partition] = servers
	node.setTopo(newTopo)
	return servers, nil
}

func (node *node) promotePartitionMaster(partition uint64) (*net.TCPAddr, error) {
	nodes, err := queryCluster()
	if err != nil {
		return nil, err
	}
	var newMaster *nodeStatus
	for _, node := range nodes {
		if newMaster == nil {
			newMaster = &node
		} else {
			if node.MasterPartitionsCount < newMaster.MasterPartitionsCount {
				newMaster = &node
			}
		}
	}
	if newMaster == nil {
		return nil, errors.New("cluster has no usable node")
	}
	return net.ResolveTCPAddr("tcp", newMaster.Addr)
}

//func (node *node) setPartitionServers(partition uint64, servers *partitionNodes) error {
//	metadataKey := fmt.Sprintf("partition_%v_servers", partition)
//	encodedServers, err := jsoniter.Marshal(servers)
//	if err != nil {
//		countlog.Error("event!topo.failed to marshal partition servers metadata",
//			"err", err,
//			"partition", partition,
//			"metadataKey", metadataKey,
//			"encodedServers", encodedServers)
//		return fmt.Errorf("failed to marshal partition servers metadata: %s", err.Error())
//	}
//	err = kvstore.SetMetadata(metadataKey, encodedServers)
//	if err != nil {
//		countlog.Error("event!topo.failed to set servers",
//			"err", err,
//			"metadataKey", metadataKey,
//			"encodedServers", encodedServers,
//			"partition", partition)
//		return err
//	}
//	return nil
//}

func (node *node) RebalanceInBackground() {

}
