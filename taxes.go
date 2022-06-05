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
	APIPathTaxes = "taxes"

	// tax_5: 5%表示の税区分
	TaxRate5 = "tax_5"
	// tax_8: 8%表示の税区分
	TaxRate8 = "tax_8"
	// tax_r8: 軽減税率8%表示の税区分
	TaxRateR8 = "tax_r8"
	// tax_10: 10%表示の税区分
	TaxRate10 = "tax_10"
	// null: 税率未設定税区分
)

type TaxCodes struct {
	TaxCodes []TaxCode `json:"taxes"`
}

type TaxCode struct {
	// 税区分コード
	Code int32 `json:"code"`
	// 税区分名
	Name string `json:"name"`
	// 税区分名（日本語表示用）
	NameJa string `json:"name_ja"`
}

type TaxCompanies struct {
	TaxCompanies []TaxCompany `json:"taxes"`
}

type TaxCompany struct {
	// 税区分コード
	Code int32 `json:"code"`
	// 税区分名
	Name string `json:"name"`
	// 税区分名（日本語表示用）
	NameJa string `json:"name_ja"`
	// 税区分の表示カテゴリ（tax_5: 5%表示の税区分、tax_8: 8%表示の税区分、tax_r8: 軽減税率8%表示の税区分、tax_10: 10%表示の税区分、null: 税率未設定税区分）
	DisplayCategory string `json:"display_category"`
	// true: 使用する、false: 使用しない
	Available bool `json:"available"`
}

func (c *Client) GetTaxCodes(ctx context.Context, reuseTokenSource oauth2.TokenSource) (*TaxCodes, error) {
	var result TaxCodes

	v, err := query.Values(nil)
	if err != nil {
		return nil, err
	}
	err = c.call(ctx, path.Join(APIPathTaxes, "codes"), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTaxCompanies(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32) (*TaxCompanies, error) {
	var result TaxCompanies

	v, err := query.Values(nil)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathTaxes, "companies", fmt.Sprint(companyID)), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
