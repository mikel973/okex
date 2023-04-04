package main

import (
	"github.com/BurntSushi/toml"
	"github.com/mikel973/okex"
	"github.com/mikel973/okex/logs"
	"os"
)

type ClientConf struct {
	ApiKey string           `toml:"api_key"`
	Secret string           `toml:"secret"`
	Phrase string           `toml:"phrase"`
	Server test.Destination `toml:"server"`
}

var config *ClientConf

func GetConfig(confFile string) *ClientConf {
	if config == nil {
		initConfig(confFile)
	}

	return config
}

func CheckRequiredFields(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"api_key"},
		{"secret"},
		{"phrase"},
		{"server"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			logs.GetLog().Fatal("Required fields ", v)
		}
	}

	return true
}

func initConfig(confFile string) {

	_, err := os.Stat(confFile)
	if os.IsNotExist(err) {
		defaultConf := ClientConf{
			"YOUR-API-KEY",
			"YOUR-SECRET-KEY",
			"YOUR-PASS-PHRASE",
			test.NormalServer,
		}
		f, err := os.Create(confFile)
		if err != nil {

			logs.GetLog().Warn("create default config error:", err)
		}
		defer f.Close()

		if err := toml.NewEncoder(f).Encode(defaultConf); err != nil {
			logs.GetLog().Warn("write default config error:", err)
		}
	}

	if metaData, err := toml.DecodeFile(confFile, &config); err != nil {
		logs.GetLog().Fatal("read config error:", err)

	} else {
		if !CheckRequiredFields(metaData) {
			logs.GetLog().Fatal("Required fields not given")
		}
	}
}
