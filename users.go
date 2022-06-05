package freee

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathUsers = "users"
)

type GetUsersOpts struct {
	Limit *int32 `url:"limit,omitempty"`
}

type Users struct {
	User []User `json:"users"`
}

type Me struct {
	User User `json:"user"`
}

type User struct {
	// ユーザーID
	ID int32 `json:"id"`
	// メールアドレス
	Email string `json:"email"`
	// 表示ユーザー名
	DisplayName *string `json:"display_name,omitempty"`
	// 名
	FirstName *string `json:"first_name,omitempty"`
	// 姓
	LastName *string `json:"last_name,omitempty"`
	// 名（カナ）
	FirstNameKana *string `json:"first_name_kana,omitempty"`
	// 姓（カナ）
	LastNameKana *string        `json:"last_name_kana,omitempty"`
	Companies    *[]UserCompany `json:"companies,omitempty"`
}

type UserCompany struct {
	// 事業所ID
	ID int32 `json:"id"`
	// 表示名
	DisplayName string `json:"display_name"`
	// ユーザーの権限
	Role string `json:"role"`
	// カスタム権限（true: 使用する、false: 使用しない）
	UseCustomRole bool `json:"use_custom_role"`
}

type GetUsersMeOpts struct {
	Companies bool `url:"companies,omitempty"`
}

func (c *Client) GetUsers(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, opts interface{}) (*Users, error) {
	var result Users

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, APIPathUsers, http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetUsersMe(ctx context.Context, reuseTokenSource oauth2.TokenSource, opts interface{}) (*Me, error) {
	var result Me

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	err = c.call(ctx, path.Join(APIPathUsers, "me"), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
