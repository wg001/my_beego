package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Score struct {
	Id uint16 `json:"id" orm:"pk;column(id)"`
	Name string `json:"name" orm:"column(name)"`
	TypePosition string `json:"type_position" orm:"column(type_position)"`
	CourseId uint8 `json:"course_id" orm:"column(course_id)"`
	Score uint8 `json:"score" orm:"column(score)"`
}
var table string
var ormObj orm.Ormer
func init()  {
	table=(&Score{}).TableName()
}

func GetScoreByUser() (Score,error){
	o := orm.NewOrm()
	o.QueryTable("score")
	score:=Score{}
	//err := o.Read(&score)
	err:=o.Raw("SELECT * from score where id=?",3).QueryRow(&score)
	if err!=nil{
		fmt.Errorf("something wrong....%v",err)
	}
	return score,nil
}

func GetAllScore()([]*Score,error){

	var score []*Score
	ormObj=orm.NewOrm()
	query:=ormObj.QueryTable(table)
	_, _ =query.All(&score)
	return score,nil
}

////写入新的分数
//func AddScore(score *Score)bool{
//
//}


// 自定义表名（系统自动调用）
func (u *Score) TableName() string {
	return "score"
}