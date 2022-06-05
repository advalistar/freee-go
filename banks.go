package freee

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathBanks = "banks"
)

type GetBanksOpts struct {
	Offset int32  `url:"offset,omitempty"`
	Limit  int32  `url:"limit,omitempty"`
	Type   string `url:"type,omitempty"`
}

type Banks struct {
	Banks []bank `json:"banks"`
}

type bank struct {
	// 連携サービスID
	ID int32 `json:"id"`
	// 連携サービス名
	Name *string `json:"name,omitempty"`
	// 連携サービス種別: (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet)
	Type *string `json:"type,omitempty"`
	// 連携サービス名(カナ)
	NameKana *string `json:"name_kana,omitempty"`
}

func (c *Client) GetBanks(ctx context.Context, reuseTokenSource oauth2.TokenSource, opts interface{}) (*Banks, error) {
	var result Banks

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	err = c.call(ctx, APIPathBanks, http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
