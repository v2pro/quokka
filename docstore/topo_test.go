package docstore

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/quokka/kvstore/memkv"
)

func Test_rebalance_init(t *testing.T) {
	memkv.ResetKVStore()
	should := require.New(t)
	thisNodeAddr = "127.0.0.1:2515"
	topo.rebalance(map[string]*nodeStatus{
		"127.0.0.1:2515": {Addr:"127.0.0.1:2515"}})
	should.Equal("127.0.0.1:2515", topo.get(0).master)
	should.True(topo.get(0).isPromoting)
}

func Test_rebalance_new_nodes(t *testing.T) {
	memkv.ResetKVStore()
	should := require.New(t)
	thisNodeAddr = "127.0.0.1:2515"
	cluster := map[string]*nodeStatus{
		"127.0.0.1:2515": {Addr: "127.0.0.1:2515"}}
	topo.rebalance(cluster)
	should.Equal("127.0.0.1:2515", topo.get(0).master)
	should.True(topo.get(0).isPromoting)
	cluster["127.0.0.1:2516"] = &nodeStatus{Addr:"127.0.0.1:2515"}
	topo.rebalance(cluster)
	should.Equal(997 / 2, cluster[thisNodeAddr].MasterPartitionsCount)
}