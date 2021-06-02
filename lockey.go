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
