// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logchimp

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/logchimp-go/internal/apijson"
	"github.com/stainless-sdks/logchimp-go/internal/requestconfig"
	"github.com/stainless-sdks/logchimp-go/option"
	"github.com/stainless-sdks/logchimp-go/packages/respjson"
)

// AuthEmailService contains methods and other services that help with interacting
// with the logchimp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthEmailService] method instead.
type AuthEmailService struct {
	Options []option.RequestOption
}

// NewAuthEmailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthEmailService(opts ...option.RequestOption) (r AuthEmailService) {
	r = AuthEmailService{}
	r.Options = opts
	return
}

// Send a verification email to the user email address. The email is only sent if
// have configured
// [SMTP mail server](https://docs.logchimp.codecarrot.net/docs/environment-variables#mail)
// are configured at the time of deploying LogChimp.
func (r *AuthEmailService) SendVerification(ctx context.Context, opts ...option.RequestOption) (res *AuthEmailSendVerificationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/email/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AuthEmailSendVerificationResponse struct {
	Verify AuthEmailSendVerificationResponseVerify `json:"verify"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Verify      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthEmailSendVerificationResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthEmailSendVerificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthEmailSendVerificationResponseVerify struct {
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthEmailSendVerificationResponseVerify) RawJSON() string { return r.JSON.raw }
func (r *AuthEmailSendVerificationResponseVerify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
