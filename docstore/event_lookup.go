package docstore

import (
	"github.com/v2pro/quokka/kvstore"
	"sync"
)

var eventLookup []*eventLookupPartition

type eventLookupPartition struct {
	partition        uint64
	lookup           map[string]uint64
	lookupMutex      *sync.Mutex
	partitionVersion uint64
}

func (partition *eventLookupPartition) getEventId(entityId string) (uint64, uint64) {
	partition.lookupMutex.Lock()
	defer partition.lookupMutex.Unlock()
	return partition.lookup[entityId], partition.partitionVersion
}

func (partition *eventLookupPartition) setEventId(partitionVersion uint64, entityId string) {
	partition.lookupMutex.Lock()
	defer partition.lookupMutex.Unlock()
	if partition.partitionVersion < partitionVersion {
		partition.partitionVersion = partitionVersion
	}
	if partition.lookup[entityId] < partitionVersion {
		partition.lookup[entityId] = partitionVersion
	}
}

func init() {
	eventLookup = make([]*eventLookupPartition, kvstore.PartitionsCount)
	for i := 0; i < len(eventLookup); i++ {
		eventLookup[i] = &eventLookupPartition{
			partition:   uint64(i),
			lookup:      map[string]uint64{},
			lookupMutex: &sync.Mutex{},
			partitionVersion: 0,
		}
	}
}

func getEventId(partition uint64, entityId string) (uint64, uint64) {
	return eventLookup[partition].getEventId(entityId)
}

func setEventId(partition uint64, partitionVersion uint64, entityId string) {
	eventLookup[partition].setEventId(partitionVersion, entityId)
}
