package store

import (
	"fmt"
	"sync"
)

type MockStore struct {
	sync.RWMutex
	data map[string]map[string]int
}

func Create() *MockStore {
	db := &MockStore{
		data: make(map[string]map[string]int),
	}
	return db
}

func (s *MockStore) Bump(keyPath []string) error {
	s.Lock()
	if len(keyPath) != 2 {
		// @TODO: use a real database to index this
		return fmt.Errorf("to many items")
	}

	instance := keyPath[0]
	app := keyPath[1]
	if _, ok := s.data[instance]; !ok {
		s.data[instance] = make(map[string]int)
	}

	s.data[instance][app] = s.data[instance][app] + 1

	defer s.Unlock()
	return nil
}

func (s *MockStore) GetAll() (map[string]string, error) {
	s.RLock()

	data := make(map[string]string)

	totalErrorCount := 0
	for instance, apps := range s.data {
		errCount := 0
		for app, count := range apps {
			data[app] = fmt.Sprintf("%v", count)
			errCount += count
			totalErrorCount += count
		}
		data[instance] = fmt.Sprintf("%v/%v", errCount, totalErrorCount)
	}

	defer s.RUnlock()
	return data, nil
}
