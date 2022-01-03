package store

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

type store struct {
	defaultExpiration time.Duration
	items             map[string]string
	mu                sync.RWMutex
}
type Store struct {
	*store
}

func (s *store) Set(key string, x string, d time.Duration) error {
	if d == DefaultExpiration {
		d = s.defaultExpiration
	}
	s.mu.Lock()
	s.items[key] = x
	s.mu.Unlock()
	return nil
}
func (s *store) set(key string, x string, d time.Duration) {
	if d == DefaultExpiration {
		d = s.defaultExpiration
	}
	s.items[key] = x
}

func (s *store) Get(key string) (string, bool) {
	s.mu.RLock()
	item, found := s.items[key]
	if !found {
		s.mu.RUnlock()
		return "", false
	}
	s.mu.RUnlock()
	return item, true
}

func (s *store) get(key string) (string, bool) {
	item, found := s.items[key]
	if !found {
		return "", false
	}
	return item, true
}
func (s *store) Delete(k string) error {
	s.mu.Lock()
	s.delete(k)
	s.mu.Unlock()

	return nil
}
func (s *store) delete(k string) (string, bool) {
	if s != nil {
		if v, found := s.items[k]; found {
			delete(s.items, k)
			return v, true
		}
	}
	delete(s.items, k)
	return "", false
}
func (s *store) save(w io.Writer) (err error) {
	enc := gob.NewEncoder(w)
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, v := range s.items {
		gob.Register(v)
	}
	err = enc.Encode(&s.items)
	return
}
func (s *store) SaveFile(filename string) error {
	fp, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = s.save(fp)
	if err != nil {
		fp.Close()
		return err
	}
	return fp.Close()
}
func (s *store) load(r io.Reader) error {
	dec := gob.NewDecoder(r)
	items := make(map[string]string)
	err := dec.Decode(&items)
	if err == nil {
		s.mu.Lock()
		defer s.mu.Unlock()
		for i, j := range items {
			s.items[i] = j
		}

	}
	return err
}

func (s *store) LoadFile(filename string) error {
	fp, err := os.Open(filename)
	if err != nil {
		return err
	}
	err = s.load(fp)
	if err != nil {
		fp.Close()
		return err
	}
	return fp.Close()
}
func (s *store) Items() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	m := make(map[string]string, len(s.items))
	for k, v := range s.items {
		m[k] = v
	}
	return m
}
func (s *store) ItemCount() int {
	s.mu.RLock()
	n := len(s.items)
	s.mu.RUnlock()
	return n
}
func (s *store) Flush() {
	s.mu.Lock()
	s.items = make(map[string]string)
	s.mu.Unlock()
}
func newStore(de time.Duration, m map[string]string) *Store {
	if de == 0 {
		de = -1
	}
	s := &store{
		defaultExpiration: de,
		items:             m,
	}
	S := Store{s}
	return &S
}
func New(defaultExpiration, cleanupInterval time.Duration) *Store {
	items := make(map[string]string)
	return newStore(defaultExpiration, items)
}
