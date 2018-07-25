package main

import (
	_ "github.com/Linvas/wxbot/routers"
	_ "github.com/Linvas/wxbot/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
