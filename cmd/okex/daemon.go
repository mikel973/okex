package main

import (
	"github.com/mikel973/okex/events"
	"github.com/mikel973/okex/events/public"
	"github.com/mikel973/okex/logs"
	ws_public "github.com/mikel973/okex/requests/ws/public"
	"github.com/urfave/cli/v2"
	"log"
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

	client := getClient()
	logs.GetLog().Info("Starting ... ")
	errChan := make(chan *events.Error)
	subChan := make(chan *events.Subscribe)
	uSubChan := make(chan *events.Unsubscribe)
	logChan := make(chan *events.Login)
	sucChan := make(chan *events.Success)
	client.Ws.SetChannels(errChan, subChan, uSubChan, logChan, sucChan)

	obCh := make(chan *public.OrderBook)
	err := client.Ws.Public.OrderBook(ws_public.OrderBook{
		InstID:  "BTC-USD-SWAP",
		Channel: "books",
	}, obCh)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case <-logChan:
			log.Print("[Authorized]")
		case success := <-sucChan:
			log.Printf("[SUCCESS]\t%+v", success)
		case sub := <-subChan:
			channel, _ := sub.Arg.Get("channel")
			log.Printf("[Subscribed]\t%s", channel)
		case uSub := <-uSubChan:
			channel, _ := uSub.Arg.Get("channel")
			log.Printf("[Unsubscribed]\t%s", channel)
		case err := <-client.Ws.ErrChan:
			log.Printf("[Error]\t%+v", err)
			for _, datum := range err.Data {
				log.Printf("[Error]\t\t%+v", datum)
			}
		case i := <-obCh:
			ch, _ := i.Arg.Get("channel")
			log.Printf("[Event]\t%s", ch)
			for _, p := range i.Books {
				for i := len(p.Asks) - 1; i >= 0; i-- {
					log.Printf("\t\tAsk\t%+v", p.Asks[i])
				}
				for _, bid := range p.Bids {
					log.Printf("\t\tBid\t%+v", bid)
				}
			}
		case b := <-client.Ws.DoneChan:
			log.Printf("[End]:\t%v", b)
			return nil
		}
	}
	return nil
}

func stopDaemon(c *cli.Context) error {
	return nil
}
