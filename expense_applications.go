package freee

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathExpenseApplications = "expense_applications"
)

type GetExpenseApplicationsOpts struct {
	Status               string `url:"status,omitempty"`
	PayrollAttached      bool   `url:"payroll_attached,omitempty"`
	StartTransactionDate string `url:"start_transaction_date,omitempty"`
	EndTransactionDate   string `url:"end_transaction_date,omitempty"`
	ApplicationNumber    int32  `url:"application_number,omitempty"`
	Title                string `url:"title,omitempty"`
	StartIssueDate       string `url:"start_issue_date,omitempty"`
	EndIssueDate         string `url:"end_issue_date,omitempty"`
	ApplicantID          int32  `url:"applicant_id,omitempty"`
	ApproverID           int32  `url:"approver_id,omitempty"`
	MinAmount            int32  `url:"min_amount,omitempty"`
	MaxAmount            int32  `url:"max_amount,omitempty"`
	Offset               int32  `url:"offset,omitempty"`
	Limit                int32  `url:"limit,omitempty"`
}

type ExpenseApplications struct {
	ExpenseApplications []ExpenseApplication `json:"expense_applications"`
}

type ExpenseApplication struct {
	// 経費申請ID
	ID int32 `json:"id"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 申請タイトル
	Title string `json:"title"`
	// 申請日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 備考
	Description *string `json:"description,omitempty"`
	// 合計金額
	TotalAmount *int32 `json:"total_amount,omitempty"`
	// 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し)
	Status string `json:"status"`
	// 部門ID
	SectionID *int32 `json:"section_id,omitempty"`
	// メモタグID
	TagIDs *[]int32 `json:"tag_ids,omitempty"`
	// 経費申請の項目行一覧（配列）
	ExpenseApplicationLines []ExpenseApplicationLine `json:"expense_application_lines"`
	// 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます)
	DealID int32 `json:"deal_id"`
	// 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:精算済み, unsettled:清算待ち)
	DealStatus string `json:"deal_status"`
	// 申請者のユーザーID
	ApplicantID int32 `json:"applicant_id"`
	// 申請No.
	ApplicationNumber string `json:"application_number"`
	// 現在承認ステップID
	CurrentStepID *int32 `json:"current_step_id,omitempty"`
	// 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。
	CurrentRound *int32 `json:"current_round,omitempty"`
	// セグメント１ID
	Segment1TagID *int32 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID
	Segment2TagID *int32 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID
	Segment3TagID *int32 `json:"segment_3_tag_id,omitempty"`
}

type ExpenseApplicationLine struct {
	// 経費申請の項目行ID
	ID int32 `json:"id"`
	// 日付 (yyyy-mm-dd)
	TransactionDate *string `json:"transaction_date,omitempty"`
	// 内容
	Description *string `json:"description,omitempty"`
	// 金額
	Amount *int32 `json:"amount,omitempty"`
	// 経費科目ID
	ExpenseApplicationLineTemplateID *int32 `json:"expense_application_line_template_id,omitempty"`
	// 証憑ファイルID（ファイルボックスのファイルID）
	ReceiptID *int32 `json:"receipt_id,omitempty"`
}

func (c *Client) GetExpenseApplications(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, opts interface{}) (*ExpenseApplications, error) {
	var result ExpenseApplications

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, APIPathExpenseApplications, http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
