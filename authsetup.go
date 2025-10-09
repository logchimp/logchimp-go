// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logchimp

import (
	"context"
	"net/http"
	"slices"

	"github.com/logchimp/logchimp-go/internal/apijson"
	"github.com/logchimp/logchimp-go/internal/requestconfig"
	"github.com/logchimp/logchimp-go/option"
	"github.com/logchimp/logchimp-go/packages/param"
	"github.com/logchimp/logchimp-go/packages/respjson"
)

// AuthSetupService contains methods and other services that help with interacting
// with the logchimp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthSetupService] method instead.
type AuthSetupService struct {
	Options []option.RequestOption
}

// NewAuthSetupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthSetupService(opts ...option.RequestOption) (r AuthSetupService) {
	r = AuthSetupService{}
	r.Options = opts
	return
}

// Setup the site by creating the owner account.
func (r *AuthSetupService) New(ctx context.Context, body AuthSetupNewParams, opts ...option.RequestOption) (res *AuthSetupNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/setup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get site setup information.
func (r *AuthSetupService) Get(ctx context.Context, opts ...option.RequestOption) (res *AuthSetupGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/setup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AuthSetupNewResponse struct {
	User AuthUser `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthSetupNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthSetupNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthSetupGetResponse struct {
	// Is LogChimp onboarding completed?
	IsSetup bool `json:"is_setup"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSetup     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthSetupGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthSetupGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthSetupNewParams struct {
	Email    string            `json:"email,required" format:"email"`
	Password string            `json:"password,required" format:"password"`
	Name     param.Opt[string] `json:"name,omitzero"`
	// This will be shown across the website on public pages and dashboard.
	SiteTitle param.Opt[string] `json:"siteTitle,omitzero"`
	paramObj
}

func (r AuthSetupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthSetupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthSetupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
