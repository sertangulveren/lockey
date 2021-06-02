package lockey

import (
	"sync"
)

type Lockey struct {
	grandMu sync.Mutex
	store   map[string]*storeItem
}

type storeItem struct {
	mu      sync.Mutex
	reserve int
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(storeItem)
	},
}

// New Lockey.
//
// Creates a new Lockey.
func New() *Lockey {
	return &Lockey{
		grandMu: sync.Mutex{},
		store:   make(map[string]*storeItem),
	}
}

// Lock locks the key.
//
// Creates a locked mutex for the given key.
func (l *Lockey) Lock(key string) {
	l.build(key).lock()
}

// Unlock unlocks the key.
//
// Unlocks the mutex of the given key.
func (l *Lockey) Unlock(key string) {
	l.destroy(key).unlock()
}

func (l *Lockey) build(key string) *storeItem {
	l.grandMu.Lock()
	defer l.grandMu.Unlock()

	item, ok := l.store[key]
	if !ok {
		// If there is no item, get one from the pool.
		item = pool.Get().(*storeItem)
		l.store[key] = item
	}

	// A new reservation has been made.
	// Increase the count.
	item.reserve++

	return item
}

func (l *Lockey) destroy(key string) *storeItem {
	l.grandMu.Lock()
	defer l.grandMu.Unlock()

	item, ok := l.store[key]
	if !ok {
		panic("There is no such lock for key: " + key)
	}
	item.reserve--

	// Remove item from the store if there is no reservation for the key.
	if !item.isReserved() {
		// Put it back to pool.
		pool.Put(item)

		// Remove from the store.
		delete(l.store, key)
	}
	return item
}

func (i *storeItem) lock() {
	i.mu.Lock()
}

func (i *storeItem) unlock() {
	i.mu.Unlock()
}

// Is there any reservation?
func (i *storeItem) isReserved() bool {
	return i.reserve > 0
}
