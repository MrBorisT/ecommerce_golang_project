package mycache

import (
	"sync"
	"time"
)

type cache[T comparable] struct {
	//simple hash map with any value
	m map[T]*any
	//map for ttl timers
	mTTL map[T]*time.Timer
	//for thread-safety
	lock sync.RWMutex

	ttl time.Duration
}

func NewMyCache[T comparable](ttl time.Duration) *cache[T] {
	return &cache[T]{
		m:    make(map[T]*any),
		mTTL: make(map[T]*time.Timer),
		ttl:  ttl,
	}
}

func (c *cache[T]) SetValue(key T, value any) {
	c.lock.Lock()
	c.m[key] = &value
	ttlTimer := time.NewTimer(c.ttl)
	c.mTTL[key] = ttlTimer
	c.lock.Unlock()
	go func() {
		<-ttlTimer.C
		c.ClearValue(key)
	}()
}

// private method to retrieve value
func (c *cache[T]) GetValue(key T) (*any, bool) {
	c.lock.RLock()
	val, ok := c.m[key]
	c.lock.RUnlock()
	if !ok {
		return nil, false
	}
	go func() {
		c.lock.Lock()
		if timer := c.mTTL[key]; timer != nil {
			timer.Reset(c.ttl)
		}
		c.lock.Unlock()
	}()

	return val, true
}

// delete value and clear timer
func (c *cache[T]) ClearValue(key T) {
	c.lock.Lock()
	delete(c.m, key)
	if timer, ok := c.mTTL[key]; ok {
		timer.Stop()
		delete(c.mTTL, key)
	}
	c.lock.Unlock()
}
