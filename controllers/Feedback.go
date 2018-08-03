package controllers

type FeedbackController struct {
	BaseController
}

//首页主体模板
func (c *FeedbackController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


//首页统计信息
func (c *FeedbackController) Statistic() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}