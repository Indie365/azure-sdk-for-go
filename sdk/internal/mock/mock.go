// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package mock

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"

	"golang.org/x/net/http2"
)

const (
	mockReadError = "mock-read-error"
)

// Server is a wrapper around an httptest.Server.
// The serving of requests is not safe for concurrent use
// which is ok for right now as each test creates is own
// server and doesn't create additional go routines.
type Server struct {
	srv *httptest.Server

	// static is the static response, if this is not nil it's always returned.
	static *mockResponse

	// respLock synchronizes access to resp
	respLock *sync.RWMutex

	// resp is the queue of responses.  each response is taken from the front.
	resp []mockResponse

	// count tracks the number of requests that have been made.
	count int
}

func newServer() *Server {
	return &Server{respLock: &sync.RWMutex{}}
}

// NewServer creates a new Server object.
// The returned close func must be called when the server is no longer needed.
func NewServer() (*Server, func()) {
	s := newServer()
	s.srv = httptest.NewServer(http.HandlerFunc(s.serveHTTP))
	return s, func() { s.srv.Close() }
}

// NewTLSServer creates a new Server object.
// The returned close func must be called when the server is no longer needed.
func NewTLSServer() (*Server, func()) {
	s := newServer()
	s.srv = httptest.NewTLSServer(http.HandlerFunc(s.serveHTTP))
	return s, func() { s.srv.Close() }
}

// NewTLSHTTP2Server creates a new Server object with HTTP2 configuration.
// The returned close func must be called when the server is no longer needed.
// It will return nil for both the server and close func if it encountered an error
// when configuring HTTP2 for the new TLS server.
func NewTLSHTTP2Server() (*Server, func()) {
	s := newServer()
	s.srv = httptest.NewUnstartedServer(http.HandlerFunc(s.serveHTTP))
	if err := http2.ConfigureServer(s.srv.Config, new(http2.Server)); err != nil {
		return nil, nil
	}
	s.srv.TLS = s.srv.Config.TLSConfig
	s.srv.StartTLS()
	return s, func() { s.srv.Close() }
}

// ServerConfig returns a copy of the server config, this cannot be
// modified since Start or StartTLS has already been called on the server.
func (s *Server) ServerConfig() http.Server {
	return *s.srv.Config
}

// returns true if the next response is an error response
func (s *Server) isErrorResp() bool {
	// always favor static response
	if s.static != nil {
		return s.static.err != nil
	}
	s.respLock.RLock()
	defer s.respLock.RUnlock()
	if len(s.resp) > 0 {
		return s.resp[0].err != nil
	}
	panic("no more responses")
}

// returns the static response or the next response in the queue
func (s *Server) getResponse() mockResponse {
	// always favor static response
	if s.static != nil {
		return *s.static
	}
	if len(s.resp) > 0 {
		// pop off first response and return it
		s.respLock.Lock()
		defer s.respLock.Unlock()
		resp := s.resp[0]
		s.resp = s.resp[1:]
		return resp
	}
	panic("no more responses")
}

// URL returns the endpoint of the test server in URL format.
func (s *Server) URL() string {
	return s.srv.URL
}

// Do implements the azcore.Transport interface on Server.
// Calling this when the response queue is empty and no static
// response has been set will cause a panic.
func (s *Server) Do(req *http.Request) (*http.Response, error) {
	s.count++
	// error responses are returned here
	if s.isErrorResp() {
		resp := s.getResponse()
		return nil, resp.err
	}
	resp, err := s.srv.Client().Do(req)
	if err != nil {
		return resp, err
	}
	// wrap the response body in a readFailer if the mock-read-error header is set
	if resp.Header.Get(mockReadError) != "" {
		resp.Body = &readFailer{wrapped: resp.Body}
		resp.Header.Del(mockReadError)
	}
	return resp, err
}

func (s *Server) serveHTTP(w http.ResponseWriter, req *http.Request) {
	var resp mockResponse
	for {
		// grab next response from the queue
		resp = s.getResponse()
		if resp.pred == nil {
			// no predicate, we're done
			break
		} else if resp.pred(req) {
			// response applies to this request, so remove the next response
			s.getResponse()
			break
		}
	}
	if resp.delay > 0 {
		time.Sleep(resp.delay)
	}
	resp.write(w)
}

// Requests returns the number of times an HTTP request was made.
func (s *Server) Requests() int {
	return s.count
}

