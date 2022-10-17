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
	APIPathReceipts = "receipts"
)

type CreateReceiptParams struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// メモ (255文字以内)
	Description string `json:"description,omitempty"`
	// 取引日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 証憑ファイル
	Receipt []byte `json:"receipt"`
}

type Receipts struct {
	Receipts []Receipt `json:"receipts"`
}

type ReceiptResponse struct {
	Receipt Receipt `json:"receipt"`
}

type Recipts struct {
	Recipts []Receipt `json:"recipts"`
}

type Receipt struct {
	// 証憑ID
	ID int32 `json:"id"`
	// ステータス(unconfirmed:確認待ち、confirmed:確認済み、deleted:削除済み、ignored:無視)
	Status string `json:"status"`
	// メモ
	Description *string `json:"description,omitempty"`
	// MIMEタイプ
	MimeType string `json:"mime_type"`
	// 発生日
	IssueDate *string `json:"issue_date,omitempty"`
	// アップロード元種別
	Origin string `json:"origin"`
	// 作成日時（ISO8601形式）
	CreatedAt string             `json:"created_at"`
	User      UserCreatedReceipt `json:"user"`
}

type GetReceiptOpts struct {
	StartDate        string `url:"start_date"`
	EndDate          string `url:"end_date"`
	UserName         string `url:"user_name,omitempty"`
	Number           int32  `url:"number,omitempty"`
	CommentType      string `url:"comment_type,omitempty"`
	CommentImportant bool   `url:"comment_important,omitempty"`
	Category         string `url:"category,omitempty"`
	Offset           int32  `url:"offset,omitempty"`
	Limit            int32  `url:"limit,omitempty"`
}

type UserCreatedReceipt struct {
	// ユーザーID
	ID int32 `json:"id"`
	// メールアドレス
	Email string `json:"email"`
	// 表示名
	DisplayName *string `json:"display_name,omitempty"`
}

func (c *Client) CreateReceipt(ctx context.Context, reuseTokenSource oauth2.TokenSource, params CreateReceiptParams, receiptName string) (*ReceiptResponse, error) {
	postBody := map[string]string{
		"company_id":  fmt.Sprint(params.CompanyID),
		"description": params.Description,
		"issue_date":  params.IssueDate,
	}
	var result ReceiptResponse
	err := c.postFiles(ctx, APIPathReceipts, http.MethodPost, reuseTokenSource, nil, postBody, receiptName, params.Receipt, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetReceipt(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, receiptID int32) (*ReceiptResponse, error) {
	var result ReceiptResponse

	v, err := query.Values(nil)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathReceipts, fmt.Sprint(receiptID)), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetReceipts(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, opts interface{}) (*Recipts, error) {
	var result Recipts

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, APIPathReceipts, http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *Client) GetReceiptOrderList() []string {
	str := new(Receipt)

	var orderList []string
	for i := 0; i < reflect.TypeOf(*str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(*str).Field(i).Tag.Get("json"))
	}

	return orderList
}
