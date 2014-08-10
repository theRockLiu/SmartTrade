// ShareInfo.go
package TradeMe

import (
	"flag"
	"log"
	"net"
)

type SShareInfo struct {
	TcpListener net.Listener
}

func (this *SShareInfo) Init(listenPort uint16, listenIP string) bool {

	//get config info
	strListenAddr := flag.String("TcpListenAddr", "0.0.0.0:9999", "TradeMe's tcp listen address")
	flag.Parse()

	//startting.
	tcpAddr, err := net.ResolveTCPAddr("tcp", *strListenAddr)
	if err != nil {
		log.Fatal(err)
		return false
	}
	tcpListener, err := net.Listen("tcp", tcpAddr.String())
	if err != nil {
		log.Fatal(err)
		return false
	}
	this.TcpListener = tcpListener

	return true

}
