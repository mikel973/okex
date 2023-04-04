package subaccount

import (
	"github.com/mikel973/okex"
)

type (
	SubAccount struct {
		SubAcct string        `json:"subAcct,omitempty"`
		Label   string        `json:"label,omitempty"`
		Mobile  string        `json:"mobile,omitempty"`
		GAuth   bool          `json:"gAuth"`
		Enable  bool          `json:"enable"`
		TS      test.JSONTime `json:"ts"`
	}
	APIKey struct {
		SubAcct    string        `json:"subAcct,omitempty"`
		Label      string        `json:"label,omitempty"`
		APIKey     string        `json:"apiKey,omitempty"`
		SecretKey  string        `json:"secretKey,omitempty"`
		Passphrase string        `json:"Passphrase,omitempty"`
		Perm       string        `json:"perm,omitempty"`
		IP         string        `json:"ip,omitempty"`
		TS         test.JSONTime `json:"ts,omitempty"`
	}
	HistoryTransfer struct {
		SubAcct string         `json:"subAcct,omitempty"`
		Ccy     string         `json:"ccy,omitempty"`
		BillID  test.JSONInt64 `json:"billId,omitempty"`
		Type    test.BillType  `json:"type,omitempty,string"`
		TS      test.JSONTime  `json:"ts,omitempty"`
	}
	Transfer struct {
		TransID test.JSONInt64 `json:"transId"`
	}
)
