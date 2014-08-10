// ShareInfo.go
package TradeMe

type SShareInfo struct {
	ui16ListenPort uint16
	strListenIP    string
}

func (this *SShareInfo) Init(listenPort uint16, listenIP string) bool {

	return true

}
