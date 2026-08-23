package support

import (
	"context"

	sdk "github.com/michaeldcanady/servicenow-sdk-go/v2"
)

type worldKey struct{}

// World holds all per-scenario state. A fresh World is created for every scenario.
type World struct {
	Client   *sdk.ServiceNowServiceClient
	Err      error
	AuthErr  error
	Response interface{}

	// Resource tracking for cleanup
	LastSysID  string
	CreatedIDs []string

	// Auth-specific state
	CachedToken   string
	FetchedToken  string
	RevocationErr error
	TokenExpired  bool

	// Pagination state
	PageCount int
	PageSize  int
}

// NewWorld creates a fresh World for a new scenario.
func NewWorld() *World {
	return &World{
		CreatedIDs: make([]string, 0),
	}
}

// WithWorld stores a World in the context.
func WithWorld(ctx context.Context, w *World) context.Context {
	return context.WithValue(ctx, worldKey{}, w)
}

// WorldFrom retrieves the World from the context.
func WorldFrom(ctx context.Context) *World {
	w, _ := ctx.Value(worldKey{}).(*World)
	return w
}

// TrackCreation records a sys_id for cleanup after the scenario.
func (w *World) TrackCreation(sysID string) {
	if sysID != "" {
		w.CreatedIDs = append(w.CreatedIDs, sysID)
	}
}

// Reset clears per-step state while preserving the client.
func (w *World) Reset() {
	w.Err = nil
	w.AuthErr = nil
	w.Response = nil
	w.LastSysID = ""
}
