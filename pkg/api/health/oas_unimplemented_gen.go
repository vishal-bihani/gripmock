// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// Liveness implements liveness operation.
//
// The test says that the service is alive and yet.
//
// GET /liveness
func (UnimplementedHandler) Liveness(ctx context.Context) (r *MessageOK, _ error) {
	return r, ht.ErrNotImplemented
}

// Readiness implements readiness operation.
//
// The test indicates readiness to receive traffic.
//
// GET /readiness
func (UnimplementedHandler) Readiness(ctx context.Context) (r *MessageOK, _ error) {
	return r, ht.ErrNotImplemented
}
