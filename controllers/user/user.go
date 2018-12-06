package user

import (
	"github.com/astaxie/beego"
	"wyweb/controllers"
	"fmt"
	"wyweb/models/user"
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
	datalist := userInfo.GetRow(where)
	ret := make(map[string]interface{})
	if datalist.Error != nil {
		ret["status"] = 0
		ret["msg"] = datalist.Error
		//return json.Marshal(ret)
	}else {

	}

}
