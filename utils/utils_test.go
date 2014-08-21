// utils_test.go
package utils

import (
	"github.com/garyburd/redigo/redis"
	"github.com/ssdb/gossdb/ssdb"
	"testing"
)

func TestInit(t *testing.T) {
	t.Skip("skip redis")

	RdsMgr := SRedisMgr{}
	err := RdsMgr.Init("127.0.0.1:6379")
	if nil != err {
		t.Error("init failed\n")
		return
	}

	reply, err := RdsMgr.GetConn().Do("SET", "key", "value")
	if nil != err {
		t.Error("do failed: ", err)
		return
	}
	t.Log(reply)

	s, err1 := redis.String(RdsMgr.GetConn().Do("GET", "key"))
	if nil != err1 {
		t.Error("get key failed : ", s)
		return
	}

	t.Log(s)

	//log.Println(s)
}

func TestSsdb(t *testing.T) {
	ip := "127.0.0.1"
	port := 8888
	db, err := ssdb.Connect(ip, port)
	if err != nil {
		t.Error("connect db failed: ", err)
		return
	}
	defer db.Close()
	var val interface{}

	keys := []string{}
	keys = append(keys, "c")
	keys = append(keys, "d")
	val, err = db.Do("multi_get", "a", "b", keys)
	t.Logf("%s\n", val)

	db.Set("a", "xxx")
	val, err = db.Get("a")
	t.Logf("%s\n", val)
	db.Del("a")
	val, err = db.Get("a")
	t.Logf("%s\n", val)

	t.Logf("----\n")

	db.Do("zset", "z", "a", 3)
	db.Do("multi_zset", "z", "b", -2, "c", 5, "d", 3)
	resp, err := db.Do("zrscan", "z", "", "", "", 10)
	if err != nil {
		t.Error("zrscan failed : ", err)
		return
	}

	if len(resp)%2 != 1 {
		t.Error("bad response")
		return
	}

	t.Logf("Status: %s\n", resp[0])
	for i := 1; i < len(resp); i += 2 {
		t.Logf("  %s : %3s\n", resp[i], resp[i+1])
	}
	return
}
