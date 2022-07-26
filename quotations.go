package freee

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	APIPathQuotations = "quotations"
)

type GetQuotationsOpts struct {
	PartnerID       int32  `url:"partner_id,omitempty"`
	PartnerCode     string `url:"partner_code,omitempty"`
	StartIssueDate  string `url:"start_issue_date,omitempty"`
	EndIssueDate    string `url:"end_issue_date,omitempty"`
	QuotationNumber int32  `url:"quotation_number,omitempty"`
	Description     string `url:"description,omitempty"`
	QuotationStatus string `url:"quotation_status,omitempty"`
	Offset          int32  `url:"offset,omitempty"`
	Limit           int32  `url:"limit,omitempty"`
}

type Quotations struct {
	Quotations []Quotation `json:"quotations"`
}

type Quotation struct {
	// 見積書ID
	ID int32 `json:"id"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 見積日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 取引先ID
	PartnerID int32 `json:"partner_id"`
	// 取引先コード
	PartnerCode *string `json:"partner_code,omitempty"`
	// 見積書番号
	QuotationNumber string `json:"quotation_number"`
	// タイトル
	Title *string `json:"title,omitempty"`
	// 合計金額
	TotalAmount int32 `json:"total_amount"`
	// 消費税
	TotalVat *int32 `json:"total_vat,omitempty"`
	// 小計
	SubTotal *int32 `json:"sub_total,omitempty"`
	// 概要
	Description *string `json:"description,omitempty"`
	// 見積書ステータス (unsubmitted: 送付待ち, submitted: 送付済み, all: 全て)
	QuotationStatus string `json:"quotation_status"`
	// Web共有日時(最新)
	WebPublishedAt *string `json:"web_published_at,omitempty"`
	// Web共有ダウンロード日時(最新)
	WebDownloadedAt *string `json:"web_downloaded_at,omitempty"`
	// Web共有取引先確認日時(最新)
	WebConfirmedAt *string `json:"web_confirmed_at,omitempty"`
	// メール送信日時(最新)
	MailSentAt *string `json:"mail_sent_at,omitempty"`
	// 取引先名
	PartnerName *string `json:"partner_name,omitempty"`
	// 見積書に表示する取引先名
	PartnerDisplayName *string `json:"partner_display_name,omitempty"`
	// 敬称（御中、様、(空白)の3つから選択）
	PartnerTitle string `json:"partner_title"`
	// 郵便番号
	PartnerZipcode *string `json:"partner_zipcode,omitempty"`
	// 都道府県コード（-1: 設定しない、0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄
	PartnerPrefectureCode *int32 `json:"partner_prefecture_code,omitempty"`
	// 都道府県
	PartnerPrefectureName *string `json:"partner_prefecture_name,omitempty"`
	// 市区町村・番地
	PartnerAddress1 *string `json:"partner_address1,omitempty"`
	// 建物名・部屋番号など
	PartnerAddress2 *string `json:"partner_address2,omitempty"`
	// 取引先担当者名
	PartnerContactInfo *string `json:"partner_contact_info,omitempty"`
	// 事業所名
	CompanyName string `json:"company_name"`
	// 郵便番号
	CompanyZipcode *string `json:"company_zipcode,omitempty"`
	// 都道府県コード（-1: 設定しない、0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄
	CompanyPrefectureCode *int32 `json:"company_prefecture_code,omitempty"`
	// 都道府県
	CompanyPrefectureName *string `json:"company_prefecture_name,omitempty"`
	// 市区町村・番地
	CompanyAddress1 *string `json:"company_address1,omitempty"`
	// 建物名・部屋番号など
	CompanyAddress2 *string `json:"company_address2,omitempty"`
	// 事業所担当者名
	CompanyContactInfo *string `json:"company_contact_info,omitempty"`
	// メッセージ
	Message *string `json:"message,omitempty"`
	// 備考
	Notes *string `json:"notes,omitempty"`
	// 見積書レイアウト
	QuotationLayout string `json:"quotation_layout"`
	// 見積書の消費税計算方法(inclusive: 内税, exclusive: 外税)
	TaxEntryMethod string `json:"tax_entry_method"`
	// 見積内容
	QuotationContents *[]QuotationContent `json:"quotation_contents,omitempty"`
	// 見積内容
	TotalAmountPerVatRatetotalAmountPerVatRate *TotalAmountPerVatRatetotalAmountPerVatRate `json:"total_amount_per_vat_rate,omitempty"`
}

type QuotationContent struct {
	// 見積内容ID
	ID int32 `json:"id"`
	// 順序
	Order int32 `json:"order"`
	// 行の種類
	Type string `json:"type"`
	// 数量
	Qty float64 `json:"qty"`
	// 単位
	Unit string `json:"unit"`
	// 単価
	UnitPrice float64 `json:"unit_price"`
	// 内税/外税の判別とamountの税込み、税抜きについて
	Amount int32 `json:"amount"`
	// 消費税額
	Vat int32 `json:"vat"`
	// 軽減税率税区分（true: 対象、false: 対象外）
	ReducedVat bool `json:"reduced_vat"`
	// 備考
	Description string `json:"description"`
	// 勘定科目ID
	AccountItemID int32 `json:"account_item_id"`
	// 勘定科目名
	AccountItemName string `json:"account_item_name"`
	// 税区分コード
	TaxCode int32 `json:"tax_code"`
	// 品目ID
	ItemID int32 `json:"item_id"`
	// 品目
	ItemName string `json:"item_name"`
	// 部門ID
	SectionID int32 `json:"section_id"`
	// 部門
	SectionName string `json:"section_name"`
	// メモタグID
	TagIDs []int32 `json:"tag_ids"`
	// メモタグ
	TagNames []string `json:"tag_names"`
	// セグメント１ID
	Segment1TagID *string `json:"segment_1_tag_id,omitempty"`
	// セグメント１
	Segment1TagName *string `json:"segment_1_tag_name,omitempty"`
	// セグメント２ID
	Segment2TagID *string `json:"segment_2_tag_id,omitempty"`
	// セグメント２
	Segment2TagName *string `json:"segment_2_tag_name,omitempty"`
	// セグメント３ID
	Segment3TagID *string `json:"segment_3_tag_id,omitempty"`
	// セグメント３
	Segment3TagName *string `json:"segment_3_tag_name,omitempty"`
}

type TotalAmountPerVatRatetotalAmountPerVatRate struct {
	// 税率5%の税込み金額合計
	Vat5 uint64 `json:"vat_5"`
	// 税率8%の税込み金額合計
	Vat8 uint64 `json:"vat_8"`
	// 軽減税率8%の税込み金額合計
	ReducedVat8 uint64 `json:"reduced_vat_8"`
	// 税率10%の税込み金額合計
	Vat10 uint64 `json:"vat_10"`
}

func (c *Client) GetQuotations(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, opts interface{}) (*Quotations, error) {
	var result Quotations

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathQuotations), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
