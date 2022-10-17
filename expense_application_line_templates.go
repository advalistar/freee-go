package freee

import (
	"context"
	"net/http"
	"reflect"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathExpenseApplicationLineTemplates = "expense_application_line_templates"
)

type GetExpenseApplicationLineTemplatesOpts struct {
	Offset int32 `url:"offset,omitempty"`
	Limit  int32 `url:"limit,omitempty"`
}

type ExpenseApplicationLineTemplates struct {
	ExpenseApplicationLineTemplates []ExpenseApplicationLineTemplate `json:"expense_application_line_templates"`
}

type ExpenseApplicationLineTemplate struct {
	// 経費科目ID
	ID int32 `json:"id"`
	// 経費科目名
	Name string `json:"name"`
	// 勘定科目ID
	AccountItemID *int32 `json:"account_item_id,omitempty"`
	// 勘定科目名
	AccountItemName string `json:"account_item_name"`
	// 税区分コード
	TaxCode *int32 `json:"tax_code,omitempty"`
	// 税区分名
	TaxName string `json:"tax_name"`
	// 経費科目の説明
	Description *string `json:"description,omitempty"`
	// 内容の補足
	LineDescription *string `json:"line_description,omitempty"`
	// 添付ファイルの必須/任意
	RequiredReceipt *bool `json:"required_receipt,omitempty"`
}

func (c *Client) GetExpenseApplicationLineTemplates(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, opts interface{}) (*ExpenseApplicationLineTemplates, error) {
	var result ExpenseApplicationLineTemplates

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, APIPathExpenseApplicationLineTemplates, http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *Client) GetExpenseApplicationLineTemplateOrderList() []string {
	str := new(ExpenseApplicationLineTemplate)

	var orderList []string
	for i := 0; i < reflect.TypeOf(*str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(*str).Field(i).Tag.Get("json"))
	}

	return orderList
}
