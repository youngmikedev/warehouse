package service

import (
	"sync"
	"time"
)

const defaultLen = 20

type ShortSession struct {
	UserID      int
	SessionID   int
	AccessToken string
	UpdatedAt   time.Time
}

type CacheWithMaxLen struct {
	mu     sync.Mutex
	m      map[string]*cacheWithMaxLenItem
	maxLen int
}

// cacheWithMaxLenItem for token access without database
type cacheWithMaxLenItem struct {
	Item      mapSession
	CreatedAt time.Time
}

type mapSession struct {
	UID       int
	SID       int
	UpdatedAt time.Time
}

func NewCacheWithMaxLen(len int) *CacheWithMaxLen {
	if len == 0 {
		len = defaultLen
	}
	return &CacheWithMaxLen{
		m:      make(map[string]*cacheWithMaxLenItem, len),
		maxLen: len,
	}
}

func (c *CacheWithMaxLen) Set(key string, value ShortSession) {
	if len(c.m) >= c.maxLen {
		c.Pop()
	}

	c.mu.Lock()
	c.m[key] = &cacheWithMaxLenItem{
		Item: mapSession{
			UID:       value.UserID,
			SID:       value.SessionID,
			UpdatedAt: value.UpdatedAt,
		},
		CreatedAt: time.Now(),
	}
	c.mu.Unlock()
}

func (c *CacheWithMaxLen) Get(key string) (ShortSession, bool) {
	c.mu.Lock()
	value, ok := c.m[key]
	c.mu.Unlock()
	if !ok {
		return ShortSession{}, false
	}
	return ShortSession{
		UserID:      value.Item.UID,
		SessionID:   value.Item.SID,
		AccessToken: key,
	}, true
}

func (c *CacheWithMaxLen) Delete(key string) {
	c.mu.Lock()
	delete(c.m, key)
	c.mu.Unlock()
}

func (c *CacheWithMaxLen) Pop() {
	var minCreatedAt = time.Now()
	var minKey string
	c.mu.Lock()
	for k, v := range c.m {
		if v.CreatedAt.Before(minCreatedAt) {
			minKey = k
			minCreatedAt = v.CreatedAt
		}
	}
	delete(c.m, minKey)
	c.mu.Unlock()
}
