package docstore

import (
	"github.com/v2pro/quokka/kvstore"
	"github.com/json-iterator/go"
	"time"
	"net"
)

// stored in metadata server_[addr]
type nodeStatus struct {
	Addr                  *net.TCPAddr
	Heartbeat             time.Time
	MasterPartitionsCount int
	SlavePartitionsCount  int
}


// tell other nodes about myself
func joinCluster(nodeStatus nodeStatus) error {
	nodeStatus.Heartbeat = time.Now()
	status, err := jsoniter.ConfigFastest.Marshal(nodeStatus)
	if err != nil {
		return err
	}
	return kvstore.SetMetadata("node_"+nodeStatus.Addr.String(), status)
}

// get other nodes
func queryCluster() ([]nodeStatus, error) {
	iter, err := kvstore.ScanMetadata("node_", "node"+string([]byte{'_' + 1}))
	if err != nil {
		return nil, err
	}
	cluster := []nodeStatus{}
	for {
		batch, err := iter()
		if err != nil {
			return nil, err
		}
		if batch == nil {
			return cluster, nil
		}
		for _, row := range batch {
			var node nodeStatus
			err = jsoniter.ConfigFastest.Unmarshal(row.Value, &node)
			if err != nil {
				return nil, err
			}
			cluster = append(cluster, node)
		}
	}
}
