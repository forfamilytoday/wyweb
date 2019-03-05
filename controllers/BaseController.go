package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type BaseController struct {
	beego.Controller
}

//权限判断的时候使用beego.InsertFilter方法

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
	if length == "" {
		length = "10"
	}
	start := c.GetString("start")
	if start == "" {
		start = "0"
	}
	parma := c.GetString("order[0][column]")
	order := c.GetString("columns[" + parma + "][name]")

	if order == "" {
		order = c.GetString("columns[0][name]")
	}
	asc := c.GetString("order[0][dir]")

	if asc == "" {
		asc = "asc"
	}

	sqlRet = fmt.Sprintf("order by %v %v limit %v,%v", order, asc, start, length)
	return
}

func (c *BaseController) CheckLogin() {

	userName := c.GetSession("userName")
	//将登陆界面排除
	if c.Ctx.Request.RequestURI != "/user/login" && c.Ctx.Request.RequestURI != "/user/dologin" && userName == nil {
		//todo post 方式请求Redirect不起作用
		//c.Redirect("/user/login", 302)
	}
	c.Data["LoginUser"] = userName

}

