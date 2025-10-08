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

// PostService contains methods and other services that help with interacting with
// the logchimp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPostService] method instead.
type PostService struct {
	Options []option.RequestOption
}

// NewPostService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPostService(opts ...option.RequestOption) (r PostService) {
	r = PostService{}
	r.Options = opts
	return
}

// Create a new post
func (r *PostService) New(ctx context.Context, body PostNewParams, opts ...option.RequestOption) (res *PostNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "posts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a post
func (r *PostService) Update(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "posts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, nil, opts...)
	return
}

// Get all the posts
func (r *PostService) List(ctx context.Context, body PostListParams, opts ...option.RequestOption) (res *PostListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "posts/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a posts
func (r *PostService) Delete(ctx context.Context, body PostDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "posts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Get post by Slug
func (r *PostService) GetBySlug(ctx context.Context, body PostGetBySlugParams, opts ...option.RequestOption) (res *PostGetBySlugResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "posts/slug"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Post struct {
	Author    PostAuthor `json:"author,required"`
	Board     Board      `json:"board,required"`
	CreatedAt string     `json:"createdAt,required"`
	PostID    string     `json:"postId,required"`
	Slug      string     `json:"slug,required"`
	// Title of the post.
	Title     string   `json:"title,required"`
	UpdatedAt string   `json:"updatedAt,required"`
	Voters    PostVote `json:"voters,required"`
	// Post description.
	ContentMarkdown string  `json:"contentMarkdown"`
	Roadmap         Roadmap `json:"roadmap"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Author          respjson.Field
		Board           respjson.Field
		CreatedAt       respjson.Field
		PostID          respjson.Field
		Slug            respjson.Field
		Title           respjson.Field
		UpdatedAt       respjson.Field
		Voters          respjson.Field
		ContentMarkdown respjson.Field
		Roadmap         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Post) RawJSON() string { return r.JSON.raw }
func (r *Post) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostAuthor struct {
	UserID   string `json:"userId,required"`
	Username string `json:"username,required"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserID      respjson.Field
		Username    respjson.Field
		Avatar      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostAuthor) RawJSON() string { return r.JSON.raw }
func (r *PostAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostNewResponse struct {
	Post PostNewResponsePost `json:"post"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Post        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PostNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostNewResponsePost struct {
	PostID  string `json:"postId,required"`
	BoardID string `json:"boardId"`
	// Post description.
	ContentMarkdown string `json:"contentMarkdown"`
	CreatedAt       string `json:"createdAt"`
	RoadmapID       string `json:"roadmap_id"`
	Slug            string `json:"slug"`
	SlugID          string `json:"slugId"`
	// Title of the post.
	Title     string `json:"title"`
	UpdatedAt string `json:"updatedAt"`
	UserID    string `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PostID          respjson.Field
		BoardID         respjson.Field
		ContentMarkdown respjson.Field
		CreatedAt       respjson.Field
		RoadmapID       respjson.Field
		Slug            respjson.Field
		SlugID          respjson.Field
		Title           respjson.Field
		UpdatedAt       respjson.Field
		UserID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostNewResponsePost) RawJSON() string { return r.JSON.raw }
func (r *PostNewResponsePost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostListResponse struct {
	Posts []Post `json:"posts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Posts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostListResponse) RawJSON() string { return r.JSON.raw }
func (r *PostListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostGetBySlugResponse struct {
	Post PostGetBySlugResponsePost `json:"post"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Post        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostGetBySlugResponse) RawJSON() string { return r.JSON.raw }
func (r *PostGetBySlugResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostGetBySlugResponsePost struct {
	SlugID string `json:"slugId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SlugID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Post
}

// Returns the unmodified JSON received from the API
func (r PostGetBySlugResponsePost) RawJSON() string { return r.JSON.raw }
func (r *PostGetBySlugResponsePost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostNewParams struct {
	// Title of the post.
	Title   string            `json:"title,required"`
	BoardID param.Opt[string] `json:"boardId,omitzero"`
	// Post description.
	ContentMarkdown param.Opt[string] `json:"contentMarkdown,omitzero"`
	RoadmapID       param.Opt[string] `json:"roadmapId,omitzero"`
	paramObj
}

func (r PostNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PostNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostListParams struct {
	// Sort posts by creation date in 'asc' or 'desc' order.
	Created string `json:"created,required"`
	// Page number of the results to fetch.
	Page float64 `json:"page,required"`
	// "Number of posts to return in a single request. Default is 10."
	Limit     param.Opt[float64] `json:"limit,omitzero"`
	RoadmapID param.Opt[string]  `json:"roadmapId,omitzero" format:"uuid"`
	UserID    param.Opt[string]  `json:"userId,omitzero" format:"uuid"`
	// Posts will be searched from the list of provided board IDs. If a value is not
	// provided, it defaults to empty array and posts are searched from all the boards.
	BoardID []string `json:"boardId,omitzero" format:"uuid"`
	paramObj
}

func (r PostListParams) MarshalJSON() (data []byte, err error) {
	type shadow PostListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostDeleteParams struct {
	ID string `json:"id,required"`
	paramObj
}

func (r PostDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow PostDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostGetBySlugParams struct {
	Slug   string            `json:"slug,required"`
	UserID param.Opt[string] `json:"userId,omitzero" format:"uuid"`
	paramObj
}

func (r PostGetBySlugParams) MarshalJSON() (data []byte, err error) {
	type shadow PostGetBySlugParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostGetBySlugParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
