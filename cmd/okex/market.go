package main

import (
	"github.com/urfave/cli/v2"
)

var ConfigCmd = &cli.Command{
	Name:  "config",
	Usage: "Manage node config",
	Subcommands: []*cli.Command{
		configDefaultCmd,
		configUpdateCmd,
	},
}

var configDefaultCmd = &cli.Command{
	Name:  "default",
	Usage: "Print default node config",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "no-comment",
			Usage: "don't comment default values",
		},
	},
	Action: func(cctx *cli.Context) error {

		return nil
	},
}

var configUpdateCmd = &cli.Command{
	Name:  "updated",
	Usage: "Print updated node config",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "no-comment",
			Usage: "don't comment default values",
		},
	},
	Action: func(cctx *cli.Context) error {

		return nil
	},
}

var MarketCmd = &cli.Command{
	Name:  "market",
	Usage: "Get market data from okex",
	Subcommands: []*cli.Command{
		subTickerCmd,
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}

var subTickerCmd = &cli.Command{
	Name:      "root",
	Usage:     "Get a CAR's root CID",
	ArgsUsage: "filename",
	Action: func(c *cli.Context) error {
		return nil
	},
}
