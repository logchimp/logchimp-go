// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logchimp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/logchimp-go/internal/apijson"
	"github.com/stainless-sdks/logchimp-go/internal/apiquery"
	"github.com/stainless-sdks/logchimp-go/internal/requestconfig"
	"github.com/stainless-sdks/logchimp-go/option"
	"github.com/stainless-sdks/logchimp-go/packages/param"
	"github.com/stainless-sdks/logchimp-go/packages/respjson"
)

// RoadmapService contains methods and other services that help with interacting
// with the logchimp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoadmapService] method instead.
type RoadmapService struct {
	Options []option.RequestOption
}

// NewRoadmapService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRoadmapService(opts ...option.RequestOption) (r RoadmapService) {
	r = RoadmapService{}
	r.Options = opts
	return
}

// Get all roadmaps
func (r *RoadmapService) List(ctx context.Context, query RoadmapListParams, opts ...option.RequestOption) (res *RoadmapListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "roadmaps"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get roadmap by Url
func (r *RoadmapService) GetByURL(ctx context.Context, url string, opts ...option.RequestOption) (res *RoadmapGetByURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if url == "" {
		err = errors.New("missing required url parameter")
		return
	}
	path := fmt.Sprintf("roadmaps/%s", url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Search roadmaps by name
func (r *RoadmapService) SearchByName(ctx context.Context, name string, opts ...option.RequestOption) (res *RoadmapSearchByNameResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("roadmaps/search/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Roadmap struct {
	ID    any `json:"id,required"`
	Color any `json:"color,required"`
	Name  any `json:"name,required"`
	URL   any `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Color       respjson.Field
		Name        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Roadmap) RawJSON() string { return r.JSON.raw }
func (r *Roadmap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoadmapPrivate struct {
	// Roadmap creation date
	CreatedAt string  `json:"created_at,required"`
	Display   bool    `json:"display,required"`
	Index     float64 `json:"index,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Display     respjson.Field
		Index       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Roadmap
}

// Returns the unmodified JSON received from the API
func (r RoadmapPrivate) RawJSON() string { return r.JSON.raw }
func (r *RoadmapPrivate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoadmapListResponse struct {
	PageInfo   RoadmapListResponsePageInfo `json:"page_info,required"`
	Results    []RoadmapPrivate            `json:"results,required"`
	Roadmaps   []RoadmapPrivate            `json:"roadmaps,required"`
	TotalCount int64                       `json:"total_count,nullable"`
	TotalPages int64                       `json:"total_pages,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageInfo    respjson.Field
		Results     respjson.Field
		Roadmaps    respjson.Field
		TotalCount  respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoadmapListResponse) RawJSON() string { return r.JSON.raw }
func (r *RoadmapListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoadmapListResponsePageInfo struct {
	// Number of items in current page
	Count int64 `json:"count,required"`
	// Current page number
	CurrentPage int64 `json:"current_page,required"`
	// Whether there are more pages
	HasNextPage bool `json:"has_next_page,required"`
	// Cursor for the last item
	EndCursor string `json:"end_cursor,nullable"`
	// Cursor for the first item
	StartCursor string `json:"start_cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		CurrentPage respjson.Field
		HasNextPage respjson.Field
		EndCursor   respjson.Field
		StartCursor respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoadmapListResponsePageInfo) RawJSON() string { return r.JSON.raw }
func (r *RoadmapListResponsePageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoadmapGetByURLResponse struct {
	Roadmap Roadmap `json:"roadmap"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Roadmap     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoadmapGetByURLResponse) RawJSON() string { return r.JSON.raw }
func (r *RoadmapGetByURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoadmapSearchByNameResponse struct {
	Roadmaps []RoadmapPrivate `json:"roadmaps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Roadmaps    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoadmapSearchByNameResponse) RawJSON() string { return r.JSON.raw }
func (r *RoadmapSearchByNameResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoadmapListParams struct {
	// Cursor for pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Number of items to return in a request
	First param.Opt[float64] `query:"first,omitzero" json:"-"`
	// Any of "public", "private".
	Visibility []string `query:"visibility,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoadmapListParams]'s query parameters as `url.Values`.
func (r RoadmapListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
