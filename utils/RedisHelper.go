// RedisHelper.go
package utils

import (
	rds "github.com/garyburd/redigo/redis"
	"log"
)

type SRedisMgr struct {
	conn rds.Conn
}

const (
	CST_TCP = "tcp"
)

func (this *SRedisMgr) Init(strRedisServerAddress string) error {

	var err error
	this.conn, err = rds.Dial(CST_TCP, strRedisServerAddress)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (this *SRedisMgr) GetConn() rds.Conn {
	return this.conn
}
