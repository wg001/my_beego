package routers

import (
	"my_beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/:username", &controllers.MainController{})
    beego.Router("/:username", &controllers.MainController{})
    beego.Router("/wg/:sex",&controllers.User{},"post:GetUser")
    beego.Router("/wg/all",&controllers.User{},"post:GetAllScore")
    beego.Router("/wg/get_all",&controllers.User{},"post:GetUserScoreContent")
}
