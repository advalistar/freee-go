package freee

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathItems = "items"
)

type Items struct {
	Items []Item `json:"items"`
}

type ItemResponse struct {
	Item Item `json:"item"`
}

type Item struct {
	// 品目ID
	ID int32 `json:"id"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 品目名 (30文字以内)
	Name string `json:"name"`
	// 更新日(yyyy-mm-dd)
	UpdateDate string `json:"update_date"`
	// 品目の使用設定（true: 使用する、false: 使用しない）
	Available bool `json:"available"`
	// ショートカット１ (20文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット２ (20文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
}

type GetItemsOpts struct {
	Offset uint32 `url:"offset,omitempty"`
	Limit  uint32 `url:"limit,omitempty"`
}

type ItemParams struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 品目名 (30文字以内)
	Name string `json:"name"`
	// ショートカット１ (20文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット２ (20文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
}

func (c *Client) GetItems(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID uint32, opts interface{}) (*Items, error) {
	var result Items

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, APIPathItems, http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateItem(ctx context.Context, reuseTokenSource oauth2.TokenSource, params ItemParams) (*Item, error) {
	var result ItemResponse
	err := c.call(ctx, APIPathItems, http.MethodPost, reuseTokenSource, nil, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Item, nil
}

func (c *Client) UpdateItem(ctx context.Context, reuseTokenSource oauth2.TokenSource, params ItemParams, itemID uint32) (*Item, error) {
	var result ItemResponse
	err := c.call(ctx, path.Join(APIPathItems, fmt.Sprint(itemID)), http.MethodPut, reuseTokenSource, nil, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Item, nil
}

func (c *Client) DestroyItem(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID uint32, itemID int32) error {
	v, err := query.Values(nil)
	if err != nil {
		return err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathItems, fmt.Sprint(itemID)), http.MethodDelete, reuseTokenSource, v, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
