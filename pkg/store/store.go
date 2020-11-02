package store

import "sync"

type Store struct {
	sync.RWMutex
	data map[string]int
}

func Create() *Store {
	db := &Store{
		data: make(map[string]int),
	}
	return db
}

func (s *Store) Bump(key string) error {
	s.Lock()
	s.data[key] = s.data[key] + 1
	s.Unlock()
	return nil
}

func (s *Store) GetAll() (map[string]int, error) {
	s.RLock()
	data := s.data
	s.RUnlock()
	return data, nil
}
