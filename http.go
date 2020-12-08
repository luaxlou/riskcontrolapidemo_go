package riskcontrolapidemo_go

import (
	"strconv"
	"time"

	"github.com/luaxlou/gohttpclientv2"
)

type Balance struct {
	Available int64 `json:"available"`
	Overdraft int64 `json:"overdraft"`
}

type BalanceRes struct {
	Data Balance `json:"data"`
}

func GetBalance(merchantKey string) Balance {

	uri := HOST + "/balances"

	var res BalanceRes

	gohttpclientv2.New().Get(uri).Header("MERCHANT_KEY", merchantKey).Exec().RenderJSON(&res)

	return res.Data

}

const BILL_TYPE_CONSUME = "consume"
const BILL_TYPE_DEPOSIT = "deposit"

type Bill struct {
	BillType string `json:"bill_type"`

	BeforeAmount int64  `json:"before_amount"`
	Change       int64  `json:"change"`
	Memo         string `json:"memo"`

	BillTime  time.Time `json:"bill_time"`
	CreatedAt time.Time `json:"created_at"`
}

type BillQueryRes struct {
	Count int    `json:"count"`
	Data  []Bill `json:"data"`
}

type BillRes struct {
	Data BillQueryRes
}

func GetBills(merchantKey string, offset, limit int, billType string) BillQueryRes {

	if limit > 50 {
		limit = 50
	}

	if limit == 0 {
		limit = 50
	}

	r := gohttpclientv2.New().Get(HOST+"/bills").Header("MERCHANT_KEY", merchantKey).
		Query("offset", strconv.Itoa(offset)).
		Query("limit", strconv.Itoa(limit)).
		Query("bill_type", billType)

	var bs BillRes
	r.Exec().RenderJSON(&bs)

	return bs.Data

}
