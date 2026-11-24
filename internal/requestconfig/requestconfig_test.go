package requestconfig

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

// trackedBody records whether Close was called on a response body.
type trackedBody struct {
	*strings.Reader
	closed bool
}

func (b *trackedBody) Close() error {
	b.closed = true
	return nil
}

// doerFunc adapts a function to the HTTPDoer interface.
type doerFunc func(*http.Request) (*http.Response, error)

func (f doerFunc) Do(req *http.Request) (*http.Response, error) { return f(req) }

func newTestConfig(t *testing.T, ctx context.Context, doer HTTPDoer, maxRetries int) *RequestConfig {
	t.Helper()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "items", nil)
	if err != nil {
		t.Fatal(err)
	}
	base, _ := url.Parse("http://example.test/")
	return &RequestConfig{
		MaxRetries:     maxRetries,
		Request:        req,
		BaseURL:        base,
		CustomHTTPDoer: doer,
	}
}

func response(status int, body *trackedBody) *http.Response {
	h := http.Header{}
	h.Set("Retry-After-Ms", "0") // keep retries instant in tests
	return &http.Response{StatusCode: status, Header: h, Body: body}
}

func TestExecuteClosesBodiesOfRetriedResponses(t *testing.T) {
	var bodies []*trackedBody
	doer := doerFunc(func(*http.Request) (*http.Response, error) {
		b := &trackedBody{Reader: strings.NewReader("{}")}
		bodies = append(bodies, b)
		if len(bodies) < 3 {
			return response(http.StatusServiceUnavailable, b), nil
		}
		return response(http.StatusOK, b), nil
	})

	var res *http.Response
	cfg := newTestConfig(t, context.Background(), doer, 2)
	cfg.ResponseInto = &res

	if err := cfg.Execute(); err != nil {
		t.Fatalf("Execute: %v", err)
	}
	if len(bodies) != 3 {
		t.Fatalf("expected 3 attempts, got %d", len(bodies))
	}
	for i, b := range bodies[:2] {
		if !b.closed {
			t.Errorf("body of retried attempt %d was not closed", i+1)
		}
	}
	if bodies[2].closed {
		t.Error("body of the final response was closed before the caller could read it")
	}
}

func TestExecuteClosesBodyWhenContextCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	body := &trackedBody{Reader: strings.NewReader("{}")}
	doer := doerFunc(func(*http.Request) (*http.Response, error) {
		cancel() // context is canceled after the response arrives
		return response(http.StatusOK, body), nil
	})

	err := newTestConfig(t, ctx, doer, 0).Execute()
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
	if !body.closed {
		t.Error("response body leaked: not closed when Execute returned the context error")
	}
}
