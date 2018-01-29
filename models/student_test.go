package models

import (
	"testing"
	"fmt"
	//"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"time"
	"github.com/garyburd/redigo/redis"
	"beego/logs"
)

//func init()  {
//orm.RegisterDriver("mysql", orm.DRMySQL)
//// 注册默认数据库
//// 我的mysql的root用户密码为tom，打算把数据表建立在名为test数据库里
//// 备注：此处第一个参数必须设置为“default”（因为我现在只有一个数据库），否则编译报错说：必须有一个注册DB的别名为 default
//orm.RegisterModel(new(Score))
//orm.RegisterModel(new(Student))
//orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/localstudy?charset=utf8")
//// 自动建表
//orm.RunSyncdb("default", false, true)
//}

const(
	REDIS_SETNX ="SETNX"
	REDIS_GET ="GET"
	REDIS_EXPIRE ="EXPIRE"

)

func TestGetStudent(t *testing.T) {
	s:=GetStudent()
	fmt.Printf("----------------%v",s)
}
func TestGetStudentInfo(t *testing.T) {
	s:=GetStudentInfo()
	fmt.Printf("----------------%v",s)
}

func TestFunc(t *testing.T) {
	bm, error := cache.NewCache("redis", `{"conn":"127.0.0.1:6379","key":"collectionName","dbNum":"1","password":""}`)
	if error != nil {
		fmt.Println("redis error:", error)
	}
	bm.Put("test", "aksjdfklasjdfklasjdfkljaskldfjaskldjfaksl", time.Second*100)
	v := bm.Get("test")
	fmt.Println("value:", string(v.([]byte)))
}

func TestRedigo(t *testing.T)  {
	rs,err:=redis.Dial("tcp","127.0.0.1:6379")
	defer rs.Close()
	if err!=nil{
		fmt.Println(err)
	}
	key:="wanggang"
	value:="大笨蛋---"
	n,errSet:=rs.Do(REDIS_SETNX,key,value)
	if errSet!=nil{
		fmt.Printf("jksladjfasdf:%v\n",errSet)
	}
	fmt.Printf("---------------%v\n",n)
	if n==int64(1){
		fmt.Println("ok")
		n,_:=rs.Do(REDIS_EXPIRE,key,10*time.Second)
		logs.Info(n)
		if n==int64(1){
			fmt.Println("sssssssssssuccess")
		}
	}else if n==int64(0){
		fmt.Println("failllllllllllllllll")
	}
	v,err1:=redis.String(rs.Do(REDIS_GET,key))
	if err1!=nil{
		fmt.Println(err1)
	}
	fmt.Println(v)

}
