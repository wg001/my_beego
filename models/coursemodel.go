package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Course struct {
	Id         uint   `json:"id"`
	CourseName string `json:"course_name"`
	Status     uint8  `json:"status"`
}

const COURSETBALE = "course" //对应表

//获取课程相关信息
func GetCourseInfoById(courseId uint) (interface{},error) {
	o := orm.NewOrm()
	//o.QueryTable(COURSETBALE)
	sql:="SELECT * FROM "+COURSETBALE+" WHERE id="+strconv.Itoa(int(courseId))
	var maps []orm.Params
	count,err:=o.Raw(sql).Values(&maps)
	if err!=nil{
		return nil,err
	}
	if count==0{
		return nil,nil
	}
	return maps,nil
}
