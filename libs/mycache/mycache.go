package mycache

import (
	"sync"
	"time"
)

type cache struct {
	//simple hash map with any value
	m map[string]any
	//map for ttl timers
	mTTL map[string]*time.Timer
	//for thread-safety
	lock sync.RWMutex

	ttl time.Duration
}

func NewMyCache(ttl time.Duration) *cache {
	return &cache{
		m:    make(map[string]any),
		mTTL: make(map[string]*time.Timer),
		ttl:  ttl,
	}
}

func (c *cache) SetValue(key string, value any) {
	c.lock.Lock()
	c.m[key] = value
	c.mTTL[key] = time.NewTimer(c.ttl)
	c.lock.Unlock()
	go func() {
		<-c.mTTL[key].C
		c.clearValue(key)
	}()
}

//private method to retrieve value
func (c *cache) getValue(key string) (any, bool) {
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

func (c *cache) clearValue(key string) {
	c.lock.Lock()
	delete(c.m, key)
	if timer, ok := c.mTTL[key]; ok {
		timer.Stop()
		delete(c.mTTL, key)
	}
	c.lock.Unlock()
}

func (c *cache) GetInt32(key string) (*int32, bool) {
	val, ok := c.getValue(key)
	if !ok {
		return nil, false
	}
	return castValue[int32](val)
}

func (c *cache) GetInt64(key string) (*int64, bool) {
	val, ok := c.getValue(key)
	if !ok {
		return nil, false
	}
	return castValue[int64](val)
}

func (c *cache) GetUint32(key string) (*uint32, bool) {
	val, ok := c.getValue(key)
	if !ok {
		return nil, false
	}
	return castValue[uint32](val)
}

func (c *cache) GetUint64(key string) (*uint64, bool) {
	val, ok := c.getValue(key)
	if !ok {
		return nil, false
	}
	return castValue[uint64](val)
}

func (c *cache) GetString(key string) (*string, bool) {
	val, ok := c.getValue(key)
	if !ok {
		return nil, false
	}
	return castValue[string](val)
}
