// main.go
package main

import (
	"SmartTrade/utils"
)

const strMyInfo string = "just test"

func main() {
	//print base information
	utils.LogBaseInfo(strMyInfo)

	//
	StartCSService()

	//
	StartBSService()

	//wait for stopping...
	utils.WaitStopping()
}
