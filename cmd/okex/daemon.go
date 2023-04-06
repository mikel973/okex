package main

import (
	"github.com/urfave/cli/v2"
)

var daemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a long-running daemon process",
	Subcommands: []*cli.Command{
		daemonStopCmd,
	},
	Action: startDaemon,
}

var daemonStopCmd = &cli.Command{
	Name:   "stop",
	Usage:  "Stop daemon process",
	Action: stopDaemon,
}

func startDaemon(c *cli.Context) error {
	return nil
}

func stopDaemon(c *cli.Context) error {
	return nil
}
