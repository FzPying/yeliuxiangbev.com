package home

import "github.com/yeliuxiangbev.com/application/controllers"

type MainController struct {
	controllers.BaseController
}

func (mc *MainController) Get() {
	mc.Data["Website"] = "beego.me"
	mc.Data["Email"] = "astaxie@gmail.com"
	mc.TplName = "home/index.html"
}
