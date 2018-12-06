package index

import (
	"github.com/astaxie/beego"
	"wyweb/controllers"
)

func init() {
	beego.AutoRouter(&IndexController{})
}

type IndexController struct {
	controllers.BaseController
}

func (c *IndexController) Index() {
	c.TplName = "index/index.html"
}
