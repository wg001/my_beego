package models

import (
	"testing"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 注册默认数据库
	// 我的mysql的root用户密码为tom，打算把数据表建立在名为test数据库里
	// 备注：此处第一个参数必须设置为“default”（因为我现在只有一个数据库），否则编译报错说：必须有一个注册DB的别名为 default
	orm.RegisterModel(new(Score))
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/localstudy?charset=utf8")
	// 自动建表
	orm.RunSyncdb("default", false, true)
}

func Test_GetScoreByUser(t *testing.T) {
	score,_:=GetScoreByUser()
	fmt.Print(score)
}

func TestGetAllScore(t *testing.T) {
	data,_:=GetAllScore()
	fmt.Println(data)
}