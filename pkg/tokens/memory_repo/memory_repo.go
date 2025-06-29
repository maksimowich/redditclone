package tokens_memory_repo

import (
	"fmt"
	"sync"
	"time"
)

type TokensMemoryRepo struct {
	tokens map[string]time.Time
	mu     *sync.Mutex
}

func (s *TokensMemoryRepo) Add(
	token string,
	expires time.Time,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tokens[token] = expires
	return nil
}

func (s *TokensMemoryRepo) Delete(
	token string,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.tokens, token)
	return nil
}

func (s *TokensMemoryRepo) IsValid(
	token string,
) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	expires, exists := s.tokens[token]
	fmt.Printf("expires: %s, exists: %t\n", expires, exists)
	fmt.Printf("time.Now().Before(expires): %t\n", time.Now().Before(expires))
	return exists && time.Now().Before(expires), nil
}

func NewTokensMemoryRepo() *TokensMemoryRepo {
	return &TokensMemoryRepo{
		tokens: make(map[string]time.Time),
		mu:     &sync.Mutex{},
	}
}
