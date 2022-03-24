package main

import "sync"

type Barrier struct {
	total        int
	count        int
	mutex        *sync.Mutex
	conditionVar *sync.Cond
}

func NewBarrier(size int) *Barrier {
	lockToUse := &sync.Mutex{}
	condToUse := sync.NewCond(lockToUse)
	return &Barrier{
		total:        size,
		count:        size,
		mutex:        lockToUse,
		conditionVar: condToUse,
	}
}

func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.count--
	if b.count == 0 {
		b.count = b.total
		b.conditionVar.Broadcast()
	} else {
		b.conditionVar.Wait()
	}
	b.mutex.Unlock()
}
