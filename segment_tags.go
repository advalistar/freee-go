package freee

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"reflect"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathSegments = "segments"
	SegmentID1      = uint32(1)
	SegmentID2      = uint32(2)
	SegmentID3      = uint32(3)
)

type SegmentTags struct {
	SegmentTags []SegmentTag `json:"segment_tags"`
}

type SegmentTagResponse struct {
	SegmentTag SegmentTag `json:"segment_tag"`
}

type SegmentTag struct {
	// セグメントタグID
	ID int32 `json:"id"`
	// セグメントタグ名
	Name string `json:"name"`
	// 備考
	Description *string `json:"description"`
	// ショートカット１ (20文字以内)
	Shortcut1 *string `json:"shortcut1"`
	// ショートカット２ (20文字以内)
	Shortcut2 *string `json:"shortcut2"`
}

type SegmentTagParams struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// セグメントタグ名 (30文字以内)
	Name string `json:"name"`
	// 備考 (30文字以内)
	Description *string `json:"description,omitempty"`
	// ショートカット１ (20文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット２ (20文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
}

type GetSegmentTagsOpts struct {
	Offset int32 `url:"offset,omitempty"`
	Limit  int32 `url:"limit,omitempty"`
}

func (c *Client) GetSegmentTags(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, segmentID int32, opts interface{}) (*SegmentTags, error) {
	var result SegmentTags

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathSegments, fmt.Sprint(segmentID), "tags"), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateSegmentTag(ctx context.Context, reuseTokenSource oauth2.TokenSource, segmentID int32, params SegmentTagParams) (*SegmentTag, error) {
	var result SegmentTagResponse
	err := c.call(ctx, path.Join(APIPathSegments, fmt.Sprint(segmentID), "tags"), http.MethodPost, reuseTokenSource, nil, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.SegmentTag, nil
}

func (c *Client) UpdateSegmentTag(ctx context.Context, reuseTokenSource oauth2.TokenSource, segmentID int32, id int32, params SegmentTagParams) (*SegmentTag, error) {
	var result SegmentTagResponse
	err := c.call(ctx, path.Join(APIPathSegments, fmt.Sprint(segmentID), "tags", fmt.Sprint(id)), http.MethodPut, reuseTokenSource, nil, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.SegmentTag, nil
}

func (c *Client) DestroySegmentTag(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, segmentID int32, id int32) error {
	v, err := query.Values(nil)
	if err != nil {
		return err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathSegments, fmt.Sprint(segmentID), "tags", fmt.Sprint(id)), http.MethodDelete, reuseTokenSource, v, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *Client) GetSegmentTagOrderList() []string {
	str := new(SegmentTag)

	var orderList []string
	for i := 0; i < reflect.TypeOf(*str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(*str).Field(i).Tag.Get("json"))
	}

	return orderList
}
