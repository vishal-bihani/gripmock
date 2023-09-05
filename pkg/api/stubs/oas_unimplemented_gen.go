// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// AddStub implements addStub operation.
//
// Add a new stub to the store.
//
// POST /stubs
func (UnimplementedHandler) AddStub(ctx context.Context, req AddStubReq) (r AddStubOK, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteStubByID implements deleteStubByID operation.
//
// The method removes the stub by ID.
//
// DELETE /stubs/{uuid}
func (UnimplementedHandler) DeleteStubByID(ctx context.Context, params DeleteStubByIDParams) (r DeleteStubByIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListStubs implements listStubs operation.
//
// The list of stubs is required to view all added stubs.
//
// GET /stubs
func (UnimplementedHandler) ListStubs(ctx context.Context) (r StubList, _ error) {
	return r, ht.ErrNotImplemented
}

// PurgeStubs implements purgeStubs operation.
//
// Completely clears the stub storage.
//
// DELETE /stubs
func (UnimplementedHandler) PurgeStubs(ctx context.Context) error {
	return ht.ErrNotImplemented
}

// SearchStubs implements searchStubs operation.
//
// Performs a search for a stub by the given conditions.
//
// POST /stubs/search
func (UnimplementedHandler) SearchStubs(ctx context.Context, req *SearchRequest) (r *SearchResponse, _ error) {
	return r, ht.ErrNotImplemented
}
