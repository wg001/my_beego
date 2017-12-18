package domain

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"fmt"
	"time"
	"strconv"
)

var redis cache.Cache

func init()  {
	redisObj, err := cache.NewCache("redis", `{"conn": "127.0.0.1:6379"}`)
	if err != nil {
		fmt.Printf("init err,%v",err)
	}
	redis = redisObj
}


func Setdata(){
	timeoutDuration := 1000 * time.Second
	for i:=0;i<100;i++{
		if err := redis.Put("xxx:"+strconv.Itoa(i), 1, timeoutDuration); err != nil {
			fmt.Printf("init err %v", err)
		}
	}
}