package manage

import (
	"github.com/fzpying/yeliuxiangbev.com/application/controllers"
	"github.com/fzpying/yeliuxiangbev.com/application/models"
)

type EntryController struct {
	controllers.BaseController
}

// @router /manage/entry [get]
func (eC *EntryController) Get() {
	Users := models.Users{}
	Users.Id = 1
	Users.Name = "FzPying"
	models.AddUser(&Users)

	eC.TplName = "manage/entry.html"
}
