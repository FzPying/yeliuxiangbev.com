package routers

import (
	"github.com/fzpying/beego"
	"github.com/fzpying/yeliuxiangbev.com/application/controllers/home"
	"github.com/fzpying/yeliuxiangbev.com/application/controllers/manage"
)

func init() {
	beego.Router("/", &home.MainController{})

	beego.Include(&manage.EntryController{})
}
