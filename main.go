package main

import (
	"github.com/fzpying/beego"
	_ "github.com/fzpying/yeliuxiangbev.com/application/routers"
	_ "github.com/fzpying/yeliuxiangbev.com/conf"
)

func main() {

	beego.Run()
}
