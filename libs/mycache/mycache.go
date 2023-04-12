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
	c.mTTL[key] = time.NewTimer(c.ttl)
	c.lock.Unlock()
	go func() {
		<-c.mTTL[key].C
		c.ClearValue(key)
	}()
}

//private method to retrieve value
func (c *cache[T]) getValue(key T) (*any, bool) {
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

//helper converter using generics
func castValue[T any](val any) (*T, bool) {
	res, ok := val.(T)
	if !ok {
		return nil, false
	}
	return &res, true
}

func (c *cache[T]) ClearValue(key T) {
	c.lock.Lock()
	delete(c.m, key)
	if timer, ok := c.mTTL[key]; ok {
		timer.Stop()
		delete(c.mTTL, key)
	}
	c.lock.Unlock()
}

func (c *cache[T]) GetInt32(key T) (*int32, bool) {
	val, ok := c.getValue(key)
	if !ok {
		return nil, false
	}
	return castValue[int32](val)
}

func (c *cache[T]) GetInt64(key T) (*int64, bool) {
	val, ok := c.getValue(key)
	if !ok {
		return nil, false
	}
	return castValue[int64](val)
}

func (c *cache[T]) GetUint32(key T) (*uint32, bool) {
	val, ok := c.getValue(key)
	if !ok {
		return nil, false
	}
	return castValue[uint32](val)
}

func (c *cache[T]) GetUint64(key T) (*uint64, bool) {
	val, ok := c.getValue(key)
	if !ok {
		return nil, false
	}
	return castValue[uint64](val)
}

func (c *cache[T]) GetString(key T) (*string, bool) {
	val, ok := c.getValue(key)
	if !ok {
		return nil, false
	}
	return castValue[string](val)
}

func (c *cache[T]) GetRawValue(key T) (*any, bool) {
	val, ok := c.getValue(key)
	if !ok {
		return nil, false
	}
	return val, true
}
