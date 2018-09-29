package routers

import (
	"github.com/fzpying/beego"
	"github.com/yeliuxiangbev.com/application/controllers/home"
	"github.com/yeliuxiangbev.com/application/controllers/manage"
)

func init() {
	beego.Router("/", &home.MainController{})

	beego.Include(&manage.EntryController{})
}
