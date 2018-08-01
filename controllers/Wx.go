package controllers

type WxController struct {
	BaseController
}

//首页主体模板
func (c *WxController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


//首页统计信息
func (c *WxController) Statistic() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}