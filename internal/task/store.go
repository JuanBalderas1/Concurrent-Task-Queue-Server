package task

import "sync"

type Store struct {
	mu    sync.RWMutex
	tasks map[int]Task
}

func NewStore() *Store {
	return &Store{
		tasks: make(map[int]Task),
	}
}

func (s *Store) Save(t Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks[t.ID] = t
}

func (s *Store) Get(id int) (Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	t, exists := s.tasks[id]
	return t, exists
}
