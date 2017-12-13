package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	userName:=c.Ctx.Input.Param(":username")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	beego.Debug("获取到的名字是:"+userName)
	c.TplName = "index.tpl"
}
