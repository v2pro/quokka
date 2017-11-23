package docstore

import (
	"github.com/v2pro/quokka/kvstore"
	"errors"
)

type memRow []byte
type memPartition [][]byte
var memPartitions = make([]memPartition, kvstore.PartitionsCount)

func resetMemKVStore() {
	memPartitions = make([]memPartition, kvstore.PartitionsCount)
	kvstore.Get = memGet
	kvstore.Append = memAppend
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