package main

import (
	"github.com/urfave/cli/v2"
)

var utilsCmd = &cli.Command{
	Name:  "util",
	Usage: "Utility for okex",
	Subcommands: []*cli.Command{
		utilsGenCmd,
	},
}

var utilsGenCmd = &cli.Command{
	Name:   "gen",
	Usage:  "gen ",
	Action: UtilsGen,
}

func UtilsGen(c *cli.Context) error {

	//
	//errChan := make(chan *events.Error)
	//subChan := make(chan *events.Subscribe)
	//uSubChan := make(chan *events.Unsubscribe)
	//logChan := make(chan *events.Login)
	//sucChan := make(chan *events.Success)
	//client.Ws.SetChannels(errChan, subChan, uSubChan, logChan, sucChan)
	//
	//obCh := make(chan *public.OrderBook)
	//err = client.Ws.Public.OrderBook(ws_public_requests.OrderBook{
	//	InstID:  "BTC-USD-SWAP",
	//	Channel: "books",
	//}, obCh)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//for {
	//	select {
	//	case <-logChan:
	//		log.Print("[Authorized]")
	//	case success := <-sucChan:
	//		log.Printf("[SUCCESS]\t%+v", success)
	//	case sub := <-subChan:
	//		channel, _ := sub.Arg.Get("channel")
	//		log.Printf("[Subscribed]\t%s", channel)
	//	case uSub := <-uSubChan:
	//		channel, _ := uSub.Arg.Get("channel")
	//		log.Printf("[Unsubscribed]\t%s", channel)
	//	case err := <-client.Ws.ErrChan:
	//		log.Printf("[Error]\t%+v", err)
	//		for _, datum := range err.Data {
	//			log.Printf("[Error]\t\t%+v", datum)
	//		}
	//	case i := <-obCh:
	//		ch, _ := i.Arg.Get("channel")
	//		log.Printf("[Event]\t%s", ch)
	//		for _, p := range i.Books {
	//			for i := len(p.Asks) - 1; i >= 0; i-- {
	//				log.Printf("\t\tAsk\t%+v", p.Asks[i])
	//			}
	//			for _, bid := range p.Bids {
	//				log.Printf("\t\tBid\t%+v", bid)
	//			}
	//		}
	//	case b := <-client.Ws.DoneChan:
	//		log.Printf("[End]:\t%v", b)
	//		return
	//	}
	//}
	return nil
}
