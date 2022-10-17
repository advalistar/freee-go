package freee

import (
	"context"
	"net/http"
	"reflect"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathApprovalFlowRoutes = "approval_flow_routes"
)

type GetApprovalFlowRoutesOpts struct {
	IncludedUserID int32  `url:"included_user_id,omitempty"`
	Usage          string `url:"usage,omitempty"`
	RequestFormID  int32  `url:"request_form_id,omitempty"`
}

type ApprovalFlowRoutes struct {
	ApprovalFlowRoutes []ApprovalFlowRoute `json:"approval_flow_routes"`
}

type ApprovalFlowRoute struct {
	// 申請経路ID
	ID int32 `json:"id"`
	// 申請経路名
	Name *string `json:"name,omitempty"`
	// 申請経路の説明
	Description *string `json:"description,omitempty"`
	// 更新したユーザーのユーザーID
	UserID *int32 `json:"user_id,omitempty"`
	// システム作成の申請経路かどうか
	DefinitionSystem *bool `json:"definition_system,omitempty"`
	// 最初の承認ステップのID
	FirstStepID *int32 `json:"first_step_id,omitempty"`
	// 申請種別（申請経路を使用できる申請種別を示します。例えば、ApprovalRequest の場合は、各種申請で使用できる申請経路です。）
	Usages *[]string `json:"usages,omitempty"`
	// 申請経路で利用できる申請フォームID配列
	RequestFormIDs *[]int32 `json:"request_form_ids,omitempty"`
	// 基本経路として設定されているかどうか
	DefaultRoute bool `json:"default_route"`
}

func (c *Client) GetApprovalFlowRoutes(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, opts interface{}) (*ApprovalFlowRoutes, error) {
	var result ApprovalFlowRoutes

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, APIPathApprovalFlowRoutes, http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *Client) GetApprovalFlowRouteOrderList() []string {
	str := new(ApprovalFlowRoute)

	var orderList []string
	for i := 0; i < reflect.TypeOf(*str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(*str).Field(i).Tag.Get("json"))
	}

	return orderList
}
