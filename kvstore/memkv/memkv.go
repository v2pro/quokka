package memkv

import (
	"github.com/v2pro/quokka/kvstore"
	"errors"
)

type memRow []byte
type memPartition [][]byte

var memPartitions = make([]memPartition, kvstore.PartitionsCount)
var memMetadata = []*kvstore.MetadataRow{}
var memMonotonics = []map[string]uint64{}

func ResetKVStore() {
	memPartitions = make([]memPartition, kvstore.PartitionsCount)
	memMonotonics = make([]map[string]uint64, kvstore.PartitionsCount)
	for i := 0; i < len(memMonotonics); i++ {
		memMonotonics[i] = map[string]uint64{}
	}
	memMetadata = []*kvstore.MetadataRow{}
	kvstore.Get = memGet
	kvstore.Append = memAppend
	kvstore.GetMetadata = memGetMetadata
	kvstore.SetMetadata = memSetMetadata
	kvstore.ScanMetadata = memScanMetadata
	kvstore.GetMonotonic = memGetMonotonic
	kvstore.SetMonotonic = memSetMonotonic
}

func memGet(partition uint64, namespace string, rowKey uint64) ([]byte, error) {
	if partition < 0 || partition >= uint64(len(memPartitions)) {
		return nil, errors.New("partition not found")
	}
	rows := memPartitions[partition]
	if rowKey <= 0 || rowKey >= uint64(len(rows)) {
		return nil, errors.New("row not found")
	}
	return rows[rowKey], nil
}

func memAppend(partition uint64, namespace string, rowKey uint64, rowValue []byte) error {
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

func memGetMonotonic(partition uint64, namespace string, key string) (uint64, error) {
	return memMonotonics[partition][key], nil
}

func memSetMonotonic(partition uint64, namespace string, key string, value uint64) error {
	if value > memMonotonics[partition][key] {
		memMonotonics[partition][key] = value
	}
	return nil
}

func memGetMetadata(key string) ([]byte, error) {
	for _, row := range memMetadata {
		if row.Key == key {
			return row.Value, nil
		}
	}
	return nil, nil
}

func memSetMetadata(key string, value []byte) error {
	for _, row := range memMetadata {
		if row.Key == key {
			row.Value = value
			return nil
		}
	}
	memMetadata = append(memMetadata, &kvstore.MetadataRow{
		Key:   key,
		Value: value,
	})
	return nil
}

func memScanMetadata(fromKey string, toKey string) (kvstore.MetadataRowIterator, error) {
	returned := false
	return func() ([]kvstore.MetadataRow, error) {
		if returned {
			return nil, nil
		}
		returned = true
		found := []kvstore.MetadataRow{}
		for _, row := range memMetadata {
			if row.Key >= fromKey && row.Key < toKey {
				found = append(found, *row)
			}
		}
		return found, nil
	}, nil
}
