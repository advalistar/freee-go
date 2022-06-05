package freee

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathReports = "reports"
)

type GetReportsOpts struct {
	// 会計年度
	FiscalYear int32 `url:"fiscal_year,omitempty"`
	// 発生月で絞込：開始会計月(1-12)。指定されない場合、現在の会計年度の期首月が指定されます。
	StartMonth int32 `url:"start_month,omitempty"`
	// 発生月で絞込：終了会計月(1-12)(会計年度が10月始まりでstart_monthが11なら11, 12, 1, ... 9のいずれかを指定する)。指定されない場合、現在の会計年度の期末月が指定されます。
	EndMonth int32 `url:"end_month,omitempty"`
	// 発生日で絞込：開始日(yyyy-mm-dd)
	StartDate string `url:"start_date,omitempty"`
	// 発生日で絞込：終了日(yyyy-mm-dd)
	EndDate string `url:"end_date,omitempty"`
	// 勘定科目の表示（勘定科目: account_item, 決算書表示:group）。指定されない場合、勘定科目: account_itemが指定されます。
	AccountItemDisplayType string `url:"account_item_display_type,omitempty"`
	// 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag） ※勘定科目はaccount_item_display_typeが「group」の時のみ指定できます
	BreakdownDisplayType string `url:"breakdown_display_type,omitempty"`
	// 取引先IDで絞込（0を指定すると、取引先が未選択で絞り込めます）
	PartnerID int32 `url:"partner_id,omitempty"`
	// 取引先コードで絞込（事業所設定で取引先コードの利用を有効にしている場合のみ利用可能です）
	PartnerCode string `url:"partner_code,omitempty"`
	// 品目IDで絞込（0を指定すると、品目が未選択で絞り込めます）
	ItemID int32 `url:"item_id,omitempty"`
	// 部門IDで絞込（0を指定すると、部門が未選択で絞り込めます）
	SectionID int32 `url:"section_id,omitempty"`
	// 決算整理仕訳で絞込（決算整理仕訳のみ: only, 決算整理仕訳以外: without）。指定されない場合、決算整理仕訳以外: withoutが指定されます。
	Adjustment string `url:"adjustment,omitempty"`
	// 配賦仕訳で絞込（配賦仕訳のみ：only,配賦仕訳以外：without）。指定されない場合、配賦仕訳を含む金額が返却されます。
	CostAllocation string `url:"cost_allocation,omitempty"`
	// 承認ステータスで絞込 (未承認を除く: without_in_progress (デフォルト)、全てのステータス: all)
	// 個人: プレミアムプラン、法人: プロフェッショナルプラン以上で指定可能です。
	// 事業所の設定から仕訳承認フローの利用を有効にした場合に指定可能です。
	ApprovalFlowStatus string `url:"approval_flow_status,omitempty"`
}

type TrialBSResponse struct {
	TrialBS Report `json:"trial_bs"`
}

