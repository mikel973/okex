package account

import "github.com/mikel973/okex"

type (
	GetBalance struct {
		Ccy []string `json:"ccy,omitempty"`
	}
	GetPositions struct {
		InstID   []string            `json:"instId,omitempty"`
		PosID    []string            `json:"posId,omitempty"`
		InstType test.InstrumentType `json:"instType,omitempty"`
	}
	GetAccountAndPositionRisk struct {
		InstType test.InstrumentType `json:"instType,omitempty"`
	}
	GetBills struct {
		Ccy      string              `json:"ccy,omitempty"`
		After    int64               `json:"after,omitempty,string"`
		Before   int64               `json:"before,omitempty,string"`
		Limit    int64               `json:"limit,omitempty,string"`
		InstType test.InstrumentType `json:"instType,omitempty"`
		MgnMode  test.MarginMode     `json:"mgnMode,omitempty"`
		CtType   test.ContractType   `json:"ctType,omitempty"`
		Type     test.BillType       `json:"type,omitempty,string"`
		SubType  test.BillSubType    `json:"subType,omitempty,string"`
	}
	SetPositionMode struct {
		PositionMode test.PositionType `json:"positionMode"`
	}
	SetLeverage struct {
		Lever   int64             `json:"lever,string"`
		InstID  string            `json:"instId,omitempty"`
		Ccy     string            `json:"ccy,omitempty"`
		MgnMode test.MarginMode   `json:"mgnMode"`
		PosSide test.PositionSide `json:"posSide,omitempty"`
	}
	GetMaxBuySellAmount struct {
		Ccy    string         `json:"ccy,omitempty"`
		Px     float64        `json:"px,string,omitempty"`
		InstID []string       `json:"instId"`
		TdMode test.TradeMode `json:"tdMode"`
	}
	GetMaxAvailableTradeAmount struct {
		Ccy        string         `json:"ccy,omitempty"`
		InstID     string         `json:"instId"`
		ReduceOnly bool           `json:"reduceOnly,omitempty"`
		TdMode     test.TradeMode `json:"tdMode"`
	}
	IncreaseDecreaseMargin struct {
		InstID     string            `json:"instId"`
		Amt        float64           `json:"amt,string"`
		PosSide    test.PositionSide `json:"posSide"`
		ActionType test.CountAction  `json:"actionType"`
	}
	GetLeverage struct {
		InstID  []string        `json:"instId"`
		MgnMode test.MarginMode `json:"mgnMode"`
	}
	GetMaxLoan struct {
		InstID  string          `json:"instId"`
		MgnCcy  string          `json:"mgnCcy,omitempty"`
		MgnMode test.MarginMode `json:"mgnMode"`
	}
	GetFeeRates struct {
		InstID   string              `json:"instId,omitempty"`
		Uly      string              `json:"uly,omitempty"`
		Category test.FeeCategory    `json:"category,omitempty,string"`
		InstType test.InstrumentType `json:"instType"`
	}
	GetInterestAccrued struct {
		InstID  string          `json:"instId,omitempty"`
		Ccy     string          `json:"ccy,omitempty"`
		After   int64           `json:"after,omitempty,string"`
		Before  int64           `json:"before,omitempty,string"`
		Limit   int64           `json:"limit,omitempty,string"`
		MgnMode test.MarginMode `json:"mgnMode,omitempty"`
	}
	SetGreeks struct {
		GreeksType test.GreekType `json:"greeksType"`
	}
)
