package repository

import (
	"net"
	"sync"
)

type CachedResults struct {
	mu sync.Locker
	v  map[string]map[string]interface{}
}

func NewCachedResults() *CachedResults {
	var mutex sync.Locker = &sync.Mutex{}

	return &CachedResults{mu: mutex, v: make(map[string]map[string]interface{})}
}

func (c *CachedResults) Set(database string, ip net.IP, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[database] = make(map[string]interface{})
	c.v[database][ip.String()] = value
}

func (c *CachedResults) Value(database string, ip net.IP) (interface{}, bool) {
	val, ok := c.v[database][ip.String()]

	return val, ok
}
