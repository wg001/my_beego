package domain

import (
	"testing"
	"beego/logs"
)

func Test_Setdata(t *testing.T)  {
	logs.Info("start")
	DataTest("wanggang")
	logs.Info("end")
}
