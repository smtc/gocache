package gocache

import (
	"encoding/json"
	"time"

	"github.com/pmylund/go-cache"
)

type LocalCache struct {
	*cache.Cache
}

func getLocalCache() *LocalCache {
	return &LocalCache{cache.New(expiration, cleanupInterval)}
}

func (c *LocalCache) Set(k string, x interface{}, d time.Duration) {
	c.Cache.Set(k, x, d)
}

func (c *LocalCache) Get(k string) (interface{}, bool) {
	return c.Cache.Get(k)
}

func (c *LocalCache) Store(k string, x interface{}, d time.Duration) {
	b, err := json.Marshal(x)
	if err == nil {
		c.Cache.Set(k, b, d)
	}
}

func (c *LocalCache) Retrieve(k string, v interface{}) (bool, error) {
	x, suc := c.Cache.Get(k)
	if !suc {
		return false, nil
	}
	err := json.Unmarshal(x.([]byte), v)
	return suc, err
}

func (c *LocalCache) Add(k string, x interface{}, d time.Duration) error {
	return c.Cache.Add(k, x, d)
}

func (c *LocalCache) Replace(k string, x interface{}, d time.Duration) error {
	return c.Cache.Replace(k, x, d)
}

func (c *LocalCache) Delete(k string) {
	c.Cache.Delete(k)
}
