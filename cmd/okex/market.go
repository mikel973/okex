package main

import (
	"github.com/mikel973/okex/logs"
	requests "github.com/mikel973/okex/requests/rest/market"
	"github.com/urfave/cli/v2"
)

var marketCmd = &cli.Command{
	Name:  "market",
	Usage: "Market data from okex",
	Subcommands: []*cli.Command{
		subTickerCmd,
	},
}

var subTickerCmd = &cli.Command{
	Name:  "ticker",
	Usage: "Show current ticker from market",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "id",
			Usage: "Instrument ID, e.g. BTC-USDT.",
			Value: "BTC-USDT",
		},
	},
	Action: getTicker,
}

func getTicker(c *cli.Context) error {
	client := getClient()

	instId := c.String("id")
	req := requests.GetTicker{InstId: instId}
	resp, err := client.Rest.Market.GetTicker(req)
	if err != nil {
		logs.GetLog().Fatal(err)
	}
	logs.GetLog().Debugf("Response:%+v", resp)
	for _, tck := range resp.Tickers {
		logs.GetLog().Infof("Ticker: %+v", tck)
	}

	return nil
}
