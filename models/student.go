package models

import (
	"time"
	//"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com1/astaxie/beego/logs"
)

type Student struct {
	Id         uint32    `json:"id"`
	Name       string    `json:"name"`
	Birthday   string    `json:"birthday"`
	Sex        uint8     `json:"sex"`
	Email      string    `json:"email"`
	Address    string    `json:"address"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
	Score      []*Score  `json:"score" orm:"reverse(many)"`
}

type StudentInfo struct {
	Student *Student
	Score   []*Score `orm:"reverse(many)"`
}

var studentTable string

func init() {
	studentTable = (&Student{}).TableName()
}

func GetStudent() Student {
	ormObj := orm.NewOrm()
	strudent := Student{}
	table := (&strudent).TableName()
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id>0", table)

	err := ormObj.Raw(sql).QueryRow(&strudent)
	if err != nil {
		fmt.Printf("--------%v", err)
	}
	return strudent
}

/**
这是一种写法，原生sql的
 */
func GetStudentInfo() ([]orm.Params,error) {
	ormObj := orm.NewOrm()
	//strudent := StudentInfo{}
	sql := fmt.Sprintf("SELECT s.id,s.address,s.create_time,s.name,s.email,c.course_id,c.score FROM %s as s left join score as c on c.user_id=s.id WHERE s.id>0", studentTable)
	logs.Info(">>>>>" + sql)
	var maps []orm.Params
	//_,err := ormObj.Raw(sql).QueryRows(&strudent)
	_,err := ormObj.Raw(sql).Values(&maps)
	if err != nil {
		fmt.Printf("vvvvvvvvvvvv,%v", err)
		return nil,err
	}
	logs.Info(maps)
	return maps,nil
}

func (student *Student) TableName() string {
	return "student"
}