type Report struct {
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 会計年度(条件に指定した時、または条件に月、日条件がない時のみ含まれる）
	FiscalYear *int32 `json:"fiscal_year,omitempty"`
	// 発生月で絞込：開始会計月(1-12)(条件に指定した時のみ含まれる）
	StartMonth *int32 `json:"start_month,omitempty"`
	// 発生月で絞込：終了会計月(1-12)(条件に指定した時のみ含まれる）
	EndMonth *int32 `json:"end_month,omitempty"`
	// 発生日で絞込：開始日(yyyy-mm-dd)(条件に指定した時のみ含まれる）
	StartDate *string `json:"start_date,omitempty"`
	// 発生日で絞込：開始日(yyyy-mm-dd)(条件に指定した時のみ含まれる）
	EndDate *string `json:"end_date,omitempty"`
	// 勘定科目の表示（勘定科目: account_item, 決算書表示:group）(条件に指定した時のみ含まれる）
	AccountItemDisplayType *string `json:"account_item_display_type,omitempty"`
	// 内訳の表示（取引先: partner, 品目: item, 部門: section, 勘定科目: account_item, セグメント1(法人向けプロフェッショナル, 法人向けエンタープライズプラン): segment_1_tag, セグメント2(法人向け エンタープライズプラン):segment_2_tag, セグメント3(法人向け エンタープライズプラン): segment_3_tag）(条件に指定した時のみ含まれる）
	BreakdownDisplayType *string `json:"breakdown_display_type,omitempty"`
	// 取引先ID(条件に指定した時のみ含まれる）
	PartnerID *int32 `json:"partner_id,omitempty"`
	// 取引先コード(条件に指定した時のみ含まれる）
	PartnerCode *string `json:"partner_code,omitempty"`
	// 品目ID(条件に指定した時のみ含まれる）
	ItemID *int32 `json:"item_id,omitempty"`
	// 部門ID(条件に指定した時のみ含まれる）
	SectionID *int32 `json:"section_id,omitempty"`
	// 決算整理仕訳のみ: only, 決算整理仕訳以外: without(条件に指定した時のみ含まれる）
	Adjustment *string `json:"adjustment,omitempty"`
	// 未承認を除く: without_in_progress (デフォルト), 全てのステータス: all(条件に指定した時のみ含まれる）
	ApprovalFlowStatus *string `json:"approval_flow_status,omitempty"`
	// 作成日時
	CReatedAt *string   `json:"created_at,omitempty"`
	Balances  []Balance `json:"balances,omitempty"`
	// 集計結果が最新かどうか
	UpToDate bool `json:"up_to_date"`
	// 作成日時
	UpToDateReasons *[]UpToDateReason `json:"up_to_date_reasons,omitempty"`

	// 配賦仕訳のみ：only,配賦仕訳以外：without(条件に指定した時のみ含まれる）
	CostAllocation *string `json:"cost_allocation,omitempty"`
}

type Balance struct {
	// 勘定科目ID(勘定科目の時のみ含まれる)
	AccountItemID *int32 `json:"account_item_id,omitempty"`
	// 勘定科目名(勘定科目の時のみ含まれる)
	AccountItemName *string `json:"account_item_name,omitempty"`
	// 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる)
	AccountGroupName *string `json:"account_group_name,omitempty"`
	// breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる
	Partners *[]BalanceBreakdown `json:"partners,omitempty"`
	// breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる
	Items *[]BalanceBreakdown `json:"items,omitempty"`
	// breakdown_display_type:section, account_item_display_type:account_item指定時のみ含まれる
	Sections *[]BalanceBreakdown `json:"sections,omitempty"`
	// breakdown_display_type:segment_1_tag, account_item_display_type:account_item指定時のみ含まれる
	Segment1Tags *[]BalanceBreakdown `json:"segment_1_tags,omitempty"`
	// breakdown_display_type:segment_2_tag, account_item_display_type:account_item指定時のみ含まれる
	Segment2Tags *[]BalanceBreakdown `json:"segment_2_tags,omitempty"`
	// breakdown_display_type:segment_3_tag, account_item_display_type:account_item指定時のみ含まれる
	Segment3Tags *[]BalanceBreakdown `json:"segment_3_tags,omitempty"`
	// 勘定科目カテゴリー名
	AccountCategoryName *string `json:"account_category_name,omitempty"`
	// 合計行(勘定科目カテゴリーの時のみ含まれる)
	TotalLine *bool `json:"total_line,omitempty"`
	// 階層レベル
	HierarchyLevel *int32 `json:"hierarchy_level,omitempty"`
	// 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる)
	ParentAccountCategoryName *string `json:"parent_account_category_name,omitempty"`
	// 期首残高
	OpeningBalance *int32 `json:"opening_balance,omitempty"`
	// 借方金額
	DebitAmount *int32 `json:"debit_amount,omitempty"`
	// 貸方金額
	CReditAmount *int32 `json:"credit_amount,omitempty"`
	// 期末残高
	ClosingBalance *int32 `json:"closing_balance,omitempty"`
	// 構成比
	CompositionRatio *float64 `json:"composition_ratio,omitempty"`

	// 前年度期末残高
	LastYearClosingBalance *int32 `json:"last_year_closing_balance,omitempty"`
	// 前年比
	YearOnYear *float64 `json:"year_on_year,omitempty"`

	// 前々年度期末残高
	TwoYearsBeforeClosingBalance *int32 `json:"two_years_before_closing_balance,omitempty"`
}

type BalanceBreakdown struct {
	ID   int32   `json:"id"`
	Name *string `json:"name,omitempty"`
	// 期首残高
	OpeningBalance *int32 `json:"opening_balance,omitempty"`
	// 借方金額
	DebitAmount *int32 `json:"debit_amount,omitempty"`
	// 貸方金額
	CReditAmount *int32 `json:"credit_amount,omitempty"`
	// 期末残高
	ClosingBalance *int32 `json:"closing_balance,omitempty"`
	// 構成比
	CompositionRatio *float64 `json:"composition_ratio,omitempty"`

	// 前年度期末残高
	LastYearClosingBalance *int32 `json:"last_year_closing_balance,omitempty"`
	// 前年比
	YearOnYear *float64 `json:"year_on_year,omitempty"`

	// 前々年度期末残高
	TwoYearsBeforeClosingBalance *int32 `json:"two_years_before_closing_balance,omitempty"`
}

type TrialBSTwoYearsResponse struct {
	TrialBSTwoYears Report `json:"trial_bs_two_years"`
}

type TrialBSThreeYearsResponse struct {
	TrialBSThreeYears Report `json:"trial_bs_three_years"`
}

type TrialPLResponse struct {
	TrialPL Report `json:"trial_pl"`
}

type TrialPLTwoYearsResponse struct {
	TrialPLTwoYears Report `json:"trial_pl_two_years"`
}

type TrialPLThreeYearsResponse struct {
	TrialPLThreeYears Report `json:"trial_pl_three_years"`
}

type TrialCRResponse struct {
	TrialCR Report `json:"trial_cr"`
}

type TrialCRTwoYearsResponse struct {
	TrialCRTwoYears Report `json:"trial_cr_two_years"`
}

type TrialCRThreeYearsResponse struct {
	TrialCRThreeYears Report `json:"trial_cr_three_years"`
}

func (c *Client) GetTrialBS(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32) (*TrialBSResponse, error) {
	var result TrialBSResponse

	v, err := query.Values(nil)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathReports, "trial_bs"), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTrialBSTwoYears(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32) (*TrialBSTwoYearsResponse, error) {
	var result TrialBSTwoYearsResponse

	v, err := query.Values(nil)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathReports, "trial_bs_two_years"), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTrialBSThreeYears(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32) (*TrialBSThreeYearsResponse, error) {
	var result TrialBSThreeYearsResponse

	v, err := query.Values(nil)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathReports, "trial_bs_three_years"), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTrialPL(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32) (*TrialPLResponse, error) {
	var result TrialPLResponse

	v, err := query.Values(nil)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathReports, "trial_pl"), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTrialPLTwoYears(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32) (*TrialPLTwoYearsResponse, error) {
	var result TrialPLTwoYearsResponse

	v, err := query.Values(nil)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathReports, "trial_pl_two_years"), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTrialPLThreeYears(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32) (*TrialPLThreeYearsResponse, error) {
	var result TrialPLThreeYearsResponse

	v, err := query.Values(nil)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathReports, "trial_pl_three_years"), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTrialCR(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32) (*TrialCRResponse, error) {
	var result TrialCRResponse

	v, err := query.Values(nil)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathReports, "trial_cr"), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTrialCRTwoYears(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32) (*TrialCRTwoYearsResponse, error) {
	var result TrialCRTwoYearsResponse

	v, err := query.Values(nil)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathReports, "trial_cr_two_years"), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTrialCRThreeYears(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32) (*TrialCRThreeYearsResponse, error) {
	var result TrialCRThreeYearsResponse

	v, err := query.Values(nil)
	if err != nil {
		return nil, err
	}
	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathReports, "trial_cr_three_years"), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
