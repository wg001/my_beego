package controllers

import (
	"github.com/astaxie/beego"
	"my_beego/models"
	"my_beego/dto"
	"Test2/framework/beego/logs"
	"net/http"
)

type User struct {
	beego.Controller
}
/**

adsfa
 */
func (c *User) GetUser() {
	userName := c.Ctx.Input.Param(":sex")
	beego.Debug("获取到的名字是:" + userName)
	score, err := models.GetScoreByUser()
	var msg  = "success"
	var code uint16	=	dto.RIGHT_CODE
	if err != nil {
		msg = err.Error()
		code=dto.ERROR_CODE
	}

	c.Data["json"] =dto.ReponseDTO{Code: code, Message: msg, Data: score}
	c.ServeJSON()
}
func (c *User) GetStudent() {
	userName := c.Ctx.Input.Param(":sex")
	beego.Debug("xueshengf :" + userName)
	score, err := models.GetScoreByUser()
	var msg  = "success"
	var code uint16	=	dto.RIGHT_CODE
	if err != nil {
		msg = err.Error()
		code=dto.ERROR_CODE
	}

	c.Data["json"] =dto.ReponseDTO{Code: code, Message: msg, Data: score}
	c.ServeJSON()
}
func (c *User) GetAllScore() {
	var respnose dto.ReponseDTO
	logs.Info(">>>>>>>>>>>>>>>>>")
	score, err := models.GetAllScore()
	logs.Info("-----------------")
	var msg = "success"
	var code uint16	=	dto.RIGHT_CODE
	if err != nil {
		msg = err.Error()
		code=dto.ERROR_CODE
	}
	respnose.Message = msg
	respnose.Data	=	score
	respnose.Code	=	code

	c.Data["json"] =respnose
	c.ServeJSON()
}
func (c *User) GetUserScoreContent() {
	result, err := models.GetStudentInfo()
	var msg = "success"
	var code uint16	=	dto.RIGHT_CODE
	if err != nil {
		msg = err.Error()
		code=dto.ERROR_CODE
	}

	c.Data["json"] =dto.ReponseDTO{Code:code,Message:msg,Data:result}
	c.ServeJSON()
}
