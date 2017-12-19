package domain

import (
	//"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"my_beego/models"
	"encoding/json"
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
func GetDataFromList(key string) {
	defer redisConn.Close()
	models.Init()
	count:=0
	var studentMap []*models.Student
	for {
		studentObj := models.Student{}

		obj, err := redis.Bytes(redisConn.Do("rpop", key))
		if err != nil {
			fmt.Printf("do redis wrong>>>>>%v\n", err)
		}
		if len(obj) == 0 {
			fmt.Printf("------%v\n", obj)
			return
		}
		json.Unmarshal(obj, &studentObj)
		studentMap = append(studentMap, &studentObj)
		if count%10000==0{
			affectRow := models.SaveStudentMutil(studentMap)
			fmt.Println(affectRow)
		}
		fmt.Println(studentObj)

	}

}
