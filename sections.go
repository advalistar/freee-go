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
	APIPathSections = "sections"
)

type Sections struct {
	Sections []Section `json:"sections"`
}

type SectionResponse struct {
	Section Section `json:"section"`
}

type Section struct {
	// 品目ID
	ID int32 `json:"id"`
	// 品目名 (30文字以内)
	Name string `json:"name"`
	// 部門の使用設定（true: 使用する、false: 使用しない）
	Available bool `json:"available"`
	// 正式名称（255文字以内）
	LongName *string `json:"long_name,omitempty"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// ショートカット１ (20文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット２ (20文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
	// 部門階層
	IndentCount *int32 `json:"indent_count,omitempty"`
	// 親部門ID
	ParentID *int32 `json:"parent_id,omitempty"`
}

type SectionParams struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 部門名 (30文字以内)
	Name string `json:"name"`
	// 正式名称 (255文字以内)
	LongName *string `json:"long_name,omitempty"`
	// ショートカット１ (20文字以内)
	Shortcut1 *string `json:"shortcut1,omitempty"`
	// ショートカット２ (20文字以内)
	Shortcut2 *string `json:"shortcut2,omitempty"`
	// 親部門ID (ビジネスプラン以上)
	ParentID *int32 `json:"parent_id,omitempty"`
}

func (c *Client) GetSections(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32) (*Sections, error) {
	var result Sections

	v, err := query.Values(nil)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, APIPathSections, http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateSection(ctx context.Context, reuseTokenSource oauth2.TokenSource, params SectionParams) (*Section, error) {
	var result SectionResponse
	err := c.call(ctx, APIPathSections, http.MethodPost, reuseTokenSource, nil, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Section, nil
}

func (c *Client) UpdateSection(ctx context.Context, reuseTokenSource oauth2.TokenSource, sectionID int32, params SectionParams) (*Section, error) {
	var result SectionResponse
	err := c.call(ctx, path.Join(APIPathSections, fmt.Sprint(sectionID)), http.MethodPut, reuseTokenSource, nil, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Section, nil
}

func (c *Client) DestroySection(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, sectionID int32) error {
	v, err := query.Values(nil)
	if err != nil {
		return err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathSections, fmt.Sprint(sectionID)), http.MethodDelete, reuseTokenSource, v, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *Client) GetSectionOrderList() []string {
	str := new(Section)

	var orderList []string
	for i := 0; i < reflect.TypeOf(*str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(*str).Field(i).Tag.Get("json"))
	}

	return orderList
}
