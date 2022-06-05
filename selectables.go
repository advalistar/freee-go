package freee

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathSelectables = "selectables"
)

type GetSelectablesOpts struct {
	Includes string `url:"includes"`
}

type Selectables struct {
	AccountCategories []AccountCategory `json:"account_categories"`
	AccountGroups     []AccountGroup    `json:"account_groups"`
}

type AccountCategory struct {
	// 収支
	Balance string `json:"balance"`
	// 事業形態（個人事業主: personal、法人: corporate）
	OrgCode string `json:"org_code"`
	// カテゴリーコード
	Role string `json:"role"`
	// カテゴリー名
	Title string `json:"title"`
	// カテゴリーの説明
	Desc *string `json:"desc,omitempty"`
	// 勘定科目の一覧
	AccountItems []SelectableAccountItem `json:"account_items"`
}

type SelectableAccountItem struct {
	// 勘定科目ID
	ID int32 `json:"id"`
	// 勘定科目
	Name *string `json:"name,omitempty"`
	// 勘定科目の説明
	Desc *string `json:"desc,omitempty"`
	// 勘定科目の説明（詳細）
	Help *string `json:"help,omitempty"`
	// ショートカット
	Shortcut   *string     `json:"shortcut,omitempty"`
	DefaultTax *DefaultTax `json:"default_tax,omitempty"`
}

type DefaultTax struct {
	TaxRate5 *TaxRate `json:"tax_rate_5,omitempty"`
	TaxRate8 *TaxRate `json:"tax_rate_8,omitempty"`
}

type TaxRate struct {
	// 税区分ID
	ID int32 `json:"id"`
	// 税区分
	Name string `json:"name"`
}

type AccountGroup struct {
	// 決算書表示名（小カテゴリー）ID
	ID int32 `json:"id"`
	// 決算書表示名
	Name string `json:"name"`
	// 年度ID
	AccountStructureID int32 `json:"account_structure_id"`
	// 勘定科目カテゴリーID
	AccountCategoryID int32 `json:"account_category_id"`
	// 詳細パラメータの種類
	DetailType *int32 `json:"detail_type,omitempty"`
	// 並び順
	Index int32 `json:"index"`
	// 作成日時
	CreatedAt *string `json:"created_at,omitempty"`
	// 更新日時
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (c *Client) GetSelectables(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, opts interface{}) (*Selectables, error) {
	var result Selectables

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join("forms", APIPathSelectables), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
