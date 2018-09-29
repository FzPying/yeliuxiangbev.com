package main

import (
	"github.com/fzpying/beego"
	_ "github.com/yeliuxiangbev.com/application/routers"
	_ "github.com/yeliuxiangbev.com/conf"
)

func main() {
	beego.Run()
}
