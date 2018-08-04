package main

import (
	_ "github.com/Linvas/xbot/routers"
	_ "github.com/Linvas/xbot/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
