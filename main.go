package gocache

import "time"

type ICache interface {
	Set(k string, x interface{}, d time.Duration)
	Get(k string) (interface{}, bool)
	Add(k string, x interface{}, d time.Duration) error
	Store(k string, x interface{}, d time.Duration)
	Retrieve(k string, v interface{}) (bool, error)
	Replace(k string, x interface{}, d time.Duration) error
	Delete(string)
}

var (
	expiration      = 20 * time.Minute
	cleanupInterval = 60 * time.Second
	c               ICache
)

func SetExpiration(expiration, cleanupInterval time.Duration) {
	expiration = expiration
	cleanupInterval = cleanupInterval
}

func GetCache() ICache {
	if c == nil {
		c = getLocalCache()
	}
	return c
}
