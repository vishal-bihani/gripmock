package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/gripmock/stuber"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/bavix/features"
	"github.com/bavix/gripmock/internal/domain/rest"
	"github.com/bavix/gripmock/pkg/grpcreflector"
	"github.com/bavix/gripmock/pkg/jsondecoder"
	"github.com/bavix/gripmock/pkg/yaml2json"
)

var (
	ErrServiceIsMissing = errors.New("service name is missing")
	ErrMethodIsMissing  = errors.New("method name is missing")
)

type RestServer struct {
	ok        atomic.Bool
	stuber    *stuber.Budgerigar
	convertor *yaml2json.Convertor
	caser     cases.Caser
	reflector *grpcreflector.GReflector
}

var _ rest.ServerInterface = &RestServer{}

func NewRestServer(path string, reflector *grpcreflector.GReflector) (*RestServer, error) {
	server := &RestServer{
		stuber:    stuber.NewBudgerigar(features.New(stuber.MethodTitle)),
		convertor: yaml2json.New(),
		caser:     cases.Title(language.English, cases.NoLower),
		reflector: reflector,
	}

	if path != "" {
		server.readStubs(path) // TODO: someday you will need to rewrite this code
		server.ok.Store(true)
	}

	return server, nil
}

func (h *RestServer) ServicesList(w http.ResponseWriter, r *http.Request) {
	services, err := h.reflector.Services(r.Context())
	if err != nil {
		return
	}

	results := make([]rest.Service, len(services))

	for i, service := range services {
		serviceMethods, err := h.reflector.Methods(r.Context(), service.ID)
		if err != nil {
			continue
		}

		restMethods := make([]rest.Method, len(serviceMethods))
		for i, serviceMethod := range serviceMethods {
			restMethods[i] = rest.Method{
				Id:   serviceMethod.ID,
				Name: serviceMethod.Name,
			}
		}

		results[i] = rest.Service{
			Id:      service.ID,
			Name:    service.Name,
			Package: service.Package,
			Methods: restMethods,
		}
	}

	if err := json.NewEncoder(w).Encode(results); err != nil {
		h.responseError(err, w)
	}
}

func (h *RestServer) ServiceMethodsList(w http.ResponseWriter, r *http.Request, serviceID string) {
	methods, err := h.reflector.Methods(r.Context(), serviceID)
	if err != nil {
		return
	}

	results := make([]rest.Method, len(methods))
	for i, method := range methods {
		results[i] = rest.Method{
			Id:   method.ID,
			Name: method.Name,
		}
	}

	if err := json.NewEncoder(w).Encode(results); err != nil {
		h.responseError(err, w)
	}
}

func (h *RestServer) liveness(w http.ResponseWriter) {
	if err := json.NewEncoder(w).Encode(rest.MessageOK{Message: "ok", Time: time.Now()}); err != nil {
		h.responseError(err, w)
	}
}

func (h *RestServer) Readiness(w http.ResponseWriter, _ *http.Request) {
	if !h.ok.Load() {
		w.WriteHeader(http.StatusServiceUnavailable)

		return
	}

	h.liveness(w)
}

func (h *RestServer) Liveness(w http.ResponseWriter, _ *http.Request) {
	h.liveness(w)
}

func (h *RestServer) AddStub(w http.ResponseWriter, r *http.Request) {
	byt, err := io.ReadAll(r.Body)
	if err != nil {
		h.responseError(err, w)

		return
	}

	defer r.Body.Close()

	var inputs []*stuber.Stub

	if err := jsondecoder.UnmarshalSlice(byt, &inputs); err != nil {
		h.responseError(err, w)

		return
	}

	for _, stub := range inputs {
		if err := validateStub(stub); err != nil {
			h.responseError(err, w)

			return
		}
	}

	if err := json.NewEncoder(w).Encode(h.stuber.PutMany(inputs...)); err != nil {
		h.responseError(err, w)

		return
	}
}

func (h *RestServer) DeleteStubByID(w http.ResponseWriter, _ *http.Request, uuid rest.ID) {
	h.stuber.DeleteByID(uuid)

	w.WriteHeader(http.StatusNoContent)
}

