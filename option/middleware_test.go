package option_test

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"testing"

	"github.com/logchimp/logchimp-go"
	"github.com/logchimp/logchimp-go/option"
)

type mockLogger struct {
	infoCalled  bool
	errorCalled bool
	lastMsg     string
}

func (m *mockLogger) Info(msg string, args ...any) {
	m.infoCalled = true
	m.lastMsg = msg
}

func (m *mockLogger) Error(msg string, args ...any) {
	m.errorCalled = true
	m.lastMsg = msg
}

func (m *mockLogger) Warn(msg string, args ...any)                                       {}
func (m *mockLogger) Debug(msg string, args ...any)                                      {}
func (m *mockLogger) Log(ctx context.Context, level slog.Level, msg string, args ...any) {}

type closureTransport struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (t *closureTransport) Do(req *http.Request) (*http.Response, error) {
	return t.fn(req)
}

func TestWithLogger(t *testing.T) {
	logger := &mockLogger{}
	client := logchimp.NewClient(
		option.WithHTTPClient(&closureTransport{
			fn: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString("{}")),
				}, nil
			},
		}),
		option.WithLogger(logger),
	)

	client.Auth.Login(context.Background(), logchimp.AuthLoginParams{
		Email:    "test@example.com",
		Password: "password",
	})

	if !logger.infoCalled {
		t.Error("Expected logger.Info to be called")
	}
	if logger.lastMsg != "HTTP request completed" {
		t.Errorf("Expected message 'HTTP request completed', got %q", logger.lastMsg)
	}
}
