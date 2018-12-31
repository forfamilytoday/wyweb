package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	c.CheckLogin()
}

type BaseSearch struct {
	start int
	limit int
	order string
}

func (c *BaseController) RetOrder() (sqlRet string) {

	length := c.GetString("length")
	start := c.GetString("start")
	parma := c.GetString("order[0][column]")
	order := c.GetString("columns[" + parma + "][name]")
	asc := c.GetString("order[0][dir]")

	sqlRet = fmt.Sprintf("order by %v %v limit %v,%v", order, asc, start, length)
	return
}

func (c *BaseController) CheckLogin() {

	userName := c.GetSession("userName")
	//将登陆界面排除
	if c.Ctx.Request.RequestURI != "/user/login" && c.Ctx.Request.RequestURI != "/user/dologin" && userName == nil {
		c.Redirect("/user/login", 302)
	}
	c.Data["LoginUser"] = userName

}
