package models

import (
	"time"
	//"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com1/astaxie/beego/logs"
	"encoding/json"
	"strconv"
	"github.com/360EntSecGroup-Skylar/excelize"
	//"github.com/tealeg/xlsx"
	"go/src/log"
)

type Student struct {
	Id         uint32    `json:"id"`
	Name       string    `json:"name"`
	Birthday   string    `json:"birthday"`
	Sex        uint8     `json:"sex"`
	Email      string    `json:"email"`
	Address    string    `json:"address"`
	UpdateTime string `json:"update_time"`
	CreateTime string `json:"create_time"`
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
	sql := fmt.Sprintf("SELECT s.id,s.address,s.create_time,s.name,s.email,c.course_id,c.score,cc.course_name FROM %s as s left join score as c on c.user_id=s.id LEFT JOIN course as cc ON c.course_id=cc.id  WHERE s.id>0", studentTable)
	logs.Info(">>>>>" + sql)
	var maps []orm.Params
	//_,err := ormObj.Raw(sql).QueryRows(&strudent)
	_,err := ormObj.Raw(sql).Values(&maps)
	if err != nil {
		fmt.Printf("vvvvvvvvvvvv,%v", err)
		return nil,err
	}
	bytes,_:=json.Marshal(maps)
	logs.Info(string(bytes))
	return maps,nil
}

func (student *Student) TableName() string {
	return "student"
}

func SaveStudent(){
	ormObj:=orm.NewOrm()
	var studentMap []Student
	for i := 0; i<2000000; i++ {
		timeNow:=time.Now().Format("2006-01-02 15:04:05")
		studentObj := Student{}
		studentObj.Address = "jk" + strconv.Itoa(i)
		studentObj.Birthday = timeNow
		studentObj.Email = strconv.Itoa(i+1) + "@qq.com"
		studentObj.Name = "4542" + strconv.Itoa(i)
		studentObj.Sex = 1
		studentObj.CreateTime=timeNow
		studentObj.UpdateTime=timeNow
		studentMap=append(studentMap, studentObj)
		//affectrows, err := ormObj.Insert(&studentObj)
		if len(studentMap)%1000==0{
			affectRow,err:=ormObj.InsertMulti(200,studentMap)
			if err!=nil{
				fmt.Println(err)
			}
			fmt.Println("影响行数："+strconv.Itoa(int(affectRow)))
			studentMap = []Student{}
		}
		//if err != nil {
		//	fmt.Println(err)
		//}
		//logs.Info("影响行数" + strconv.Itoa(int(affectrows)))
	}
	affectRow,err:=ormObj.InsertMulti(200,studentMap)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("影响行数："+strconv.Itoa(int(affectRow)))
	fmt.Println(studentMap)
	fmt.Println("over")
}
/**
这是一种写法，原生sql的
*/
func GetStudentInfoLimit(offset uint) ([]orm.Params,error) {
	ormObj := orm.NewOrm()
	//strudent := StudentInfo{}
	//param:=[]string{studentTable,strconv.Itoa(int(offset)),}
	length:=100000
	//offStart:=uint()
	sql := fmt.Sprintf("SELECT s.id,s.address,s.create_time,s.name,s.email FROM %s as s WHERE s.id>0 and s.id<=100 limit %d,%d",studentTable,int(offset)*length,length)
	logs.Info(">>>>>" + sql)
	var maps []orm.Params
	//_,err := ormObj.Raw(sql).QueryRows(&strudent)
	_,err := ormObj.Raw(sql).Values(&maps)
	if err != nil {
		fmt.Printf("vvvvvvvvvvvv,%v", err)
		return nil,err
	}
	bytes,_:=json.Marshal(maps)
	logs.Info(string(bytes))
	return maps,nil
}
func GetData2Excel(){
	var result []orm.Params
	limit :=uint(0)
	for {
		param,err:=GetStudentInfoLimit(limit)
		if err!=nil{
			fmt.Printf(">>>>>>>>%v \n",err)
		}
		if len(param)==0{
			break
		}
		limit = limit+1
		slice := make([]orm.Params, len(result)+len(param))
		copy(slice,result)
		copy(slice[len(result):],param)
		param=nil
		result = slice
		slice=nil
	}

	fmt.Println(">>>>>>>>>>>>>>")
	fmt.Println(len(result))
	fmt.Print(result)
	xlsx:=excelize.NewFile()
	xlsx.SaveAs("./Workbook1.xlsx")

	//c := make(chan int)
	//d := make(chan int)
	sheetPath:=fmt.Sprintf("./Workbook%s.xlsx","1")
	xlsx,err := excelize.OpenFile(sheetPath)
	if err!=nil{
		logs.Info("<<<<<<<<<<<<<%v\n",err)
	}
	excelSource:=make(chan *excelize.File,2)
	go func(result2 []orm.Params) {
		InsertExcel(result2,excelSource,len(result)/2,"2",xlsx)
	}(result[len(result)/2:])


	go func(result1 []orm.Params) {
		InsertExcel(result1,excelSource,0,"1",xlsx)
	}(result[:len(result)/2])
	saveExcel(excelSource,2)
	fmt.Println(">>>>>>>>>>ok")

}

func saveExcel(c <-chan *excelize.File,count int){
	for i:=0;i<count;i++{
		v:=<-c
		log.Printf("cccccccccccccc%v\n",v)
		v.Save()
	}
}

func InsertExcel(result []orm.Params, c chan<- *excelize.File,start int,tag string,xlsx *excelize.File)  {

	logs.Info("\n get resource")
	//if err!=nil{
	//	fmt.Printf("This is tag:%s--%v",tag,err)
	//	xlsx = excelize.NewFile()
	//	xlsx.SaveAs(sheetPath)
	//	xlsx,_=excelize.OpenFile(sheetPath)
	//}
	//xlsx := excelize.NewFile()
	sheetName:="Sheet1"
	for key,value:=range result{
		xlsx.SetCellValue(sheetName,"A"+strconv.Itoa(key+start+1),value["address"])
		xlsx.SetCellValue(sheetName,"B"+strconv.Itoa(key+start+1),value["create_time"])
		xlsx.SetCellValue(sheetName,"C"+strconv.Itoa(key+start+1),value["name"])
		xlsx.SetCellValue(sheetName,"D"+strconv.Itoa(key+start+1),value["email"])
		xlsx.SetCellValue(sheetName,"E"+strconv.Itoa(key+start+1),value["id"])
	}
	//err = xlsx.Save()
	c<-xlsx
	//if err != nil {
	//	logs.Info("------------------------%s",tag)
	//	fmt.Println(err)
	//}
	logs.Info("\n in goroutine")

}
