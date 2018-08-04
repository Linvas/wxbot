package controllers

import (
	"github.com/Linvas/xbot/enums"
)

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Index() {
	c.Data["error"] = c.GetString(":error")
	c.setTpl("home/error.html", "shared/layout_pullbox.html")
}

//首页主体模板
func (c *ErrorController) NotFound() {
	if c.Ctx.Request.Method == "POST" {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", "")
	}
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	c.setTpl()
}

//首页统计信息
func (c *ErrorController) Statistic() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}