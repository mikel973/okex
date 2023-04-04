package account

import (
	"github.com/mikel973/okex"
)

type (
	Balance struct {
		TotalEq     test.JSONFloat64  `json:"totalEq"`
		IsoEq       test.JSONFloat64  `json:"isoEq"`
		AdjEq       test.JSONFloat64  `json:"adjEq,omitempty"`
		OrdFroz     test.JSONFloat64  `json:"ordFroz,omitempty"`
		Imr         test.JSONFloat64  `json:"imr,omitempty"`
		Mmr         test.JSONFloat64  `json:"mmr,omitempty"`
		MgnRatio    test.JSONFloat64  `json:"mgnRatio,omitempty"`
		NotionalUsd test.JSONFloat64  `json:"notionalUsd,omitempty"`
		Details     []*BalanceDetails `json:"details,omitempty"`
		UTime       test.JSONTime     `json:"uTime"`
	}
	BalanceDetails struct {
		Ccy           string           `json:"ccy"`
		Eq            test.JSONFloat64 `json:"eq"`
		CashBal       test.JSONFloat64 `json:"cashBal"`
		IsoEq         test.JSONFloat64 `json:"isoEq,omitempty"`
		AvailEq       test.JSONFloat64 `json:"availEq,omitempty"`
		DisEq         test.JSONFloat64 `json:"disEq"`
		AvailBal      test.JSONFloat64 `json:"availBal"`
		FrozenBal     test.JSONFloat64 `json:"frozenBal"`
		OrdFrozen     test.JSONFloat64 `json:"ordFrozen"`
		Liab          test.JSONFloat64 `json:"liab,omitempty"`
		Upl           test.JSONFloat64 `json:"upl,omitempty"`
		UplLib        test.JSONFloat64 `json:"uplLib,omitempty"`
		CrossLiab     test.JSONFloat64 `json:"crossLiab,omitempty"`
		IsoLiab       test.JSONFloat64 `json:"isoLiab,omitempty"`
		MgnRatio      test.JSONFloat64 `json:"mgnRatio,omitempty"`
		Interest      test.JSONFloat64 `json:"interest,omitempty"`
		Twap          test.JSONFloat64 `json:"twap,omitempty"`
		MaxLoan       test.JSONFloat64 `json:"maxLoan,omitempty"`
		EqUsd         test.JSONFloat64 `json:"eqUsd"`
		NotionalLever test.JSONFloat64 `json:"notionalLever,omitempty"`
		StgyEq        test.JSONFloat64 `json:"stgyEq"`
		IsoUpl        test.JSONFloat64 `json:"isoUpl,omitempty"`
		UTime         test.JSONTime    `json:"uTime"`
	}
	Position struct {
		InstID      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		LiabCcy     string              `json:"liabCcy,omitempty"`
		OptVal      string              `json:"optVal,omitempty"`
		Ccy         string              `json:"ccy"`
		PosID       string              `json:"posId"`
		TradeID     string              `json:"tradeId"`
		Pos         test.JSONFloat64    `json:"pos"`
		AvailPos    test.JSONFloat64    `json:"availPos,omitempty"`
		AvgPx       test.JSONFloat64    `json:"avgPx"`
		Upl         test.JSONFloat64    `json:"upl"`
		UplRatio    test.JSONFloat64    `json:"uplRatio"`
		Lever       test.JSONFloat64    `json:"lever"`
		LiqPx       test.JSONFloat64    `json:"liqPx,omitempty"`
		Imr         test.JSONFloat64    `json:"imr,omitempty"`
		Margin      test.JSONFloat64    `json:"margin,omitempty"`
		MgnRatio    test.JSONFloat64    `json:"mgnRatio"`
		Mmr         test.JSONFloat64    `json:"mmr"`
		Liab        test.JSONFloat64    `json:"liab,omitempty"`
		Interest    test.JSONFloat64    `json:"interest"`
		NotionalUsd test.JSONFloat64    `json:"notionalUsd"`
		ADL         test.JSONFloat64    `json:"adl"`
		Last        test.JSONFloat64    `json:"last"`
		DeltaBS     test.JSONFloat64    `json:"deltaBS"`
		DeltaPA     test.JSONFloat64    `json:"deltaPA"`
		GammaBS     test.JSONFloat64    `json:"gammaBS"`
		GammaPA     test.JSONFloat64    `json:"gammaPA"`
		ThetaBS     test.JSONFloat64    `json:"thetaBS"`
		ThetaPA     test.JSONFloat64    `json:"thetaPA"`
		VegaBS      test.JSONFloat64    `json:"vegaBS"`
		VegaPA      test.JSONFloat64    `json:"vegaPA"`
		PosSide     test.PositionSide   `json:"posSide"`
		MgnMode     test.MarginMode     `json:"mgnMode"`
		InstType    test.InstrumentType `json:"instType"`
		CTime       test.JSONTime       `json:"cTime"`
		UTime       test.JSONTime       `json:"uTime"`
	}
	BalanceAndPosition struct {
		EventType test.EventType    `json:"eventType"`
		PTime     test.JSONTime     `json:"pTime"`
		UTime     test.JSONTime     `json:"uTime"`
		PosData   []*Position       `json:"posData"`
		BalData   []*BalanceDetails `json:"balData"`
	}
	PositionAndAccountRisk struct {
		AdjEq   test.JSONFloat64                     `json:"adjEq,omitempty"`
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		TS      test.JSONTime                        `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string           `json:"ccy"`
		Eq    test.JSONFloat64 `json:"eq"`
		DisEq test.JSONFloat64 `json:"disEq"`
	}
	PositionAndAccountRiskPositionData struct {
		InstID      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		Ccy         string              `json:"ccy"`
		NotionalCcy test.JSONFloat64    `json:"notionalCcy"`
		Pos         test.JSONFloat64    `json:"pos"`
		NotionalUsd test.JSONFloat64    `json:"notionalUsd"`
		PosSide     test.PositionSide   `json:"posSide"`
		InstType    test.InstrumentType `json:"instType"`
		MgnMode     test.MarginMode     `json:"mgnMode"`
	}
	Bill struct {
		Ccy       string              `json:"ccy"`
		InstID    string              `json:"instId"`
		Notes     string              `json:"notes"`
		BillID    string              `json:"billId"`
		OrdID     string              `json:"ordId"`
		BalChg    test.JSONFloat64    `json:"balChg"`
		PosBalChg test.JSONFloat64    `json:"posBalChg"`
		Bal       test.JSONFloat64    `json:"bal"`
		PosBal    test.JSONFloat64    `json:"posBal"`
		Sz        test.JSONFloat64    `json:"sz"`
		Pnl       test.JSONFloat64    `json:"pnl"`
		Fee       test.JSONFloat64    `json:"fee"`
		From      test.AccountType    `json:"from,string"`
		To        test.AccountType    `json:"to,string"`
		InstType  test.InstrumentType `json:"instType"`
		MgnMode   test.MarginMode     `json:"MgnMode"`
		Type      test.BillType       `json:"type,string"`
		SubType   test.BillSubType    `json:"subType,string"`
		TS        test.JSONTime       `json:"ts"`
	}
	Config struct {
		Level      string            `json:"level"`
		LevelTmp   string            `json:"levelTmp"`
		AcctLv     string            `json:"acctLv"`
		AutoLoan   bool              `json:"autoLoan"`
		UID        string            `json:"uid"`
		GreeksType test.GreekType    `json:"greeksType"`
		PosMode    test.PositionType `json:"posMode"`
	}
	PositionMode struct {
		PosMode test.PositionType `json:"posMode"`
	}
	Leverage struct {
		InstID  string            `json:"instId"`
		Lever   test.JSONFloat64  `json:"lever"`
		MgnMode test.MarginMode   `json:"mgnMode"`
		PosSide test.PositionSide `json:"posSide"`
	}
	MaxBuySellAmount struct {
		InstID  string           `json:"instId"`
		Ccy     string           `json:"ccy"`
		MaxBuy  test.JSONFloat64 `json:"maxBuy"`
		MaxSell test.JSONFloat64 `json:"maxSell"`
	}
	MaxAvailableTradeAmount struct {
		InstID    string           `json:"instId"`
		AvailBuy  test.JSONFloat64 `json:"availBuy"`
		AvailSell test.JSONFloat64 `json:"availSell"`
	}
	MarginBalanceAmount struct {
		InstID  string            `json:"instId"`
		Amt     test.JSONFloat64  `json:"amt"`
		PosSide test.PositionSide `json:"posSide,string"`
		Type    test.CountAction  `json:"type,string"`
	}
	Loan struct {
		InstID  string           `json:"instId"`
		MgnCcy  string           `json:"mgnCcy"`
		Ccy     string           `json:"ccy"`
		MaxLoan test.JSONFloat64 `json:"maxLoan"`
		MgnMode test.MarginMode  `json:"mgnMode"`
		Side    test.OrderSide   `json:"side,string"`
	}
	Fee struct {
		Level    string              `json:"level"`
		Taker    test.JSONFloat64    `json:"taker"`
		Maker    test.JSONFloat64    `json:"maker"`
		Delivery test.JSONFloat64    `json:"delivery,omitempty"`
		Exercise test.JSONFloat64    `json:"exercise,omitempty"`
		Category test.FeeCategory    `json:"category,string"`
		InstType test.InstrumentType `json:"instType"`
		TS       test.JSONTime       `json:"ts"`
	}
	InterestAccrued struct {
		InstID       string           `json:"instId"`
		Ccy          string           `json:"ccy"`
		Interest     test.JSONFloat64 `json:"interest"`
		InterestRate test.JSONFloat64 `json:"interestRate"`
		Liab         test.JSONFloat64 `json:"liab"`
		MgnMode      test.MarginMode  `json:"mgnMode"`
		TS           test.JSONTime    `json:"ts"`
	}
	InterestRate struct {
		Ccy          string           `json:"ccy"`
		InterestRate test.JSONFloat64 `json:"interestRate"`
	}
	Greek struct {
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		Ccy   string           `json:"ccy"`
		MaxWd test.JSONFloat64 `json:"maxWd"`
	}
)