func (h *RestServer) BatchStubsDelete(w http.ResponseWriter, r *http.Request) {
	byt, err := io.ReadAll(r.Body)
	if err != nil {
		h.responseError(err, w)

		return
	}

	defer r.Body.Close()

	var inputs []uuid.UUID

	if err := jsondecoder.UnmarshalSlice(byt, &inputs); err != nil {
		h.responseError(err, w)

		return
	}

	if len(inputs) > 0 {
		h.stuber.DeleteByID(inputs...)
	}
}

func (h *RestServer) ListUsedStubs(w http.ResponseWriter, _ *http.Request) {
	if err := json.NewEncoder(w).Encode(h.stuber.Used()); err != nil {
		h.responseError(err, w)
	}
}

func (h *RestServer) ListUnusedStubs(w http.ResponseWriter, _ *http.Request) {
	if err := json.NewEncoder(w).Encode(h.stuber.Unused()); err != nil {
		h.responseError(err, w)
	}
}

func (h *RestServer) ListStubs(w http.ResponseWriter, _ *http.Request) {
	if err := json.NewEncoder(w).Encode(h.stuber.All()); err != nil {
		h.responseError(err, w)
	}
}

func (h *RestServer) PurgeStubs(w http.ResponseWriter, _ *http.Request) {
	h.stuber.Clear()

	w.WriteHeader(http.StatusNoContent)
}

func (h *RestServer) SearchStubs(w http.ResponseWriter, r *http.Request) {
	query, err := stuber.NewQuery(r)
	if err != nil {
		h.responseError(err, w)

		return
	}

	defer r.Body.Close()

	result, err := h.stuber.FindByQuery(query)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		h.writeResponseError(err, w)

		return
	}

	if result.Found() == nil {
		w.WriteHeader(http.StatusNotFound)
		h.writeResponseError(stubNotFoundError2(query, result), w)

		return
	}

	if err := json.NewEncoder(w).Encode(result.Found().Output); err != nil {
		h.responseError(err, w)
	}
}

func (h *RestServer) FindByID(w http.ResponseWriter, _ *http.Request, uuid rest.ID) {
	stub := h.stuber.FindByID(uuid)
	if stub == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	if err := json.NewEncoder(w).Encode(stub); err != nil {
		h.responseError(err, w)
	}
}

func (h *RestServer) responseError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)

	h.writeResponseError(err, w)
}

func (h *RestServer) writeResponseError(err error, w http.ResponseWriter) {
	if err := json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	}); err != nil {
		h.responseError(err, w)
	}
}

// readStubs reads all the stubs from the given directory and its subdirectories,
// and adds them to the server's stub store.
// The stub files can be in yaml or json format.
// If a file is in yaml format, it will be converted to json format.
func (h *RestServer) readStubs(pathDir string) {
	files, err := os.ReadDir(pathDir)
	if err != nil {
		log.Printf("can't read stubs from %s: %v", pathDir, err)

		return
	}

	for _, file := range files {
		// If the file is a directory, recursively read its stubs.
		if file.IsDir() {
			h.readStubs(path.Join(pathDir, file.Name()))

			continue
		}

		// Read the stub file and add it to the server's stub store.
		stubs, err := h.readStub(path.Join(pathDir, file.Name()))
		if err != nil {
			log.Printf("cant read stubs from %s: %v", file.Name(), err)

			continue
		}

		h.stuber.PutMany(stubs...)
	}
}

// readStub reads a stub file and returns a slice of stubs.
// The stub file can be in yaml or json format.
// If the file is in yaml format, it will be converted to json format.
func (h *RestServer) readStub(path string) ([]*stuber.Stub, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		file, err = h.convertor.Execute(path, file)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal file %s: %w", path, err)
		}
	}

	var stubs []*stuber.Stub
	if err := jsondecoder.UnmarshalSlice(file, &stubs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file %s: %v %w", path, string(file), err)
	}

	return stubs, nil
}

// validateStub validates if the stub is valid or not.
func validateStub(stub *stuber.Stub) error {
	if stub.Service == "" {
		return ErrServiceIsMissing
	}

	if stub.Method == "" {
		return ErrMethodIsMissing
	}

	if stub.Input.Contains == nil && stub.Input.Equals == nil && stub.Input.Matches == nil {
		return errors.New("input cannot be empty") //nolint:err113
	}

	if stub.Output.Error == "" && stub.Output.Data == nil && stub.Output.Code == nil {
		return errors.New("output cannot be empty") //nolint:err113
	}

	return nil
}
