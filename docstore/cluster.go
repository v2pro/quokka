package docstore

import (
	"net"
	"sync"
	"github.com/v2pro/quokka/kvstore"
	"fmt"
	"errors"
	"github.com/json-iterator/go"
	"github.com/v2pro/plz/countlog"
	"net/http"
	"bytes"
	"io/ioutil"
)

type partitionServers struct {
	mutex  *sync.Mutex
	Master *net.TCPAddr
	Slaves []*net.TCPAddr
}
var cluster []*partitionServers

func init() {
	cluster = make([]*partitionServers, kvstore.PartitionsCount)
	for i := 0; i < len(cluster); i++ {
		cluster[i] = &partitionServers{
			mutex: &sync.Mutex{},
		}
	}
}

func getMaster(partition uint64) (*net.TCPAddr, error) {
	if partition < 0 || partition >= uint64(len(cluster)) {
		return nil, errors.New("partition out of range")
	}
	var err error
	servers := cluster[partition]
	if servers.Master == nil {
		servers, err = refreshPartitionServers(partition)
		if err != nil {
			return nil, err
		}
	}
	return servers.Master, nil
}

func refreshPartitionServers(partition uint64) (*partitionServers, error) {
	metadataKey := fmt.Sprintf("partition_%v_servers", partition)
	encodedServers, err := kvstore.GetMetadata(metadataKey)
	if err != nil {
		countlog.Error("event!cluster.failed to get servers",
			"err", err,
			"metadataKey", metadataKey,
			"partition", partition)
		return nil, err
	}
	if len(encodedServers) == 0 {
		countlog.Error("event!cluster.received empty server list",
			"partition", partition,
			"metadataKey", metadataKey,
			"encodedServers", encodedServers)
		return nil, errors.New("received empty server list")
	}
	servers := cluster[partition]
	servers.mutex.Lock()
	defer servers.mutex.Unlock()
	err = jsoniter.Unmarshal(encodedServers, servers)
	if err != nil {
		countlog.Error("event!cluster.failed to unmarshal partition servers metadata",
			"err", err,
			"partition", partition,
			"metadataKey", metadataKey,
			"encodedServers", encodedServers)
		return nil, fmt.Errorf("failed to unmarshal partition servers metadata: %s", err.Error())
	}
	if servers.Master == nil {
		return nil, errors.New("master unknown")
	}
	return servers, nil
}

func setPartitionServers(partition uint64, servers *partitionServers) error {
	metadataKey := fmt.Sprintf("partition_%v_servers", partition)
	encodedServers, err := jsoniter.Marshal(servers)
	if err != nil {
		countlog.Error("event!cluster.failed to marshal partition servers metadata",
			"err", err,
			"partition", partition,
			"metadataKey", metadataKey,
			"encodedServers", encodedServers)
		return fmt.Errorf("failed to marshal partition servers metadata: %s", err.Error())
	}
	err = kvstore.SetMetadata(metadataKey, encodedServers)
	if err != nil {
		countlog.Error("event!cluster.failed to set servers",
			"err", err,
			"metadataKey", metadataKey,
			"encodedServers", encodedServers,
			"partition", partition)
		return err
	}
	return nil
}
