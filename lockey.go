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
