package routers

import (
	"github.com/fzpying/beego"
	"github.com/fzpying/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/fzpying/yeliuxiangbev.com/application/controllers/manage:EntryController"] = append(beego.GlobalControllerRouter["github.com/fzpying/yeliuxiangbev.com/application/controllers/manage:EntryController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/manage/entry`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

}
