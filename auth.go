// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logchimp

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/logchimp-go/internal/apijson"
	"github.com/stainless-sdks/logchimp-go/internal/requestconfig"
	"github.com/stainless-sdks/logchimp-go/option"
	"github.com/stainless-sdks/logchimp-go/packages/param"
	"github.com/stainless-sdks/logchimp-go/packages/respjson"
)

// AuthService contains methods and other services that help with interacting with
// the logchimp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	Options []option.RequestOption
	Setup   AuthSetupService
	Email   AuthEmailService
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r AuthService) {
	r = AuthService{}
	r.Options = opts
	r.Setup = NewAuthSetupService(opts...)
	r.Email = NewAuthEmailService(opts...)
	return
}

// Login to account using email and password.
func (r *AuthService) Login(ctx context.Context, body AuthLoginParams, opts ...option.RequestOption) (res *AuthLoginResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new user account.
func (r *AuthService) Signup(ctx context.Context, body AuthSignupParams, opts ...option.RequestOption) (res *AuthSignupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/signup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AuthUser struct {
	AuthToken string `json:"authToken"`
	// URL of the user's avatar
	Avatar string `json:"avatar" format:"url"`
	// Email address of the user
	Email string `json:"email" format:"email"`
	// User's full name
	Name   string `json:"name"`
	UserID string `json:"userId"`
	// User's username
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthToken   respjson.Field
		Avatar      respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		UserID      respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthUser) RawJSON() string { return r.JSON.raw }
func (r *AuthUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthLoginResponse struct {
	User AuthUser `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthLoginResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthLoginResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthSignupResponse struct {
	User AuthUser `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthSignupResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthSignupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthLoginParams struct {
	Email    string `json:"email,required" format:"email"`
	Password string `json:"password,required" format:"password"`
	paramObj
}

func (r AuthLoginParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthLoginParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthLoginParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthSignupParams struct {
	Email    string `json:"email,required" format:"email"`
	Password string `json:"password,required" format:"password"`
	paramObj
}

func (r AuthSignupParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthSignupParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthSignupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
