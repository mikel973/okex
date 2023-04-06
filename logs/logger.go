package logs

import (
	logging "github.com/ipfs/go-log/v2"
)

var log *logging.ZapEventLogger

func GetLog() *logging.ZapEventLogger {
	if log == nil {
		log = logging.Logger("okex")
		logging.SetLogLevel("okex", "INFO")
	}

	return log
}

func SetLogLevel(level string) {
	logging.SetLogLevel("okex", level)
}
