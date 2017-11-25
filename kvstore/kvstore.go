package kvstore

import (
	"errors"
	"github.com/spaolacci/murmur3"
)

const PartitionsCount = 997
var KeyConflictError = errors.New("key conflict")

func HashToPartition(entityId string) uint64 {
	hash := murmur3.Sum64([]byte(entityId))
	partitionIndex := hash % PartitionsCount
	return partitionIndex
}

// kvstore for event log

type Row struct {
	Key   uint64
	Value []byte
}
type RowIterator func() ([]Row, error)

var Get = func(partition uint64, rowKey uint64) ([]byte, error) {
	return nil, errors.New("not implemented")
}

var Scan = func(partition uint64, fromRowKey uint64) (RowIterator, error) {
	return nil, errors.New("not implemented")
}

var Append = func(partition uint64, rowKey uint64, rowValue []byte) error {
	return errors.New("not implemented")
}

// metadata:
// partition master and slaves

var GetMetadata = func(key string) ([]byte, error) {
	return nil, errors.New("not implemented")
}

var SetMetadata = func(key string, value []byte) error {
	return errors.New("not implemented")
}

// entity lookup:
// entity id => event id
// command lookup:
// command id => event id

var GetMonotonic = func(partition uint64, key string) (uint64, error) {
	return 0, errors.New("not implemented")
}

var SetMonotonic = func(partition uint64, key string, value uint64) error {
	return errors.New("not implemented")
}
