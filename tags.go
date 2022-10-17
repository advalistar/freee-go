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
	APIPathTags = "tags"
)

type Tags struct {
	Tags []Tag `json:"tags"`
}

type TagResponse struct {
	Tag Tag `json:"tag"`
}

type Tag struct {
	// タグID
	ID int32 `json:"id"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 名前(30文字以内)
	Name string `json:"name"`
	// 更新日(yyyy-mm-dd)
	UpdateDate string `json:"update_date"`
	// ショートカット1 (255文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット2 (255文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
}

type TagParams struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// メモタグ名 (30文字以内)
	Name string `json:"name"`
	// メモタグ検索用 (20文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// メモタグ検索用 (20文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
}

type GetTagsOpts struct {
	StartUpdateDate string `url:"start_update_date,omitempty"`
	EndUpdateDate   string `url:"end_update_date,omitempty"`
	Offset          int32  `url:"offset,omitempty"`
	Limit           int32  `url:"limit,omitempty"`
}

func (c *Client) GetTags(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, opts interface{}) (*Tags, error) {
	var result Tags

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, APIPathTags, http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateTag(ctx context.Context, reuseTokenSource oauth2.TokenSource, params TagParams) (*Tag, error) {
	var result TagResponse
	err := c.call(ctx, APIPathTags, http.MethodPost, reuseTokenSource, nil, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Tag, nil
}

func (c *Client) GetTag(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, tagID int32, opts interface{}) (*Tags, error) {
	var result Tags

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathTags, fmt.Sprint(tagID)), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateTag(ctx context.Context, reuseTokenSource oauth2.TokenSource, tagID int32, params TagParams) (*Tag, error) {
	var result TagResponse
	err := c.call(ctx, path.Join(APIPathTags, fmt.Sprint(tagID)), http.MethodPut, reuseTokenSource, nil, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Tag, nil
}

func (c *Client) DestroyTag(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, tagID int32) error {
	v, err := query.Values(nil)
	if err != nil {
		return err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathTags, fmt.Sprint(tagID)), http.MethodDelete, reuseTokenSource, v, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *Client) GetTagOrderList() []string {
	str := new(Tag)

	var orderList []string
	for i := 0; i < reflect.TypeOf(*str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(*str).Field(i).Tag.Get("json"))
	}

	return orderList
}
