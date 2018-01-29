package domain

import (
	//"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"my_beego/models"
	"encoding/json"
	"strconv"
	"beego/logs"
	"time"
	"sync"
)

var redisConn redis.Conn

func init() {
	//redisObj, err := cache.NewCache("redis", `{"conn": "127.0.0.1:6379"}`)
	//if err != nil {
	//	fmt.Printf("init err,%v",err)
	//}
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	redisConn = conn
}

//
//
//func Setdata(){
//	timeoutDuration := 1000 * time.Second
//	for i:=0;i<100;i++{
//		if err := redis.Put("xxx:"+strconv.Itoa(i), 1, timeoutDuration); err != nil {
//			fmt.Printf("init err %v", err)
//		}
//	}
//}
//
var l sync.Mutex

func GetDataFromList(key string, redisConn redis.Conn, c chan<- int) {
	logs.Debug("goroutine start...")
	logs.Debug(redisConn)
	count := 0
	var studentMap []models.Student
	for {
		studentObj := models.Student{}
		l.Lock()
		obj, err := redis.Bytes(redisConn.Do("rpop", key))
		if err != nil {
			fmt.Printf("do redis wrong>>>>>%v\n", err)
		}
		l.Unlock()
		if len(obj) == 0 {
			fmt.Printf("------%v\n", obj)
			break
		}
		json.Unmarshal(obj, &studentObj)
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		studentObj.CreateTime = timeNow
		studentObj.UpdateTime = timeNow
		logs.Debug("-------------")
		logs.Debug(studentObj)
		studentMap = append(studentMap, studentObj)
		logs.Debug("========")
		logs.Debug(studentMap)
		if count%20000 == 0 {
			affectRow := models.SaveStudentMutil(studentMap)
			fmt.Println(affectRow)
			studentMap= []models.Student{}
		}
		fmt.Println(studentObj)
	}

	c <- 1
}

func DataTest(key string) {
	models.Init()
	c := make(chan int)
	gocount := 20
	for i := 0; i < gocount; i++ {
		go func(key string, redisConn redis.Conn, c chan<- int, ) {
			GetDataFromList(key, redisConn, c)
		}(key, redisConn, c)
	}
	x := 0
	for i := 0; i < gocount; i++ {
		x = x + <-c
	}
	//defer (*redisConn).Close()
	fmt.Println(">>>>>>>>>>>>>--" + strconv.Itoa(x))
	close(c)
}
