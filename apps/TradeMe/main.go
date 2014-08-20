// main.go
package main

import (
	"SmartTrade/TradeMe"
	"SmartTrade/utils"
)

const (
	APP_NAME  = "TradeMe "
	VER_INFO  = " 0.1"
	CONF_FILE = "conf/conf.json"
)

func main() {
	utils.LogAppInfo(APP_NAME + VER_INFO)

	si := TradeMe.SShareInfo{}

	si.Init(CONF_FILE)

	si.WaitingExit()
}
