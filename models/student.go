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
	"go/src/log"
	"sync"
)

type Student struct {
	Id         uint32 `json:"id"`
	Name       string `json:"name"`
	Birthday   string `json:"birthday"`
	Sex        uint8  `json:"sex"`
	Email      string `json:"email"`
	Address    string `json:"address"`
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
func GetStudentInfo() ([]orm.Params, error) {
	ormObj := orm.NewOrm()
	//strudent := StudentInfo{}
	sql := fmt.Sprintf("SELECT s.id,s.address,s.create_time,s.name,s.email,c.course_id,c.score,cc.course_name FROM %s as s left join score as c on c.user_id=s.id LEFT JOIN course as cc ON c.course_id=cc.id  WHERE s.id>0", studentTable)
	logs.Info(">>>>>" + sql)
	var maps []orm.Params
	//_,err := ormObj.Raw(sql).QueryRows(&strudent)
	_, err := ormObj.Raw(sql).Values(&maps)
	if err != nil {
		fmt.Printf("vvvvvvvvvvvv,%v", err)
		return nil, err
	}
	bytes, _ := json.Marshal(maps)
	logs.Info(string(bytes))
	return maps, nil
}

func (student *Student) TableName() string {
	return "student"
}

func SaveStudentByObj(student *Student) int64 {
	if student == nil {
		fmt.Println("ssssstudent wrong")
	}
	ormObj := orm.NewOrm()
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	if student.UpdateTime==""{
		student.UpdateTime=timeNow
	}
	if student.CreateTime==""{
		student.CreateTime=timeNow
	}
	affectRow, err := ormObj.Insert(student)
	if err != nil {
		fmt.Printf("Insert Stundet wrong---%v",err)
	}
	return affectRow
}

func SaveStudentMutil(studentMap []Student)int64{
	ormObj := orm.NewOrm()
	affectRow, err := ormObj.InsertMulti(len(studentMap), studentMap)
	if err!=nil{
		logs.Debug("*****************")
		logs.Debug(studentMap)
		logs.Debug("*****************")
		fmt.Printf("something wrong %v\n",err)
	}
	fmt.Printf("affect rows :%d",affectRow)
	return affectRow
}

func SaveStudent() {
	ormObj := orm.NewOrm()
	var studentMap []Student
	for i := 0; i < 2000000; i++ {
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		studentObj := Student{}
		studentObj.Address = "jk" + strconv.Itoa(i)
		studentObj.Birthday = timeNow
		studentObj.Email = strconv.Itoa(i+1) + "@qq.com"
		studentObj.Name = "4542" + strconv.Itoa(i)
		studentObj.Sex = 1
		studentObj.CreateTime = timeNow
		studentObj.UpdateTime = timeNow
		studentMap = append(studentMap, studentObj)
		//affectrows, err := ormObj.Insert(&studentObj)
		if len(studentMap)%1000 == 0 {
			affectRow, err := ormObj.InsertMulti(200, studentMap)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("影响行数：" + strconv.Itoa(int(affectRow)))
			studentMap = []Student{}
		}
		//if err != nil {
		//	fmt.Println(err)
		//}
		//logs.Info("影响行数" + strconv.Itoa(int(affectrows)))
	}
	affectRow, err := ormObj.InsertMulti(200, studentMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("影响行数：" + strconv.Itoa(int(affectRow)))
	fmt.Println(studentMap)
	fmt.Println("over")
}

/**
原生sql的
*/
func GetStudentInfoLimit(offset uint) ([]orm.Params, error) {
	ormObj := orm.NewOrm()
	length := 100000
	sql := fmt.Sprintf("SELECT s.id,s.address,s.create_time,s.name,s.email FROM %s as s WHERE s.id>0 and s.id<=500000 limit %d,%d", studentTable, int(offset)*length, length)
	//logs.Info(">>>>>" + sql)
	var maps []orm.Params
	//_,err := ormObj.Raw(sql).QueryRows(&strudent)
	_, err := ormObj.Raw(sql).Values(&maps)
	if err != nil {
		fmt.Printf("vvvvvvvvvvvv,%v", err)
		return nil, err
	}
	bytes, _ := json.Marshal(maps)
	logs.Info(string(bytes))
	return maps, nil
}

type ExcelObj struct {
	xlsx  *excelize.File
	wlock *sync.RWMutex
}

var m *sync.RWMutex

func GetData2Excel() {
	var result []orm.Params
	limit := uint(0)
	logs.Debug("Start>>>>>>>>>>>>>>>>>>")
	for {
		param, err := GetStudentInfoLimit(limit)
		if err != nil {
			fmt.Printf(">>>>>>>>%v \n", err)
		}
		if len(param) == 0 {
			break
		}
		limit = limit + 1
		slice := make([]orm.Params, len(result)+len(param))
		copy(slice, result)
		copy(slice[len(result):], param)
		param = nil
		result = slice
		slice = nil
	}
	m = new(sync.RWMutex)

	fmt.Println(">>>>>>>>>>>>>>")
	fmt.Println(len(result))
	xlsx := excelize.NewFile()
	sheetPath := fmt.Sprintf("./Workbook%s.xlsx", "1")
	xlsx.SaveAs(sheetPath)
	logs.Info("\n get resource")
	c := make(chan int)
	//d := make(chan int)
	xlsx, err := excelize.OpenFile(sheetPath)
	if err != nil {
		logs.Info("<<<<<<<<<<<<<%v\n", err)
	}
	gocount := 1
	for i := 0; i < gocount; i++ {
		go func(result2 []orm.Params, xlsx *excelize.File, c chan<- int, start int) {
			Save1(result2, xlsx, c, start)
		}(result[i*len(result)/gocount:(i+1)*len(result)/gocount], xlsx, c, i*len(result)/gocount)
		//Save2(result[i*len(result)/gocount:(i+1)*len(result)/gocount],xlsx,i*len(result)/gocount)
	}
	var x int
	for j := 0; j < gocount; j++ {
		x = x + <-c
	}
	if x >= gocount {
		fmt.Println("oooooooooooooooook")
		logs.Debug("START SAVE")
		xlsx.Save()
		logs.Debug("SAVE FINISHED")
	}
	//Save1(result[:len(result)/2],xlsx)
	logs.Debug("<<<<<<<<<<<END")
}

func saveExcel(c <-chan *excelize.File, count int) {
	for i := 0; i < count; i++ {
		v := <-c
		log.Printf("cccccccccccccc%v\n", v)
		v.Save()
	}
}

func Save1(result []orm.Params, xlsx *excelize.File, c chan<- int, start int) {
	sheetName := "Sheet1"
	logs.Debug("TAG_%d PREPARE DATA", start)
	m.Lock()
	for key, value := range result {
		if key%10000 == 0 {
			logs.Debug("进行到了》》" + strconv.Itoa(key))
		}
		xlsx.SetCellValue(sheetName, "A"+strconv.Itoa(key+start+1), value["address"])
		xlsx.SetCellValue(sheetName, "B"+strconv.Itoa(key+start+1), value["create_time"])
		xlsx.SetCellValue(sheetName, "C"+strconv.Itoa(key+start+1), value["name"])
		xlsx.SetCellValue(sheetName, "D"+strconv.Itoa(key+start+1), value["email"])
		xlsx.SetCellValue(sheetName, "E"+strconv.Itoa(key+start+1), value["id"])
	}
	m.Unlock()
	logs.Debug("TAG_%d PREPARE FINISHED", start)
	c <- 1
}

func Save2(result []orm.Params, xlsx *excelize.File, start int) {
	sheetName := "Sheet1"
	logs.Debug("tag_%d PREPARE", start)
	for key, value := range result {
		xlsx.SetCellValue(sheetName, "A"+strconv.Itoa(key+start+1), value["address"])
		xlsx.SetCellValue(sheetName, "B"+strconv.Itoa(key+start+1), value["create_time"])
		xlsx.SetCellValue(sheetName, "C"+strconv.Itoa(key+start+1), value["name"])
		xlsx.SetCellValue(sheetName, "D"+strconv.Itoa(key+start+1), value["email"])
		xlsx.SetCellValue(sheetName, "E"+strconv.Itoa(key+start+1), value["id"])
	}
	//m.Lock()
	//logs.Debug("tag_%d START",start)
	//xlsx.Save()
	//logs.Debug("tag_%d FINISHED",start)
	//m.Unlock()
}
