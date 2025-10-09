// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logchimp

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/logchimp/logchimp-go/internal/apijson"
	"github.com/logchimp/logchimp-go/internal/apiquery"
	"github.com/logchimp/logchimp-go/internal/requestconfig"
	"github.com/logchimp/logchimp-go/option"
	"github.com/logchimp/logchimp-go/packages/param"
	"github.com/logchimp/logchimp-go/packages/respjson"
)

// BoardService contains methods and other services that help with interacting with
// the logchimp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBoardService] method instead.
type BoardService struct {
	Options []option.RequestOption
}

// NewBoardService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBoardService(opts ...option.RequestOption) (r BoardService) {
	r = BoardService{}
	r.Options = opts
	return
}

// Get all the boards
func (r *BoardService) List(ctx context.Context, query BoardListParams, opts ...option.RequestOption) (res *BoardListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "boards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Board struct {
	BoardID   string `json:"boardId,required" format:"uuid"`
	Color     string `json:"color,required" format:"color"`
	CreatedAt string `json:"createdAt,required"`
	Name      string `json:"name,required"`
	URL       string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoardID     respjson.Field
		Color       respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Board) RawJSON() string { return r.JSON.raw }
func (r *Board) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BoardListResponse struct {
	Boards []BoardListResponseBoard `json:"boards"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Boards      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BoardListResponse) RawJSON() string { return r.JSON.raw }
func (r *BoardListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BoardListResponseBoard struct {
	PostCount string `json:"post_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PostCount   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Board
}

// Returns the unmodified JSON received from the API
func (r BoardListResponseBoard) RawJSON() string { return r.JSON.raw }
func (r *BoardListResponseBoard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BoardListParams struct {
	// Borders are sorted by date creation.
	Created param.Opt[string] `query:"created,omitzero" json:"-"`
	Limit   param.Opt[string] `query:"limit,omitzero" json:"-"`
	Page    param.Opt[string] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BoardListParams]'s query parameters as `url.Values`.
func (r BoardListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
