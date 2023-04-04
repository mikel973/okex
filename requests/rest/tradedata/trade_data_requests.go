package tradedata

import "github.com/mikel973/okex"

type (
	GetTakerVolume struct {
		Ccy      string              `json:"ccy"`
		Begin    int64               `json:"before,omitempty,string"`
		End      int64               `json:"limit,omitempty,string"`
		InstType test.InstrumentType `json:"instType"`
		Period   test.BarSize        `json:"period,string,omitempty"`
	}
	GetRatio struct {
		Ccy    string       `json:"ccy"`
		Begin  int64        `json:"before,omitempty,string"`
		End    int64        `json:"limit,omitempty,string"`
		Period test.BarSize `json:"period,string,omitempty"`
	}
	GetOpenInterestAndVolumeStrike struct {
		Ccy     string       `json:"ccy"`
		ExpTime string       `json:"expTime"`
		Period  test.BarSize `json:"period,string,omitempty"`
	}
)
