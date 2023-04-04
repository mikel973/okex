package public

import "github.com/mikel973/okex"

type (
	GetInstruments struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		InstType test.InstrumentType `json:"instType"`
	}
	GetDeliveryExerciseHistory struct {
		Uly      string              `json:"uly"`
		After    int64               `json:"after,omitempty,string"`
		Before   int64               `json:"before,omitempty,string"`
		Limit    int64               `json:"limit,omitempty,string"`
		InstType test.InstrumentType `json:"instType"`
	}
	GetOpenInterest struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		InstType test.InstrumentType `json:"instType"`
	}
	GetFundingRate struct {
		InstID string `json:"instId"`
	}
	GetLimitPrice struct {
		InstID string `json:"instId"`
	}
	GetOptionMarketData struct {
		Uly     string `json:"uly"`
		ExpTime string `json:"expTime,omitempty"`
	}
	GetEstimatedDeliveryExercisePrice struct {
		Uly     string `json:"uly"`
		ExpTime string `json:"expTime,omitempty"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Uly        string  `json:"uly"`
		Ccy        string  `json:"ccy,omitempty"`
		DiscountLv float64 `json:"discountLv,string"`
	}
	GetLiquidationOrders struct {
		InstID   string              `json:"instId,omitempty"`
		Ccy      string              `json:"ccy,omitempty"`
		Uly      string              `json:"uly,omitempty"`
		After    int64               `json:"after,omitempty,string"`
		Before   int64               `json:"before,omitempty,string"`
		Limit    int64               `json:"limit,omitempty,string"`
		InstType test.InstrumentType `json:"instType"`
		MgnMode  test.MarginMode     `json:"mgnMode,omitempty"`
		Alias    test.AliasType      `json:"alias,omitempty"`
		State    test.OrderState     `json:"state,omitempty"`
	}
	GetMarkPrice struct {
		InstID   string              `json:"instId,omitempty"`
		Uly      string              `json:"uly,omitempty"`
		InstType test.InstrumentType `json:"instType"`
	}
	GetPositionTiers struct {
		InstID   string              `json:"instId,omitempty"`
		Uly      string              `json:"uly,omitempty"`
		InstType test.InstrumentType `json:"instType"`
		TdMode   test.TradeMode      `json:"tdMode"`
		Tier     test.JSONInt64      `json:"tier,omitempty"`
	}
	GetUnderlying struct {
		InstType test.InstrumentType `json:"instType"`
	}
	Status struct {
		State string `json:"state,omitempty"`
	}
)
