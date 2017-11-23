package docstore

import (
	"github.com/v2pro/quokka/kvstore"
	"errors"
)

type memRow []byte
type memPartition [][]byte
var memPartitions = make([]memPartition, kvstore.PartitionsCount)
var memMetadata = map[string][]byte{}

func resetMemKVStore() {
	memPartitions = make([]memPartition, kvstore.PartitionsCount)
	memMetadata = map[string][]byte{}
	kvstore.Get = memGet
	kvstore.Append = memAppend
	kvstore.GetMetadata = memGetMetadata
	kvstore.SetMetadata = memSetMetadata
}

func memGet(partition uint64, rowKey uint64) ([]byte, error)  {
	if partition < 0 || partition >= uint64(len(memPartitions)) {
		return nil, errors.New("partition not found")
	}
	rows := memPartitions[partition]
	if rowKey <= 0 || rowKey >= uint64(len(rows)) {
		return nil, errors.New("row not found")
	}
	return rows[rowKey], nil
}

func memAppend(partition uint64, rowKey uint64, rowValue []byte) error {
	if partition < 0 || partition >= uint64(len(memPartitions)) {
		return errors.New("partition not found")
	}
	rows := memPartitions[partition]
	if rows == nil {
		rows = make(memPartition, 1)
	}
	if rowKey != uint64(len(rows)) {
		return errors.New("row key should be +1")
	}
	rows = append(rows, rowValue)
	memPartitions[partition] = rows
	return nil
}

func memGetMetadata(key string) ([]byte, error) {
	return memMetadata[key], nil
}

func memSetMetadata(key string, value []byte) error {
	memMetadata[key] = value
	return nil
}