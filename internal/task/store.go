package task

import (
	"sort"
	"sync"
	"time"
)

type Store struct {
	mu     sync.RWMutex
	tasks  map[int]Task
	nextID int
}

func NewStore() *Store {
	return &Store{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

func (s *Store) Create(t Task) Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()

	t.ID = s.nextID
	t.Status = "queued"
	t.Attempts = 0
	t.Error = ""
	t.CreatedAt = now
	t.UpdatedAt = now

	if t.MaxRetries <= 0 {
		t.MaxRetries = 3
	}

	s.tasks[t.ID] = t
	s.nextID++

	return t
}

func (s *Store) Save(t Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t.UpdatedAt = time.Now()
	s.tasks[t.ID] = t
}

func (s *Store) Get(id int) (Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	t, exists := s.tasks[id]
	return t, exists
}

func (s *Store) GetAll() []Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]Task, 0, len(s.tasks))

	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	return tasks
}
