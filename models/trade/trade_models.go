package trade

import (
	"github.com/mikel973/okex"
)

type (
	PlaceOrder struct {
		ClOrdID string           `json:"clOrdId"`
		Tag     string           `json:"tag"`
		SMsg    string           `json:"sMsg"`
		SCode   test.JSONInt64   `json:"sCode"`
		OrdID   test.JSONFloat64 `json:"ordId"`
	}
	CancelOrder struct {
		OrdID   string           `json:"ordId"`
		ClOrdID string           `json:"clOrdId"`
		SMsg    string           `json:"sMsg"`
		SCode   test.JSONFloat64 `json:"sCode"`
	}
	AmendOrder struct {
		OrdID   string           `json:"ordId"`
		ClOrdID string           `json:"clOrdId"`
		ReqID   string           `json:"reqId"`
		SMsg    string           `json:"sMsg"`
		SCode   test.JSONFloat64 `json:"sCode"`
	}
	ClosePosition struct {
		InstID  string            `json:"instId"`
		PosSide test.PositionSide `json:"posSide"`
	}
	Order struct {
		InstID      string              `json:"instId"`
		Ccy         string              `json:"ccy"`
		OrdID       string              `json:"ordId"`
		ClOrdID     string              `json:"clOrdId"`
		TradeID     string              `json:"tradeId"`
		Tag         string              `json:"tag"`
		Category    string              `json:"category"`
		FeeCcy      string              `json:"feeCcy"`
		RebateCcy   string              `json:"rebateCcy"`
		Px          test.JSONFloat64    `json:"px"`
		Sz          test.JSONInt64      `json:"sz"`
		Pnl         test.JSONFloat64    `json:"pnl"`
		AccFillSz   test.JSONInt64      `json:"accFillSz"`
		FillPx      test.JSONFloat64    `json:"fillPx"`
		FillSz      test.JSONInt64      `json:"fillSz"`
		FillTime    test.JSONFloat64    `json:"fillTime"`
		AvgPx       test.JSONFloat64    `json:"avgPx"`
		Lever       test.JSONFloat64    `json:"lever"`
		TpTriggerPx test.JSONFloat64    `json:"tpTriggerPx"`
		TpOrdPx     test.JSONFloat64    `json:"tpOrdPx"`
		SlTriggerPx test.JSONFloat64    `json:"slTriggerPx"`
		SlOrdPx     test.JSONFloat64    `json:"slOrdPx"`
		Fee         test.JSONFloat64    `json:"fee"`
		Rebate      test.JSONFloat64    `json:"rebate"`
		State       test.OrderState     `json:"state"`
		TdMode      test.TradeMode      `json:"tdMode"`
		PosSide     test.PositionSide   `json:"posSide"`
		Side        test.OrderSide      `json:"side"`
		OrdType     test.OrderType      `json:"ordType"`
		InstType    test.InstrumentType `json:"instType"`
		TgtCcy      test.QuantityType   `json:"tgtCcy"`
		UTime       test.JSONTime       `json:"uTime"`
		CTime       test.JSONTime       `json:"cTime"`
	}
	TransactionDetail struct {
		InstID   string              `json:"instId"`
		OrdID    string              `json:"ordId"`
		TradeID  string              `json:"tradeId"`
		ClOrdID  string              `json:"clOrdId"`
		BillID   string              `json:"billId"`
		Tag      test.JSONFloat64    `json:"tag"`
		FillPx   test.JSONFloat64    `json:"fillPx"`
		FillSz   test.JSONFloat64    `json:"fillSz"`
		FeeCcy   test.JSONFloat64    `json:"feeCcy"`
		Fee      test.JSONFloat64    `json:"fee"`
		InstType test.InstrumentType `json:"instType"`
		Side     test.OrderSide      `json:"side"`
		PosSide  test.PositionSide   `json:"posSide"`
		ExecType test.OrderFlowType  `json:"execType"`
		TS       test.JSONTime       `json:"ts"`
	}
	PlaceAlgoOrder struct {
		AlgoID string         `json:"algoId"`
		SMsg   string         `json:"sMsg"`
		SCode  test.JSONInt64 `json:"sCode"`
	}
	CancelAlgoOrder struct {
		AlgoID string         `json:"algoId"`
		SMsg   string         `json:"sMsg"`
		SCode  test.JSONInt64 `json:"sCode"`
	}
	AlgoOrder struct {
		InstID       string              `json:"instId"`
		Ccy          string              `json:"ccy"`
		OrdID        string              `json:"ordId"`
		AlgoID       string              `json:"algoId"`
		ClOrdID      string              `json:"clOrdId"`
		TradeID      string              `json:"tradeId"`
		Tag          string              `json:"tag"`
		Category     string              `json:"category"`
		FeeCcy       string              `json:"feeCcy"`
		RebateCcy    string              `json:"rebateCcy"`
		TimeInterval string              `json:"timeInterval"`
		Px           test.JSONFloat64    `json:"px"`
		PxVar        test.JSONFloat64    `json:"pxVar"`
		PxSpread     test.JSONFloat64    `json:"pxSpread"`
		PxLimit      test.JSONFloat64    `json:"pxLimit"`
		Sz           test.JSONInt64      `json:"sz"`
		SzLimit      test.JSONInt64      `json:"szLimit"`
		ActualSz     test.JSONFloat64    `json:"actualSz"`
		ActualPx     test.JSONFloat64    `json:"actualPx"`
		Pnl          test.JSONFloat64    `json:"pnl"`
		AccFillSz    test.JSONInt64      `json:"accFillSz"`
		FillPx       test.JSONFloat64    `json:"fillPx"`
		FillSz       test.JSONInt64      `json:"fillSz"`
		FillTime     test.JSONFloat64    `json:"fillTime"`
		AvgPx        test.JSONFloat64    `json:"avgPx"`
		Lever        test.JSONFloat64    `json:"lever"`
		TpTriggerPx  test.JSONFloat64    `json:"tpTriggerPx"`
		TpOrdPx      test.JSONFloat64    `json:"tpOrdPx"`
		SlTriggerPx  test.JSONFloat64    `json:"slTriggerPx"`
		SlOrdPx      test.JSONFloat64    `json:"slOrdPx"`
		OrdPx        test.JSONFloat64    `json:"ordPx"`
		Fee          test.JSONFloat64    `json:"fee"`
		Rebate       test.JSONFloat64    `json:"rebate"`
		State        test.OrderState     `json:"state"`
		TdMode       test.TradeMode      `json:"tdMode"`
		ActualSide   test.PositionSide   `json:"actualSide"`
		PosSide      test.PositionSide   `json:"posSide"`
		Side         test.OrderSide      `json:"side"`
		OrdType      test.AlgoOrderType  `json:"ordType"`
		InstType     test.InstrumentType `json:"instType"`
		TgtCcy       test.QuantityType   `json:"tgtCcy"`
		CTime        test.JSONTime       `json:"cTime"`
		TriggerTime  test.JSONTime       `json:"triggerTime"`
	}
)
