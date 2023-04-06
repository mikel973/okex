package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func init() {
	GetConfig("./okex.toml")
}

func main() {

	app := &cli.App{
		Name:     "okex",
		Usage:    "Utility for okex api v5",
		Commands: []*cli.Command{marketCmd, daemonCmd},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return
}
