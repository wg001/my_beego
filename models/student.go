package models

import (
	"time"
	//"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Student struct {
	Id uint32 `json:"id"`
	Name string `json:"name"`
	Birthday string `json:"birthday"`
	Sex uint8	`json:"sex"`
	Email string `json:"email"`
	Address string 	`json:"address"`
	UpdateTime time.Duration `json:"update_time"`
	CreateTime time.Duration `json:"create_time"`
}

type StudentInfo struct {
	Student *Student	`json:"student"`
	Score	[]*Score `json:"score" orm:"reverse(many)`
}


func init()  {
	table=(&Student{}).TableName()
}

func GetStudent() Student {
	ormObj := orm.NewOrm()
	strudent:=Student{}
	sql:=fmt.Sprintf("SELECT * FROM %s WHERE id>0",table)

	err:=ormObj.Raw(sql).QueryRow(&strudent)
	if err!=nil{
		fmt.Printf("--------%v",err)
	}
	return strudent
}

func GetStudentInfo()StudentInfo{
	ormObj := orm.NewOrm()
	strudent:=StudentInfo{}
	sql:=fmt.Sprintf("SELECT s.*,c.* FROM %s as s left join score as c on c.user_id=s.id WHERE s.id>0",table)
	err:=ormObj.Raw(sql).QueryRow(&strudent)
	if err!=nil{
		fmt.Printf("vvvvvvvvvvvv,%v",err)
	}
	return strudent
}


func (student *Student) TableName() string {
	return "student"
}