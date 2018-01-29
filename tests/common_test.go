package test

import (
	"testing"
	"fmt"
	"beego/logs"
)

func Test_All(t *testing.T){
	logs.Debug("skadjfkafd")
	c:=make(chan int,4)
	c<-4
	c<-3
	c<-1
	close(c)
	for value:=range c {
		fmt.Println(value)
	}
	fmt.Println(len(c))
	fmt.Println("-----")
}