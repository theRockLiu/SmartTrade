// ShareInfo.go
package TradeMe

import (
	"encoding/json"
	"log"
	"net"
	"os"
)

type SShareInfo struct {
	TcpListener net.Listener
}

type SConfiguration struct {
	TcpListenAddrs []string
}

/*
{
    "TcpListenAddrs": ["0.0.0.0:9999"]
}
*/

func (this *SShareInfo) Init(strConfFile string) bool {

	file, err := os.Open(strConfFile)
	if err != nil {
		log.Fatal(err)
		return false
	}

	decoder := json.NewDecoder(file)
	conf := SConfiguration{}
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal(err)
		return false
	}

	////get config info
	//strListenAddr := flag.String("TcpListenAddr", "0.0.0.0:9999", "TradeMe's tcp listen address")
	//flag.Parse()
	//tcpAddr, err := net.ResolveTCPAddr("tcp", *strListenAddr)

	//startting.
	if len(conf.TcpListenAddrs) == 0 {
		log.Fatalf("get conf info failed, no tcp listen address!\n")
		return false
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", conf.TcpListenAddrs[0])
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

	log.Printf("init succeed!\n")

	return true

}
