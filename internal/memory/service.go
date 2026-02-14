package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/webeauford/accounts/pkg/accounts"
)

var ErrNotFound = errors.New("user not found")

type Service struct {
	mu    sync.RWMutex
	users map[string]accounts.User
}

func New() *Service {
	return &Service{
		users: make(map[string]accounts.User),
	}
}

func (s *Service) GetUser(ctx context.Context, id string) (accounts.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	u, ok := s.users[id]
	if !ok {
		return accounts.User{}, ErrNotFound
	}
	return u, nil
}

// Helper you can delete later; useful for tests/demo.
func (s *Service) PutUser(u accounts.User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users[u.ID] = u
}
