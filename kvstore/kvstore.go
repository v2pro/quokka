package kvstore

import "errors"

const PartitionsCount = 997
var KeyConflictError = errors.New("key conflict")

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
// entity to event map
// bloomfilter for unique indices

var GetMetadata = func(key string) ([]byte, error) {
	return nil, errors.New("not implemented")
}

var SetMetadata = func(key string, value []byte) error {
	return errors.New("not implemented")
}
