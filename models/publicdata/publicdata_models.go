package publicdata

import (
	"github.com/mikel973/okex"
)

type (
	Instrument struct {
		InstID    string               `json:"instId"`
		Uly       string               `json:"uly,omitempty"`
		BaseCcy   string               `json:"baseCcy,omitempty"`
		QuoteCcy  string               `json:"quoteCcy,omitempty"`
		SettleCcy string               `json:"settleCcy,omitempty"`
		CtValCcy  string               `json:"ctValCcy,omitempty"`
		CtVal     test.JSONFloat64     `json:"ctVal,omitempty"`
		CtMult    test.JSONFloat64     `json:"ctMult,omitempty"`
		Stk       test.JSONFloat64     `json:"stk,omitempty"`
		TickSz    test.JSONFloat64     `json:"tickSz,omitempty"`
		LotSz     test.JSONFloat64     `json:"lotSz,omitempty"`
		MinSz     test.JSONFloat64     `json:"minSz,omitempty"`
		Lever     test.JSONFloat64     `json:"lever"`
		InstType  test.InstrumentType  `json:"instType"`
		Category  test.FeeCategory     `json:"category,string"`
		OptType   test.OptionType      `json:"optType,omitempty"`
		ListTime  test.JSONTime        `json:"listTime"`
		ExpTime   test.JSONTime        `json:"expTime,omitempty"`
		CtType    test.ContractType    `json:"ctType,omitempty"`
		Alias     test.AliasType       `json:"alias,omitempty"`
		State     test.InstrumentState `json:"state"`
	}
	DeliveryExerciseHistory struct {
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		TS      test.JSONTime                     `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstID string                    `json:"instId"`
		Px     test.JSONFloat64          `json:"px"`
		Type   test.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		InstID   string              `json:"instId"`
		Oi       test.JSONFloat64    `json:"oi"`
		OiCcy    test.JSONFloat64    `json:"oiCcy"`
		InstType test.InstrumentType `json:"instType"`
		TS       test.JSONTime       `json:"ts"`
	}
	FundingRate struct {
		InstID          string              `json:"instId"`
		InstType        test.InstrumentType `json:"instType"`
		FundingRate     test.JSONFloat64    `json:"fundingRate"`
		NextFundingRate test.JSONFloat64    `json:"NextFundingRate"`
		FundingTime     test.JSONTime       `json:"fundingTime"`
		NextFundingTime test.JSONTime       `json:"nextFundingTime"`
	}
	LimitPrice struct {
		InstID   string              `json:"instId"`
		InstType test.InstrumentType `json:"instType"`
		BuyLmt   test.JSONFloat64    `json:"buyLmt"`
		SellLmt  test.JSONFloat64    `json:"sellLmt"`
		TS       test.JSONTime       `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string              `json:"instId"`
		InstType test.InstrumentType `json:"instType"`
		SettlePx test.JSONFloat64    `json:"settlePx"`
		TS       test.JSONTime       `json:"ts"`
	}
	OptionMarketData struct {
		InstID   string              `json:"instId"`
		Uly      string              `json:"uly"`
		InstType test.InstrumentType `json:"instType"`
		Delta    test.JSONFloat64    `json:"delta"`
		Gamma    test.JSONFloat64    `json:"gamma"`
		Vega     test.JSONFloat64    `json:"vega"`
		Theta    test.JSONFloat64    `json:"theta"`
		DeltaBS  test.JSONFloat64    `json:"deltaBS"`
		GammaBS  test.JSONFloat64    `json:"gammaBS"`
		VegaBS   test.JSONFloat64    `json:"vegaBS"`
		ThetaBS  test.JSONFloat64    `json:"thetaBS"`
		Lever    test.JSONFloat64    `json:"lever"`
		MarkVol  test.JSONFloat64    `json:"markVol"`
		BidVol   test.JSONFloat64    `json:"bidVol"`
		AskVol   test.JSONFloat64    `json:"askVol"`
		RealVol  test.JSONFloat64    `json:"realVol"`
		TS       test.JSONTime       `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string           `json:"ccy"`
		Amt          test.JSONFloat64 `json:"amt"`
		DiscountLv   test.JSONInt64   `json:"discountLv"`
		DiscountInfo []*DiscountInfo  `json:"discountInfo"`
	}
	DiscountInfo struct {
		DiscountRate test.JSONInt64 `json:"discountRate"`
		MaxAmt       test.JSONInt64 `json:"maxAmt"`
		MinAmt       test.JSONInt64 `json:"minAmt"`
	}
	SystemTime struct {
		TS test.JSONTime `json:"ts"`
	}
	LiquidationOrder struct {
		InstID    string                    `json:"instId"`
		Uly       string                    `json:"uly,omitempty"`
		InstType  test.InstrumentType       `json:"instType"`
		TotalLoss test.JSONFloat64          `json:"totalLoss"`
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string            `json:"ccy,omitempty"`
		Side    test.OrderSide    `json:"side"`
		OosSide test.PositionSide `json:"posSide"`
		BkPx    test.JSONFloat64  `json:"bkPx"`
		Sz      test.JSONFloat64  `json:"sz"`
		BkLoss  test.JSONFloat64  `json:"bkLoss"`
		TS      test.JSONTime     `json:"ts"`
	}
	MarkPrice struct {
		InstID   string              `json:"instId"`
		InstType test.InstrumentType `json:"instType"`
		MarkPx   test.JSONFloat64    `json:"markPx"`
		TS       test.JSONTime       `json:"ts"`
	}
	PositionTier struct {
		InstID       string              `json:"instId"`
		Uly          string              `json:"uly,omitempty"`
		InstType     test.InstrumentType `json:"instType"`
		Tier         test.JSONInt64      `json:"tier"`
		MinSz        test.JSONFloat64    `json:"minSz"`
		MaxSz        test.JSONFloat64    `json:"maxSz"`
		Mmr          test.JSONFloat64    `json:"mmr"`
		Imr          test.JSONFloat64    `json:"imr"`
		OptMgnFactor test.JSONFloat64    `json:"optMgnFactor,omitempty"`
		QuoteMaxLoan test.JSONFloat64    `json:"quoteMaxLoan,omitempty"`
		BaseMaxLoan  test.JSONFloat64    `json:"baseMaxLoan,omitempty"`
		MaxLever     test.JSONFloat64    `json:"maxLever"`
		TS           test.JSONTime       `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string           `json:"ccy"`
		Rate  test.JSONFloat64 `json:"rate"`
		Quota test.JSONFloat64 `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		Level         string           `json:"level"`
		IrDiscount    test.JSONFloat64 `json:"irDiscount"`
		LoanQuotaCoef int              `json:"loanQuotaCoef,string"`
	}
	State struct {
		Title       string        `json:"title"`
		State       string        `json:"state"`
		Href        string        `json:"href"`
		ServiceType string        `json:"serviceType"`
		System      string        `json:"system"`
		ScheDesc    string        `json:"scheDesc"`
		Begin       test.JSONTime `json:"begin"`
		End         test.JSONTime `json:"end"`
	}
)
