package docstore

import (
	"github.com/v2pro/quokka/kvstore"
	"errors"
)

type entityLookup struct {
	memLookup *memLookup
	kvstoreLookup *kvstoreLookup
}

func newEntityLookup() *entityLookup {
	return &entityLookup{
		memLookup: &memLookup{
			cacheSize: 5000,
			cache1: map[string]interface{}{},
			cache2: map[string]interface{}{},
		},
		kvstoreLookup: &kvstoreLookup{
			prefix: "entity_",
		},
	}
}

func (lookup *entityLookup) getEntity(partition uint64, entityType string, entityId string) (*entity, error) {
	cachedVal := lookup.memLookup.getCacheValue(entityId)
	if cachedVal != nil {
		return cachedVal.(*entity), nil
	}
	eventId, err := lookup.kvstoreLookup.getEventId(partition, entityId, lookup.memLookup.cache2StartVersion)
	if err != nil {
		return nil, err
	}
	return loadEntity(partition, entityType, entityId, eventId)
}

func (lookup *entityLookup) cacheEntity(entityId string, value *entity, version uint64) {
	lookup.memLookup.setCacheValue(entityId, value, version)
}

func (lookup *entityLookup) setEventId(partition uint64, entityId string, eventId uint64) error {
	return lookup.kvstoreLookup.setEventId(partition, entityId, eventId)
}

type commandLookup struct {
	memLookup *memLookup
	kvstoreLookup *kvstoreLookup
}

func newCommandLookup() *commandLookup {
	return &commandLookup{
		memLookup: &memLookup{
			cacheSize: 5000,
			cache1: map[string]interface{}{},
			cache2: map[string]interface{}{},
		},
		kvstoreLookup: &kvstoreLookup{
			prefix: "command_",
		},
	}
}

func (lookup *commandLookup) getCommand(partition uint64, commandId string) ([]byte, error) {
	cachedVal := lookup.memLookup.getCacheValue(commandId)
	if cachedVal != nil {
		return cachedVal.([]byte), nil
	}
	eventId, err := lookup.kvstoreLookup.getEventId(partition, commandId, lookup.memLookup.cache2StartVersion)
	if err != nil {
		return nil, err
	}
	return getCommandResponse(partition, eventId)
}

func (lookup *commandLookup) cacheCommand(commandId string, value []byte, version uint64) {
	lookup.memLookup.setCacheValue(commandId, value, version)
}

func (lookup *commandLookup) setEventId(partition uint64, commandId string, eventId uint64) error {
	return lookup.kvstoreLookup.setEventId(partition, commandId, eventId)
}

type memLookup struct {
	cacheSize int
	cache1StartVersion uint64
	cache2StartVersion uint64
	cache1 map[string]interface{}
	cache2 map[string]interface{}
}

func (lookup *memLookup) setCacheValue(key string, value interface{}, version uint64) {
	if len(lookup.cache2) < lookup.cacheSize {
		if lookup.cache2StartVersion == 0 {
			lookup.cache2StartVersion = version
		}
		lookup.cache2[key] = value
	}
	if len(lookup.cache1) < lookup.cacheSize {
		if lookup.cache1StartVersion == 0 {
			lookup.cache1StartVersion = version
		}
		lookup.cache1[key] = value
	}
	// expire cache2, move cache1 to cache2
	lookup.cache2StartVersion = lookup.cache1StartVersion
	lookup.cache2 = lookup.cache1
	lookup.cache1StartVersion = 0
	lookup.cache1 = map[string]interface{}{}
}

func (lookup *memLookup) getCacheValue(key string) interface{} {
	val := lookup.cache1[key]
	if val != nil {
		return val
	}
	return lookup.cache2[key]
}

type kvstoreLookup struct {
	prefix string
}

func (lookup *kvstoreLookup) getEventId(partition uint64, key string, minVersion uint64) (uint64, error) {
	offset, err := kvstore.GetMonotonic(partition, "offset")
	if err != nil {
		return 0, err
	}
	if offset < minVersion {
		return 0, errors.New("lookup is too old")
	}
	value, err := kvstore.GetMonotonic(partition, lookup.prefix + key)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (lookup *kvstoreLookup) setEventId(partition uint64, key string, eventId uint64) error {
	return kvstore.SetMonotonic(partition, lookup.prefix + key, eventId)
}