// AppendError appends the error to the end of the response queue.
func (s *Server) AppendError(err error) {
	s.resp = append(s.resp, mockResponse{err: err})
}

// RepeatError appends the error n number of times to the end of the response queue.
func (s *Server) RepeatError(n int, err error) {
	for i := 0; i < n; i++ {
		s.AppendError(err)
	}
}

// SetError indicates the same error should always be returned.
// Any responses set via other methods will be ignored.
func (s *Server) SetError(err error) {
	s.static = &mockResponse{err: err}
}

// AppendResponse appends the response to the end of the response queue.
// If no options are provided the default response is an http.StatusOK.
func (s *Server) AppendResponse(opts ...ResponseOption) {
	mr := mockResponse{code: http.StatusOK, headers: http.Header{}}
	for _, o := range opts {
		o.apply(&mr)
	}
	s.resp = append(s.resp, mr)
}

// RepeatResponse appends the response n number of times to the end of the response queue.
// If no options are provided the default response is an http.StatusOK.
func (s *Server) RepeatResponse(n int, opts ...ResponseOption) {
	for i := 0; i < n; i++ {
		s.AppendResponse(opts...)
	}
}

// SetResponse indicates the same response should always be returned.
// Any responses set via other methods will be ignored.
// If no options are provided the default response is an http.StatusOK.
// NOTE: does not support WithPredicate(), will cause a panic.
func (s *Server) SetResponse(opts ...ResponseOption) {
	mr := mockResponse{code: http.StatusOK, headers: http.Header{}}
	for _, o := range opts {
		o.apply(&mr)
	}
	if mr.pred != nil {
		panic("WithPredicate not supported for static responses")
	}
	s.static = &mr
}

// ResponseOption is an abstraction for configuring a mock HTTP response.
type ResponseOption interface {
	apply(mr *mockResponse)
}

type fnRespOpt func(*mockResponse)

func (fn fnRespOpt) apply(mr *mockResponse) {
	fn(mr)
}

// ResponsePredicate is a predicate that's invoked in response to an HTTP request.
type ResponsePredicate func(*http.Request) bool

type mockResponse struct {
	code    int
	body    []byte
	headers http.Header
	err     error
	rerr    bool
	delay   time.Duration
	pred    ResponsePredicate
}

func (mr mockResponse) write(w http.ResponseWriter) {
	if len(mr.headers) > 0 {
		for k, v := range mr.headers {
			for _, vv := range v {
				w.Header().Add(k, vv)
			}
		}
	}
	if mr.rerr {
		w.Header().Add(mockReadError, "true")
	}
	w.WriteHeader(mr.code)
	if mr.body != nil {
		w.Write(mr.body)
	}
}

// WithStatusCode sets the HTTP response's status code to the specified value.
func WithStatusCode(c int) ResponseOption {
	return fnRespOpt(func(mr *mockResponse) {
		mr.code = c
	})
}

// WithBody sets the HTTP response's body to the specified value.
func WithBody(b []byte) ResponseOption {
	return fnRespOpt(func(mr *mockResponse) {
		mr.body = b
	})
}

// WithHeader adds the specified header and value to the HTTP response.
func WithHeader(k, v string) ResponseOption {
	return fnRespOpt(func(mr *mockResponse) {
		mr.headers.Add(k, v)
	})
}

// WithSlowResponse will sleep for the specified duration before returning the HTTP response.
func WithSlowResponse(d time.Duration) ResponseOption {
	return fnRespOpt(func(mr *mockResponse) {
		mr.delay = d
	})
}

// WithBodyReadError returns a response that will fail when reading the body.
func WithBodyReadError() ResponseOption {
	return fnRespOpt(func(mr *mockResponse) {
		mr.rerr = true
	})
}

// WithPredicate invokes the specified predicate func on the HTTP request.
// If the predicate returns true, the response associated with the predicate is
// returned and the next response is removed from the queue. When false, the
// associated response is ignored and the next one is returned.
// NOTE: not supported for static responses, will cause a panic.
func WithPredicate(p ResponsePredicate) ResponseOption {
	return fnRespOpt(func(mr *mockResponse) {
		mr.pred = p
	})
}

type readFailer struct {
	wrapped io.ReadCloser
}

func (r *readFailer) Close() error {
	return r.wrapped.Close()
}

func (r *readFailer) Read(p []byte) (int, error) {
	return 0, errors.New("mock read failure")
}
