package controllers

import (
	"github.com/astaxie/beego"
	"my_beego/models"
	"my_beego/dto"
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
	score, err := models.GetAllScore()
	var msg string = "success"
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
