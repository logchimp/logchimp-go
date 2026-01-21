// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package option

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"time"
)

// WithDebugLog logs the HTTP request and response content.
// If the logger parameter is nil, it uses the default logger.
//
// WithDebugLog is for debugging and development purposes only.
// It should not be used in production code. The behavior and interface
// of WithDebugLog is not guaranteed to be stable.
func WithDebugLog(logger *log.Logger) RequestOption {
	return WithMiddleware(func(req *http.Request, nxt MiddlewareNext) (*http.Response, error) {
		if logger == nil {
			logger = log.Default()
		}

		if reqBytes, err := httputil.DumpRequest(req, true); err == nil {
			logger.Printf("Request Content:\n%s\n", reqBytes)
		}

		resp, err := nxt(req)
		if err != nil {
			return resp, err
		}

		if respBytes, err := httputil.DumpResponse(resp, true); err == nil {
			logger.Printf("Response Content:\n%s\n", respBytes)
		}

		return resp, err
	})
}

// WithLogger logs the HTTP request and response information using a structured
// logger.
func WithLogger(logger StructuredLogger) RequestOption {
	return WithMiddleware(func(req *http.Request, nxt MiddlewareNext) (*http.Response, error) {
		if logger == nil {
			logger = slog.Default()
		}

		start := time.Now()
		resp, err := nxt(req)
		duration := time.Since(start)

		if err != nil {
			logger.Error("HTTP request failed",
				slog.String("method", req.Method),
				slog.String("url", req.URL.String()),
				slog.Duration("duration", duration),
				slog.Any("error", err),
			)
			return resp, err
		}

		logger.Info("HTTP request completed",
			slog.String("method", req.Method),
			slog.String("url", req.URL.String()),
			slog.Int("status", resp.StatusCode),
			slog.Duration("duration", duration),
		)

		return resp, err
	})
}

// StructuredLogger is an interface for structured logging, compatible with
// [log/slog.Logger].
type StructuredLogger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
	Warn(msg string, args ...any)
	Debug(msg string, args ...any)
	Log(ctx context.Context, level slog.Level, msg string, args ...any)
}
