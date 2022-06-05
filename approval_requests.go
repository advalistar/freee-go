package freee

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathApprovalRequests = "approval_requests"
)

type GetApprovalRequestsOpts struct {
	Status               string `url:"status,omitempty"`
	ApplicationNumber    int32  `url:"application_number,omitempty"`
	Title                string `url:"title,omitempty"`
	FormID               int32  `url:"form_id,omitempty"`
	StartApplicationDate string `url:"start_application_date,omitempty"`
	EndApplicationDate   string `url:"end_application_date,omitempty"`
	ApplicantID          int32  `url:"applicant_id,omitempty"`
	MinAmount            int32  `url:"min_amount,omitempty"`
	MaxAmount            int32  `url:"max_amount,omitempty"`
	ApproverID           int32  `url:"approver_id,omitempty"`
	Offset               int32  `url:"offset,omitempty"`
	Limit                int32  `url:"limit,omitempty"`
}

type ApprovalRequests struct {
	ApprovalRequests []ApprovalRequest `json:"approval_requests"`
}

type ApprovalRequest struct {
	// 各種申請ID
	ID int32 `json:"id"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 申請日 (yyyy-mm-dd)
	ApplicationDate string `json:"application_date"`
	// 申請タイトル
	Title string `json:"title"`
	// 申請者のユーザーID
	ApplicantID int32 `json:"applicant_id"`
	// 申請No.
	ApplicationNumber string `json:"application_number"`
	// 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し)
	Status string `json:"status"`
	// 各種申請の項目一覧（配列）
	RequestItems []RequestItem `json:"request_items"`
	// 申請フォームID
	FormID int32 `json:"form_id"`
	// 現在承認ステップID
	CurrentStepID int32 `json:"current_step_id"`
	// 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。
	CurrentRound int32 `json:"current_round"`
	// 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます)
	DealID int32 `json:"deal_id"`
	// 振替伝票のID (申請ステータス:statusがapprovedで、関連する振替伝票が存在する時のみmanual_journal_idが表示されます)
	ManualJournalID int32 `json:"manual_journal_id"`
	// 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:決済済み, unsettled:未決済)
	DealStatus string `json:"deal_status"`
}

type RequestItem struct {
	// 項目ID
	ID int32 `json:"id"`
	// 項目タイプ(title: 申請タイトル, single_line: 自由記述形式 1行, multi_line: 自由記述形式 複数行, select: プルダウン, date: 日付, amount: 金額, receipt: 添付ファイル, section: 部門ID, partner: 取引先ID, ninja_sign_document: 契約書（freeeサイン連携）)
	Type string `json:"type"`
	// 項目の値
	Value string `json:"value"`
}

type ApprovalRequestsForms struct {
	ApprovalRequestsForms []ApprovalRequestsForm `json:"approval_request_forms"`
}

type ApprovalRequestsForm struct {
	// 申請フォームID
	ID int32 `json:"id"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 申請フォームの名前
	Name string `json:"name"`
	// 申請フォームの説明
	Description string `json:"description"`
	// ステータス(draft: 申請で使用しない、active: 申請で使用する、deleted: 削除済み)
	Status string `json:"status"`
	// 作成日時
	CreatedDate string `json:"created_date"`
	// 表示順（申請者が選択する申請フォームの表示順を設定できます。小さい数ほど上位に表示されます。（0を除く整数のみ。マイナス不可）未入力の場合、表示順が後ろになります。同じ数字が入力された場合、登録順で表示されます。）
	FormOrder int32 `json:"form_order"`
	// 適用された経路数
	RouteSettingCount int32 `json:"route_setting_count"`
}

func (c *Client) GetApprovalRequests(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, opts interface{}) (*ApprovalRequests, error) {
	var result ApprovalRequests

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, APIPathApprovalRequests, http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetApprovalRequestsForms(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32) (*ApprovalRequestsForms, error) {
	var result ApprovalRequestsForms

	v, err := query.Values(nil)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathApprovalRequests, "forms"), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
