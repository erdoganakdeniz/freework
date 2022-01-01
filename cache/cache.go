package cache

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

const (
	NoExpiration      time.Duration = -1
	DefaultExpiration time.Duration = 0
)

type cache struct {
	defaultExpiration time.Duration
	items             map[string]string
	mu                sync.RWMutex
}
type Cache struct {
	*cache
}

func (c *cache) Set(key string, x string, d time.Duration) error {
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	c.mu.Lock()
	c.items[key] = x
	c.mu.Unlock()
	return nil
}
func (c *cache) set(key string, x string, d time.Duration) {
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	c.items[key] = x
}

func (c *cache) Get(key string) (string, bool) {
	c.mu.RLock()
	item, found := c.items[key]
	if !found {
		c.mu.RUnlock()
		return "", false
	}
	c.mu.RUnlock()
	return item, true
}

func (c *cache) get(key string) (string, bool) {
	item, found := c.items[key]
	if !found {
		return "", false
	}
	return item, true
}
func (c *cache) Delete(k string) error {
	c.mu.Lock()
	c.delete(k)
	c.mu.Unlock()

	return nil
}
func (c *cache) delete(k string) (v string, a bool) {
	if c != nil {
		if v, found := c.items[k]; found {
			delete(c.items, k)
			return v, true
		}
	}
	delete(c.items, k)
	return "", false
}
func (c *cache) save(w io.Writer) (err error) {
	enc := gob.NewEncoder(w)
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, v := range c.items {
		gob.Register(v)
	}
	err = enc.Encode(&c.items)
	return
}
func (c *cache) SaveFile(filename string) error {
	fp, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = c.save(fp)
	if err != nil {
		fp.Close()
		return err
	}
	return fp.Close()
}
func (c *cache) load(r io.Reader) error {
	dec := gob.NewDecoder(r)
	items := make(map[string]string)
	err := dec.Decode(&items)
	if err == nil {
		c.mu.Lock()
		defer c.mu.Unlock()
		//for k, v := range items {
		//ov, found := c.items[k]
		//if !found || ov.Expired() {
		//c.items[k] = v
		//}
		//}
	}
	return err
}

func (c *cache) LoadFile(filename string) error {
	fp, err := os.Open(filename)
	if err != nil {
		return err
	}
	err = c.load(fp)
	if err != nil {
		fp.Close()
		return err
	}
	return fp.Close()
}
func (c *cache) Items() map[string]string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	m := make(map[string]string, len(c.items))
	//now := time.Now().UnixNano()
	for k, v := range c.items {

		//if v.Expiration > 0 {
		//	if now > v.Expiration {
		//		continue
		//	}
		//}
		m[k] = v
	}
	return m
}
func (c *cache) ItemCount() int {
	c.mu.RLock()
	n := len(c.items)
	c.mu.RUnlock()
	return n
}
func (c *cache) Flush() {
	c.mu.Lock()
	c.items = make(map[string]string)
	c.mu.Unlock()
}
func newCache(de time.Duration, m map[string]string) *Cache {
	if de == 0 {
		de = -1
	}
	c := &cache{
		defaultExpiration: de,
		items:             m,
	}
	C := Cache{c}
	return &C
}
func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	items := make(map[string]string)
	return newCache(defaultExpiration, items)
}
