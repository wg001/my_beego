package domain

import (
	"fmt"
)

type StudentObj struct {
	Name string
	Age  int
}

func (student *StudentObj) GetStudentInfo(id string) {
	fmt.Print("jaksdf")
}

