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

// VoteService contains methods and other services that help with interacting with
// the logchimp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoteService] method instead.
type VoteService struct {
	Options []option.RequestOption
}

// NewVoteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVoteService(opts ...option.RequestOption) (r VoteService) {
	r = VoteService{}
	r.Options = opts
	return
}

// Vote on a post
func (r *VoteService) Add(ctx context.Context, body VoteAddParams, opts ...option.RequestOption) (res *VoteAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "votes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove vote from a post
func (r *VoteService) Remove(ctx context.Context, body VoteRemoveParams, opts ...option.RequestOption) (res *VoteRemoveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "votes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type CurrentUserVote struct {
	CreatedAt string `json:"createdAt,required"`
	PostID    string `json:"postId,required"`
	UserID    string `json:"userId,required"`
	VoteID    string `json:"voteId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		PostID      respjson.Field
		UserID      respjson.Field
		VoteID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CurrentUserVote) RawJSON() string { return r.JSON.raw }
func (r *CurrentUserVote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostVote struct {
	Votes      []PostVoteVote  `json:"votes,required"`
	VotesCount float64         `json:"votesCount,required"`
	ViewerVote CurrentUserVote `json:"viewerVote"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Votes       respjson.Field
		VotesCount  respjson.Field
		ViewerVote  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostVote) RawJSON() string { return r.JSON.raw }
func (r *PostVote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostVoteVote struct {
	Username string `json:"username,required"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Username    respjson.Field
		Avatar      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	CurrentUserVote
}

// Returns the unmodified JSON received from the API
func (r PostVoteVote) RawJSON() string { return r.JSON.raw }
func (r *PostVoteVote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoteAddResponse struct {
	Voters PostVote `json:"voters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Voters      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoteAddResponse) RawJSON() string { return r.JSON.raw }
func (r *VoteAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoteRemoveResponse struct {
	Voters PostVote `json:"voters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Voters      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoteRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *VoteRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoteAddParams struct {
	// ID of the post to vote on.
	PostID string `json:"postId,required"`
	paramObj
}

func (r VoteAddParams) MarshalJSON() (data []byte, err error) {
	type shadow VoteAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoteAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoteRemoveParams struct {
	// ID of the post to vote on.
	PostID string `json:"postId,required"`
	paramObj
}

func (r VoteRemoveParams) MarshalJSON() (data []byte, err error) {
	type shadow VoteRemoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoteRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
