// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logchimp_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/logchimp-go"
	"github.com/stainless-sdks/logchimp-go/internal/testutil"
	"github.com/stainless-sdks/logchimp-go/option"
)

func TestAuthSetupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Auth.Setup.New(context.TODO(), logchimp.AuthSetupNewParams{
		Email:     "mike@example.com",
		Password:  "password",
		Name:      logchimp.String("Mike"),
		SiteTitle: logchimp.String("LogChimp"),
	})
	if err != nil {
		var apierr *logchimp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthSetupGet(t *testing.T) {
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
	_, err := client.Auth.Setup.Get(context.TODO())
	if err != nil {
		var apierr *logchimp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
