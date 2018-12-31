package user

import (
	"github.com/astaxie/beego"
	"wyweb/controllers"
	"fmt"
	"wyweb/models/user"
	"strconv"
	"time"
)

type UserController struct {
	controllers.BaseController
}

func init() {
	beego.AutoRouter(&UserController{})
}

func (c *UserController) Login() {

	c.TplName = "login/login.html"
}

func (c *UserController) DoLogin() {
	userName := c.GetString("user_name")
	password := c.GetString("password")
	where := fmt.Sprintf(" and name='%v' and password='%v'", userName, password)
	userInfo := user.User{}
	datalist, err := userInfo.GetRow(where)
	ret := make(map[string]interface{})
	if err != nil {
		ret["status"] = 0
		ret["msg"] = err.Error()
		c.Data["json"] = ret
	} else {
		ret["status"] = 1
		ret["msg"] = ""
		c.Data["json"] = ret
		//记录session
		c.SetSession("userName", datalist.Name)
	}
	c.ServeJSON()
}

func (c *UserController) UserList() {
	c.RetOrder()
	c.TplName = "user/list.html"
}

func (c *UserController) AjaxData() {

	sqlRet := c.RetOrder()

	draw := c.GetString("draw")

	retData := make(map[string]interface{})
	userInfo := new(user.User)
	where := c.GetString("search[value]")
	datalist := userInfo.GetUsers(where, sqlRet)

	if datalist.Error != nil {
		fmt.Println(datalist.Error.Error())
	}

	retData["data"] = datalist.Data
	retData["draw"] = draw
	retData["recordsTotal"] = datalist.Total
	retData["recordsFiltered"] = datalist.Total

	c.Data["json"] = retData

	c.ServeJSON()
	return

}

func (c *UserController) Del() {
	id, _ := strconv.Atoi(c.GetString("id"))
	userInfo := new(user.User)

	res := userInfo.Del(id)

	c.Data["json"] = res
	c.ServeJSON()
}

func (c *UserController) Add() {

	status, _ := strconv.Atoi(c.GetString("Status"))
	name := c.GetString("Name")
	pwd := c.GetString("Password")
	role := c.GetString("Role")

	Id := c.GetString("Id")

	userInfo := user.User{
		Name:     name,
		Status:   status,
		Password: pwd,
		Role:     role,
	}

	//处理修改逻辑
	if Id != "" {
		userInfo.Id, _ = strconv.Atoi(c.GetString("Id"))
		c.Data["json"] = userInfo.Update(&userInfo)
		c.ServeJSON()
	}

	userInfo.CreateTime = strconv.FormatInt(time.Now().Unix(), 10)
	c.Data["json"] = userInfo.Add(&userInfo)
	c.ServeJSON()
}

func (c *UserController) Edit() {

	userInfo := new(user.User)

	where := "and id=" + c.GetString("id")

	row, err := userInfo.GetRow(where)

	res := make(map[string]interface{})

	fmt.Println(err)
	if err != nil {
		res["data"] = ""
		res["status"] = 0;
		res["msg"] = "没有记录"
		c.Data["json"] = res
		c.ServeJSON()

	} else {
		res["data"] = row
		res["status"] = 1;
		res["msg"] = "已获取记录"
		c.Data["json"] = res
		c.ServeJSON()
	}
}


