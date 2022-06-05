package freee

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathJournals = "journals"
)

type GetJournalsOpts struct {
	DownloadType string   `url:"download_type"`
	VisibleTags  []string `url:"visible_tags,omitempty"`
	VisibleIDs   []string `url:"visible_ids,omitempty"`
	StartDate    string   `url:"start_date,omitempty"`
	EndDate      string   `url:"end_date,omitempty"`
}

type Journals struct {
	Journals []Journal `json:"journals"`
}

type Journal struct {
	// 受け付けID
	ID int32 `json:"id"`
	// 受け付けメッセージ
	Messages *[]string `json:"messages,omitempty"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// ダウンロード形式
	DownloadType *string `json:"download_type,omitempty"`
	// 取得開始日 (yyyy-mm-dd)
	StartDate *string `json:"start_date,omitempty"`
	// 取得終了日 (yyyy-mm-dd)
	EndDate *string `json:"end_date,omitempty"`
	// 補助科目やコメントとして出力する項目
	VisibleTags *[]string `json:"visible_tags,omitempty"`
	// 追加出力するID項目
	VisibleIDs *[]string `json:"visible_ids,omitempty"`
	// ステータス確認用URL
	StatusURL *string `json:"status_url,omitempty"`
	// 集計結果が最新かどうか
	UpToDate *bool `json:"up_to_date,omitempty"`
	// 集計結果が最新かどうか
	UpToDateReasons *[]UpToDateReason `json:"up_to_date_reasons,omitempty"`
}

type UpToDateReason struct {
	// コード
	Code string `json:"code"`
	// 集計が最新でない理由
	Message string `json:"message"`
}

func (c *Client) GetJournals(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, opts interface{}) (*Journals, error) {
	var result Journals

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathJournals), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
