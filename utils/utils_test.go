// utils_test.go
package utils

import (
	"github.com/garyburd/redigo/redis"
	"testing"
)

func TestInit(t *testing.T) {

	RdsMgr := SRedisMgr{}
	err := RdsMgr.Init("127.0.0.1:6379")
	if nil != err {
		t.Error("init failed\n")
	}

	reply, err := RdsMgr.GetConn().Do("SET", "key", "value")
	if nil != err {
		t.Error("do failed: ", err)
	}
	t.Log(reply)

	s, err1 := redis.String(RdsMgr.GetConn().Do("GET", "key"))
	if nil != err1 {
		t.Error("get key failed : ", s)
	}

	t.Log(s)

	//log.Println(s)
}
