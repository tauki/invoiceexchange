package eventhandler

import (
	"hash/fnv"
	"sync"
)

const (
	numOfInvoiceLocks = 1000
)

type invoiceLocks struct {
	locks []sync.Mutex
}

func newInvoiceLocks() *invoiceLocks {
	return &invoiceLocks{
		locks: make([]sync.Mutex, numOfInvoiceLocks),
	}
}

func (l *invoiceLocks) Lock(id string) {
	l.locks[Hash(id)%numOfInvoiceLocks].Lock()
}

func (l *invoiceLocks) Unlock(id string) {
	l.locks[Hash(id)%numOfInvoiceLocks].Unlock()
}

func Hash(data string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(data))
	return h.Sum32()
}
