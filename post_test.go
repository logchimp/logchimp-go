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

func TestPostNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Posts.New(context.TODO(), logchimp.PostNewParams{
		Title:           "title",
		BoardID:         logchimp.String("boardId"),
		ContentMarkdown: logchimp.String("contentMarkdown"),
		RoadmapID:       logchimp.String("roadmapId"),
	})
	if err != nil {
		var apierr *logchimp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPostUpdate(t *testing.T) {
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
	err := client.Posts.Update(context.TODO())
	if err != nil {
		var apierr *logchimp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPostListWithOptionalParams(t *testing.T) {
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
	_, err := client.Posts.List(context.TODO(), logchimp.PostListParams{
		Created:   "desc",
		Page:      1,
		BoardID:   []string{"11c14790-b4d3-4689-942c-55548cfdd8a6"},
		Limit:     logchimp.Float(10),
		RoadmapID: logchimp.String("11c14790-b4d3-4689-942c-55548cfdd8a6"),
		UserID:    logchimp.String("11c14790-b4d3-4689-942c-55548cfdd8a6"),
	})
	if err != nil {
		var apierr *logchimp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPostDelete(t *testing.T) {
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
	err := client.Posts.Delete(context.TODO(), logchimp.PostDeleteParams{
		ID: "id",
	})
	if err != nil {
		var apierr *logchimp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPostGetBySlugWithOptionalParams(t *testing.T) {
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
	_, err := client.Posts.GetBySlug(context.TODO(), logchimp.PostGetBySlugParams{
		Slug:   "slug",
		UserID: logchimp.String("11c14790-b4d3-4689-942c-55548cfdd8a6"),
	})
	if err != nil {
		var apierr *logchimp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
