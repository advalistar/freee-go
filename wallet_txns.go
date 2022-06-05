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
	APIPathTxns = "wallet_txns"

	WalletTypeBankAccount = "bank_account"
	WalletTypeCreditCard  = "credit_card"
	WalletTypeWallet      = "wallet" // 現金

	TxnsTypeIncome  = "income"
	TxnsTypeExpense = "expense"
)

type WalletTxnsResponse struct {
	WalletTxns []WalletTxn `json:"wallet_txns"`
}

type WalletTxnResponse struct {
	WalletTxn WalletTxn `json:"wallet_txn"`
}

type GetWalletTxnOpts struct {
	// walletable_type、walletable_idは同時に指定が必要です。
	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet)
	WalletableType string `url:"walletable_type,omitempty"`
	// 口座ID
	WalletableID int32 `url:"walletable_id,omitempty"`
	// 取引日で絞込：開始日 (yyyy-mm-dd)
	StartDate string `url:"start_date,omitempty"`
	// 取引日で絞込：終了日 (yyyy-mm-dd)
	EndDate string `url:"end_date,omitempty"`
	// 入金／出金 (入金: income, 出金: expense)
	EntrySide string `url:"entry_side,omitempty"`
	// 取得レコードのオフセット (デフォルト: 0)
	Offset int32 `url:"offset,omitempty"`
	// 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 100)
	Limit int32 `url:"limit,omitempty"`
}

type WalletTxn struct {
	// 明細ID
	ID int32 `json:"id"`
	// 事業所ID
	CompanyID int32 `json:"company_id"`
	// 取引日（yyyy-mm-dd）
	Date string `json:"date"`
	// 取引金額
	Amount int32 `json:"amount"`
	// 未決済金額
	DueAmount int32 `json:"due_amount"`
	// 残高
	Balance int32 `json:"balance"`
	// 入金/出勤（入金: income, 出勤: expense）
	EntrySide string `json:"entry_side"`
	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet)
	WalletableType string `json:"walletable_type"`
	// 口座ID
	WalletableID int32 `json:"walletable_id"`
	// 取引内容
	Description string `json:"description"`
	// 明細のステータス（消込待ち: 1, 消込済み: 2, 無視: 3, 消込中: 4）
	Status uint `json:"status"`
	// 登録時に自動登録ルールの設定が適用され、登録処理が実行された場合、 trueになります。〜を推測する、〜の消込をするの条件の場合は一致してもfalseになります。
	RuleMatched bool `json:"rule_matched"`
}

func (c *Client) GetWalletTransactions(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, opts GetWalletTxnOpts) (*WalletTxnsResponse, error) {
	var result WalletTxnsResponse

	if (opts.WalletableType != "" && opts.WalletableID == 0) || (opts.WalletableID != 0 && opts.WalletableType == "") {
		return nil, fmt.Errorf("either walletable_type or walletable_id is specified, then other value must be set")
	}

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathTxns), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetWalletTransaction(ctx context.Context, reuseTokenSource oauth2.TokenSource, companyID int32, txnID int64, opts GetWalletTxnOpts) (*WalletTxn, error) {
	var result WalletTxnResponse

	if (opts.WalletableType != "" && opts.WalletableID == 0) || (opts.WalletableID != 0 && opts.WalletableType == "") {
		return nil, fmt.Errorf("either walletable_type or walletable_id is specified, then other value must be set")
	}

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	SetCompanyID(&v, companyID)
	err = c.call(ctx, path.Join(APIPathTxns, fmt.Sprint(txnID)), http.MethodGet, reuseTokenSource, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result.WalletTxn, nil
}
