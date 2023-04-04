package funding

import "github.com/mikel973/okex"

type (
	Currency struct {
		Ccy         string `json:"ccy"`
		Name        string `json:"name"`
		Chain       string `json:"chain"`
		MinWd       string `json:"minWd"`
		MinFee      string `json:"minFee"`
		MaxFee      string `json:"maxFee"`
		CanDep      bool   `json:"canDep"`
		CanWd       bool   `json:"canWd"`
		CanInternal bool   `json:"canInternal"`
	}
	Balance struct {
		Ccy       string `json:"ccy"`
		Bal       string `json:"bal"`
		FrozenBal string `json:"frozenBal"`
		AvailBal  string `json:"availBal"`
	}
	Transfer struct {
		TransID string           `json:"transId"`
		Ccy     string           `json:"ccy"`
		Amt     test.JSONFloat64 `json:"amt"`
		From    test.AccountType `json:"from,string"`
		To      test.AccountType `json:"to,string"`
	}
	Bill struct {
		BillID string           `json:"billId"`
		Ccy    string           `json:"ccy"`
		Bal    test.JSONFloat64 `json:"bal"`
		BalChg test.JSONFloat64 `json:"balChg"`
		Type   test.BillType    `json:"type,string"`
		TS     test.JSONTime    `json:"ts"`
	}
	DepositAddress struct {
		Addr     string           `json:"addr"`
		Tag      string           `json:"tag,omitempty"`
		Memo     string           `json:"memo,omitempty"`
		PmtID    string           `json:"pmtId,omitempty"`
		Ccy      string           `json:"ccy"`
		Chain    string           `json:"chain"`
		CtAddr   string           `json:"ctAddr"`
		Selected bool             `json:"selected"`
		To       test.AccountType `json:"to,string"`
		TS       test.JSONTime    `json:"ts"`
	}
	DepositHistory struct {
		Ccy   string            `json:"ccy"`
		Chain string            `json:"chain"`
		TxID  string            `json:"txId"`
		From  string            `json:"from"`
		To    string            `json:"to"`
		DepId string            `json:"depId"`
		Amt   test.JSONFloat64  `json:"amt"`
		State test.DepositState `json:"state,string"`
		TS    test.JSONTime     `json:"ts"`
	}
	Withdrawal struct {
		Ccy   string           `json:"ccy"`
		Chain string           `json:"chain"`
		WdID  test.JSONInt64   `json:"wdId"`
		Amt   test.JSONFloat64 `json:"amt"`
	}
	WithdrawalHistory struct {
		Ccy   string               `json:"ccy"`
		Chain string               `json:"chain"`
		TxID  string               `json:"txId"`
		From  string               `json:"from"`
		To    string               `json:"to"`
		Tag   string               `json:"tag,omitempty"`
		PmtID string               `json:"pmtId,omitempty"`
		Memo  string               `json:"memo,omitempty"`
		Amt   test.JSONFloat64     `json:"amt"`
		Fee   test.JSONInt64       `json:"fee"`
		WdID  test.JSONInt64       `json:"wdId"`
		State test.WithdrawalState `json:"state,string"`
		TS    test.JSONTime        `json:"ts"`
	}
	PiggyBank struct {
		Ccy  string           `json:"ccy"`
		Amt  test.JSONFloat64 `json:"amt"`
		Side test.ActionType  `json:"side,string"`
	}
	PiggyBankBalance struct {
		Ccy      string           `json:"ccy"`
		Amt      test.JSONFloat64 `json:"amt"`
		Earnings test.JSONFloat64 `json:"earnings"`
	}
)
