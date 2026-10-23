// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"context"
	"sync"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
)

// MemoryTokenStore stores tokens in memory. It is safe for concurrent use and
// is the default store for a provider.
type MemoryTokenStore struct {
	mu     sync.RWMutex
	tokens map[TokenKey]Token
}

// NewMemoryTokenStore creates an empty in-memory token store.
func NewMemoryTokenStore() *MemoryTokenStore {
	return &MemoryTokenStore{
		tokens: make(map[TokenKey]Token),
	}
}

// Load returns a copy of the token for key, or (nil, nil) when no token exists.
func (s *MemoryTokenStore) Load(ctx context.Context, key TokenKey) (*Token, error) {
	if ctx == nil {
		return nil, snerrors.ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if s == nil {
		return nil, snerrors.ErrNilStore
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	token, ok := s.tokens[key]
	if !ok {
		return nil, nil
	}

	copied := token
	return &copied, nil
}

// Save stores a copy of token for key.
func (s *MemoryTokenStore) Save(ctx context.Context, key TokenKey, token *Token) error {
	if ctx == nil {
		return snerrors.ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if s == nil {
		return snerrors.ErrNilStore
	}
	if token == nil {
		return snerrors.ErrNilModel
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.tokens == nil {
		s.tokens = make(map[TokenKey]Token)
	}
	s.tokens[key] = *token
	return nil
}

// Delete removes the token for key. Deleting a missing key succeeds.
func (s *MemoryTokenStore) Delete(ctx context.Context, key TokenKey) error {
	if ctx == nil {
		return snerrors.ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if s == nil {
		return snerrors.ErrNilStore
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.tokens, key)
	return nil
}
