package member

import (
	"sort"
	"sync"
	"time"
)

type Store struct {
	mu      sync.RWMutex
	members map[int]Member
	nextID  int
}

func NewStore() *Store {
	return &Store{
		members: make(map[int]Member),
		nextID:  1,
	}
}

func (s *Store) Create(m Member) Member {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()

	m.ID = s.nextID
	m.Status = "active"
	m.CreatedAt = now
	m.UpdatedAt = now

	if m.Role == "" {
		m.Role = "Member"
	}

	s.members[m.ID] = m
	s.nextID++

	return m
}

func (s *Store) Get(id int) (Member, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	m, exists := s.members[id]
	return m, exists
}

func (s *Store) GetAll() []Member {
	s.mu.RLock()
	defer s.mu.RUnlock()

	members := make([]Member, 0, len(s.members))

	for _, m := range s.members {
		members = append(members, m)
	}

	sort.Slice(members, func(i, j int) bool {
		return members[i].ID < members[j].ID
	})

	return members
}
