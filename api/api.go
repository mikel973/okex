package api

import (
	"context"
	"github.com/mikel973/okex"
	"github.com/mikel973/okex/api/rest"
	"github.com/mikel973/okex/api/ws"
)

// Client is the main api wrapper of okex
type Client struct {
	Rest *rest.ClientRest
	Ws   *ws.ClientWs
	ctx  context.Context
}

// NewClient returns a pointer to a fresh Client
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, destination test.Destination) (*Client, error) {
	restURL := test.RestURL
	wsPubURL := test.PublicWsURL
	wsPriURL := test.PrivateWsURL
	switch destination {
	case test.AwsServer:
		restURL = test.AwsRestURL
		wsPubURL = test.AwsPublicWsURL
		wsPriURL = test.AwsPrivateWsURL
	case test.DemoServer:
		restURL = test.DemoRestURL
		wsPubURL = test.DemoPublicWsURL
		wsPriURL = test.DemoPrivateWsURL
	}

	r := rest.NewClient(apiKey, secretKey, passphrase, restURL, destination)
	c := ws.NewClient(ctx, apiKey, secretKey, passphrase, map[bool]test.BaseURL{true: wsPriURL, false: wsPubURL})

	return &Client{r, c, ctx}, nil
}
