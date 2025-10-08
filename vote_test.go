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

func TestVoteAdd(t *testing.T) {
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
	_, err := client.Votes.Add(context.TODO(), logchimp.VoteAddParams{
		PostID: "postId",
	})
	if err != nil {
		var apierr *logchimp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVoteRemove(t *testing.T) {
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
	_, err := client.Votes.Remove(context.TODO(), logchimp.VoteRemoveParams{
		PostID: "postId",
	})
	if err != nil {
		var apierr *logchimp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
