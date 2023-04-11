package mycache

import "sync"

type cache struct {
	//simple hash map with any value
	m map[string]interface{}
	//for thread-safety
	lock sync.RWMutex
}

func NewMyCache() *cache {
	return &cache{
		m: make(map[string]interface{}),
	}
}

func (c *cache) SetValue(key string, value interface{}) {
	c.lock.Lock()
	c.m[key] = value
	c.lock.Unlock()
}

func (c *cache) getValue(key string) (interface{}, bool) {
	c.lock.RLock()
	val, ok := c.m[key]
	c.lock.RUnlock()
	if !ok {
		return nil, false
	}
	return val, true
}

func (c *cache) GetInt32(key string) (*int32, bool) {
	var val interface{}
	var ok bool
	if val, ok = c.getValue(key); !ok {
		return nil, false
	}

	var res int32
	if res, ok = val.(int32); !ok {
		return nil, false
	}
	return &res, true
}

func (c *cache) GetInt64(key string) (*int64, bool) {
	var val interface{}
	var ok bool
	if val, ok = c.getValue(key); !ok {
		return nil, false
	}

	var res int64
	if res, ok = val.(int64); !ok {
		return nil, false
	}
	return &res, true
}

func (c *cache) GetUint32(key string) (*uint32, bool) {
	var val interface{}
	var ok bool
	if val, ok = c.getValue(key); !ok {
		return nil, false
	}

	var res uint32
	if res, ok = val.(uint32); !ok {
		return nil, false
	}
	return &res, true
}

func (c *cache) GetUint64(key string) (*uint64, bool) {
	var val interface{}
	var ok bool
	if val, ok = c.getValue(key); !ok {
		return nil, false
	}

	var res uint64
	if res, ok = val.(uint64); !ok {
		return nil, false
	}
	return &res, true
}

func (c *cache) GetString(key string) (*string, bool) {
	var val interface{}
	var ok bool
	if val, ok = c.getValue(key); !ok {
		return nil, false
	}

	var res string
	if res, ok = val.(string); !ok {
		return nil, false
	}
	return &res, true
}
