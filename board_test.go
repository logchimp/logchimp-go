// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logchimp_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/logchimp/logchimp-go"
	"github.com/logchimp/logchimp-go/internal/testutil"
	"github.com/logchimp/logchimp-go/option"
)

func TestBoardListWithOptionalParams(t *testing.T) {
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
	_, err := client.Boards.List(context.TODO(), logchimp.BoardListParams{
		Created: logchimp.String("created"),
		Limit:   logchimp.String("limit"),
		Page:    logchimp.String("page"),
	})
	if err != nil {
		var apierr *logchimp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
