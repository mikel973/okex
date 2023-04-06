package main

import (
	"context"
	"github.com/mikel973/okex/api"
	"github.com/mikel973/okex/logs"
)

func getClient() *api.Client {
	conf := GetConfig("")
	ctx := context.Background()
	client, err := api.NewClient(ctx, conf.ApiKey, conf.Secret, conf.Phrase, conf.Server)
	if err != nil {
		logs.GetLog().Fatal(err)
	}
	return client
}
