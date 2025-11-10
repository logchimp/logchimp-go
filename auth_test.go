// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logchimp_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/logchimp/logchimp-go"
	"github.com/logchimp/logchimp-go/internal/testutil"
	"github.com/logchimp/logchimp-go/option"
)

func TestAuthLogin(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := logchimp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Auth.Login(context.TODO(), logchimp.AuthLoginParams{
		Email:    "mike@example.com",
		Password: "password",
	})
	if err != nil {
		var apierr *logchimp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// Without the prism server, we can mock the response for testing purposes.
// This is useful for unit tests where we want to simulate the server's response without needing a live server.
func TestAuthLogin_MockServer(t *testing.T) {
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/auth/login" {
			t.Fatalf("wrong endpoint: %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"user": {
				"email": "mock@example.com",
				"authToken": "abc123",
				"name": "Mock User",
				"userId": "1",
				"username": "mock"
			}
		}`))
	}))
	defer mock.Close()

	client := logchimp.NewClient(
		option.WithBaseURL(mock.URL),
	)

	resp, err := client.Auth.Login(context.TODO(), logchimp.AuthLoginParams{
		Email:    "mock@example.com",
		Password: "pass",
	})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if resp.User.Email != "mock@example.com" {
		t.Errorf("got %s, want mock@example.com", resp.User.Email)
	}
}

func TestAuthSignup(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := logchimp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Auth.Signup(context.TODO(), logchimp.AuthSignupParams{
		Email:    "mike@example.com",
		Password: "password",
	})
	if err != nil {
		var apierr *logchimp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
