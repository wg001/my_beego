package models

import (
	"testing"
	"fmt"
)

func TestGetCourseInfoById(t *testing.T) {
	course,_:=GetCourseInfoById(1)
	fmt.Printf("%v",course)
}