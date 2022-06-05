package freee

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathPaymentRequests = "payment_requests"
)

type GetPaymentRequestsOpts struct {
	Status               string `url:"status,omitempty"`
	StartApplicationDate string `url:"start_application_date,omitempty"`
	EndApplicationDate   string `url:"end_application_date,omitempty"`
	StartIssueDate       string `url:"start_issue_date,omitempty"`
	EndIssueDate         string `url:"end_issue_date,omitempty"`
	ApplicationNumber    int32  `url:"application_number,omitempty"`
	Title                string `url:"title,omitempty"`
	ApplicantID          int32  `url:"applicant_id,omitempty"`
	ApproverID           int32  `url:"approver_id,omitempty"`
	MinAmount            int32  `url:"min_amount,omitempty"`
	MaxAmount            int32  `url:"max_amount,omitempty"`
	PartnerID            int32  `url:"partner_id,omitempty"`
	PartnerCode          string `url:"partner_code,omitempty"`
	PaymentMethod        string `url:"payment_method,omitempty"`
	StartPaymentDate     string `url:"start_payment_date,omitempty"`
	EndPaymentDate       string `url:"end_payment_date,omitempty"`
	DocumentCode         string `url:"document_code,omitempty"`
	Offset               int32  `url:"offset,omitempty"`
	Limit                int32  `url:"limit,omitempty"`
}

type PaymentRequests struct {
	PaymentRequests []PaymentRequest `json:"payment_requests"`
}

type PaymentRequest struct {
	// 支払依頼ID
	ID int32 `json:"id"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 申請タイトル
	Title string `json:"title"`
	// 申請日 (yyyy-mm-dd)
	ApplicationDate string `json:"application_date"`
	// 合計金額
	TotalAmount int32 `json:"total_amount"`
	// 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し)
	Status string `json:"status"`
	// 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます)
	DealID *int32 `json:"deal_id,omitempty"`
	// 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:支払済み, unsettled:支払待ち)
	DealStatus *string `json:"deal_status,omitempty"`
	// 申請者のユーザーID
	ApplicantID int32 `json:"applicant_id"`
	// 承認者（配列） 承認ステップのresource_typeがunspecified (指定なし)の場合はapproversはレスポンスに含まれません。 しかし、resource_typeがunspecifiedの承認ステップにおいて誰かが承認・却下・差し戻しのいずれかのアクションを取った後は、 approversはレスポンスに含まれるようになります。 その場合approversにはアクションを行ったステップのIDとアクションを行ったユーザーのIDが含まれます。
	Approvers Approver `json:"approvers"`
	// 申請No.
	ApplicationNumber string `json:"application_number"`
	// 現在承認ステップID
	CurrentStepID int32 `json:"current_step_id"`
	// 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。
	CurrentRound int32 `json:"current_round"`
	// 請求書番号
	DocumentCode string `json:"document_code"`
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 支払期限 (yyyy-mm-dd)
	PaymentDate string `json:"payment_date"`
	// 支払方法(none: 指定なし, domestic_bank_transfer: 国内振込, abroad_bank_transfer: 国外振込, account_transfer: 口座振替, credit_card: クレジットカード)
	PaymentMethod string `json:"payment_method"`
	// 取引先ID
	PartnerID int32 `json:"partner_id"`
	// 取引先コード
	PartnerCode string `json:"partner_code"`
	// 取引先名
	PartnerName string `json:"partner_name"`
}

type Approver struct {
	// 承認ステップID
	StepID int32 `json:"step_id"`
	// 承認者のユーザーID 下記の場合はnullになります。
	UserID int32 `json:"user_id"`
	// 承認者の承認状態
	Status string `json:"status"`
	// 代理承認済みかどうか
	IsForceAction bool `json:"is_force_action"`
	// 承認ステップの承認方法
	ResourceType string `json:"resource_type"`
}

func (c *Client) GetPaymentRequests(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, opts interface{}) (*PaymentRequests, error) {
	var result PaymentRequests

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathPaymentRequests), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
