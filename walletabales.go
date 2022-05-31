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
	APIPathWalletables = "walletables"
)

type WalletablesResponse struct {
	Walletables []Walletable `json:"walletables"`
	Meta        Meta         `json:"meta"`
}

type WalletableResponse struct {
	Walletable Walletable `json:"walletable"`
	Meta       Meta       `json:"meta"`
}

type Meta struct {
	UpToDate bool `json:"up_to_date"`
}

type GetWalletablesOpts struct {
	// 残高情報を含める
	WithBalance bool `url:"with_balance,omitempty"`
	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, その他の決済口座: wallet)
	Type string `url:"type,omitempty"`
}

type Walletable struct {
	// 口座ID
	ID uint64 `json:"id"`
	// 口座名 (255文字以内)
	Name string `json:"name"`
	// サービスID
	BankID uint64 `json:"bank_id"`
	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet)
	Type string `json:"type"`
	// 同期残高
	LastBalance int64 `json:"last_balance,omitempty"`
	// 登録残高
	WalletableBalance int64 `json:"walletable_balance,omitempty"`
}

func (c *Client) GetWalletables(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID uint32, opts GetWalletablesOpts) (*WalletablesResponse, error) {
	var result WalletablesResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathWalletables), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetWalletable(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID uint32, walletableID uint64, opts GetWalletTxnOpts) (*Walletable, error) {
	var result WalletableResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathWalletables, fmt.Sprint(walletableID)), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result.Walletable, nil
}
