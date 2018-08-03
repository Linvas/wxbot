package controllers

type MainController struct {
	BaseController
}



//首页主体模板
func (c *MainController) Index() {


	c.setTpl()
}


//首页统计信息
func (c *MainController) Statistic() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}