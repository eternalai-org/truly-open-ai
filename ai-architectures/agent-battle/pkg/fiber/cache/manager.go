package cache

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// msgp -file="manager.go" -o="manager_msgp.go" -tests=false -unexported
type item struct {
	body      []byte
	ctype     []byte
	cencoding []byte
	status    int
	exp       uint64
	headers   map[string][]byte
	// used for finding the item in an indexed heap
	heapidx int
}

//msgp:ignore manager
type manager struct {
	pool    sync.Pool
	storage fiber.Storage
}

func newManager(storage fiber.Storage) *manager {
	// Create new storage handler
	manager := &manager{
		pool: sync.Pool{
			New: func() interface{} {
				return new(item)
			},
		},
		storage: storage,
	}
	return manager
}

// acquire returns an *entry from the sync.Pool
func (m *manager) acquire() *item {
	return m.pool.Get().(*item) //nolint:forcetypeassert // We store nothing else in the pool
}

// release and reset *entry to sync.Pool
func (m *manager) release(e *item) {
	// don't release item if we using memory storage
	if m.storage != nil {
		return
	}
	e.body = nil
	e.ctype = nil
	e.status = 0
	e.exp = 0
	e.headers = nil
	m.pool.Put(e)
}

// get data from storage or memory
func (m *manager) get(key string) *item {
	it := m.acquire()
	raw, err := m.storage.Get(key)
	if err != nil {
		return it
	}
	if raw != nil {
		if _, err := it.UnmarshalMsg(raw); err != nil {
			return it
		}
	}
	return it
}

// get raw data from storage or memory
func (m *manager) getRaw(key string) []byte {
	raw, _ := m.storage.Get(key) //nolint:errcheck // TODO: Handle error here
	return raw
}

// set data to storage or memory
func (m *manager) set(key string, it *item, exp time.Duration) {
	if raw, err := it.MarshalMsg(nil); err == nil {
		_ = m.storage.Set(key, raw, exp) //nolint:errcheck // TODO: Handle error here
	}
	// we can release data because it's serialized to database
	m.release(it)
}

// set data to storage or memory
func (m *manager) setRaw(key string, raw []byte, exp time.Duration) {
	_ = m.storage.Set(key, raw, exp) //nolint:errcheck // TODO: Handle error here
}

// delete data from storage or memory
func (m *manager) del(key string) {
	_ = m.storage.Delete(key) //nolint:errcheck // TODO: Handle error here
}
