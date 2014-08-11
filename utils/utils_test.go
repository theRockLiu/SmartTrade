// utils_test.go
package utils

import (
	"log"
	"testing"
)

func TestInit(t *testing.T) {

	RdsMgr := SRedisMgr{}
	err := RdsMgr.Init("127.0.0.1:6379")
	if nil != err {
		t.Error("init failed\n")
	}

	reply, err := RdsMgr.GetConn().Do("APPEND", "key", "value")
	if nil != err {
		t.Error("do failed: ", err)
	}

	log.Println(reply)
}
