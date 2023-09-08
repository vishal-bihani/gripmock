// Package rest provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.14.0 DO NOT EDIT.
package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
	codes "google.golang.org/grpc/codes"
)

// ID defines model for ID.
type ID = openapi_types.UUID

// ListID defines model for ListID.
type ListID = []ID

// MessageOK defines model for MessageOK.
type MessageOK struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

// SearchRequest defines model for SearchRequest.
type SearchRequest struct {
	Data    interface{} `json:"data"`
	Id      *ID         `json:"id,omitempty"`
	Method  string      `json:"method"`
	Service string      `json:"service"`
}

// SearchResponse defines model for SearchResponse.
type SearchResponse struct {
	Code  *codes.Code `json:"code,omitempty"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// Stub defines model for Stub.
type Stub struct {
	Id      *ID        `json:"id,omitempty"`
	Input   StubInput  `json:"input"`
	Method  string     `json:"method"`
	Output  StubOutput `json:"output"`
	Service string     `json:"service"`
}

// StubInput defines model for StubInput.
type StubInput struct {
	Contains map[string]interface{} `json:"contains,omitempty"`
	Equals   map[string]interface{} `json:"equals,omitempty"`
	Matches  map[string]interface{} `json:"matches,omitempty"`
}

// StubList defines model for StubList.
type StubList = []Stub

// StubOutput defines model for StubOutput.
type StubOutput struct {
	Code  *codes.Code            `json:"code,omitempty"`
	Data  map[string]interface{} `json:"data"`
	Error string                 `json:"error"`
}

// AddStubJSONBody defines parameters for AddStub.
type AddStubJSONBody struct {
	union json.RawMessage
}

// AddStubJSONRequestBody defines body for AddStub for application/json ContentType.
type AddStubJSONRequestBody AddStubJSONBody

// BatchStubsDeleteJSONRequestBody defines body for BatchStubsDelete for application/json ContentType.
type BatchStubsDeleteJSONRequestBody = ListID

// SearchStubsJSONRequestBody defines body for SearchStubs for application/json ContentType.
type SearchStubsJSONRequestBody = SearchRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Liveness check
	// (GET /health/liveness)
	Liveness(w http.ResponseWriter, r *http.Request)
	// Readiness check
	// (GET /health/readiness)
	Readiness(w http.ResponseWriter, r *http.Request)
	// Remove all stubs
	// (DELETE /stubs)
	PurgeStubs(w http.ResponseWriter, r *http.Request)
	// Getting a list of stubs
	// (GET /stubs)
	ListStubs(w http.ResponseWriter, r *http.Request)
	// Add a new stub to the store
	// (POST /stubs)
	AddStub(w http.ResponseWriter, r *http.Request)
	// Deletes a pack by IDs
	// (POST /stubs/batchDelete)
	BatchStubsDelete(w http.ResponseWriter, r *http.Request)
	// Stub storage search
	// (POST /stubs/search)
	SearchStubs(w http.ResponseWriter, r *http.Request)
	// Getting a list of unused stubs
	// (GET /stubs/unused)
	ListUnusedStubs(w http.ResponseWriter, r *http.Request)
	// Deletes stub by ID
	// (DELETE /stubs/{uuid})
	DeleteStubByID(w http.ResponseWriter, r *http.Request, uuid ID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// Liveness operation middleware
func (siw *ServerInterfaceWrapper) Liveness(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Liveness(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Readiness operation middleware
func (siw *ServerInterfaceWrapper) Readiness(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Readiness(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PurgeStubs operation middleware
func (siw *ServerInterfaceWrapper) PurgeStubs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PurgeStubs(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListStubs operation middleware
func (siw *ServerInterfaceWrapper) ListStubs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListStubs(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AddStub operation middleware
func (siw *ServerInterfaceWrapper) AddStub(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddStub(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// BatchStubsDelete operation middleware
func (siw *ServerInterfaceWrapper) BatchStubsDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.BatchStubsDelete(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// SearchStubs operation middleware
func (siw *ServerInterfaceWrapper) SearchStubs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SearchStubs(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListUnusedStubs operation middleware
func (siw *ServerInterfaceWrapper) ListUnusedStubs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListUnusedStubs(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteStubByID operation middleware
func (siw *ServerInterfaceWrapper) DeleteStubByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "uuid" -------------
	var uuid ID

	err = runtime.BindStyledParameter("simple", false, "uuid", mux.Vars(r)["uuid"], &uuid)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "uuid", Err: err})
		return
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteStubByID(w, r, uuid)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{})
}

type GorillaServerOptions struct {
	BaseURL          string
	BaseRouter       *mux.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r *mux.Router) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r *mux.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options GorillaServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = mux.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.HandleFunc(options.BaseURL+"/health/liveness", wrapper.Liveness).Methods("GET")

	r.HandleFunc(options.BaseURL+"/health/readiness", wrapper.Readiness).Methods("GET")

	r.HandleFunc(options.BaseURL+"/stubs", wrapper.PurgeStubs).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/stubs", wrapper.ListStubs).Methods("GET")

	r.HandleFunc(options.BaseURL+"/stubs", wrapper.AddStub).Methods("POST")

	r.HandleFunc(options.BaseURL+"/stubs/batchDelete", wrapper.BatchStubsDelete).Methods("POST")

	r.HandleFunc(options.BaseURL+"/stubs/search", wrapper.SearchStubs).Methods("POST")

	r.HandleFunc(options.BaseURL+"/stubs/unused", wrapper.ListUnusedStubs).Methods("GET")

	r.HandleFunc(options.BaseURL+"/stubs/{uuid}", wrapper.DeleteStubByID).Methods("DELETE")

	return r
}
