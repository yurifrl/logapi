package store

import (
	"sync"
	"time"

	"github.com/k0kubun/pp"
)

type Store struct {
	sync.RWMutex
	data map[string]int
	last time.Time
}

func Create() *Store {
	db := &Store{
		data: make(map[string]int),
	}
	return db
}

func (s *Store) Bump(key string, time time.Time) error {
	s.Lock()
	s.data[key] = s.data[key] + 1
	pp.Println("=========")
	pp.Println(time)
	pp.Println(s.last)
	pp.Println("=========")
	if time.After(s.last) {
		s.last = time
	}
	defer s.Unlock()
	return nil
}

func (s *Store) Last() (time.Time, error) {
	return s.last, nil
}

func (s *Store) GetAll() (map[string]int, error) {
	s.RLock()
	data := s.data
	defer s.RUnlock()
	return data, nil
}
