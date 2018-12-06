package user

import "github.com/astaxie/beego/orm"

type User struct {
	Id       int    `"orm:column(id);auto;type(uint);pk"`
	Name     string `"orm:column(name);size(30);type(string);unique;"`
	Password string `"orm:column(password);size(30);type(string);"`
	Role     string `"orm:column(role);size(50);type(string)"`
	Status   int	`"orm:column(status);size(4);type(tinyint)"`
}

type UserDataRet struct {
	Total	int
	Error	error
	Data	User
}

func init() {
	orm.RegisterModel(new(User))
}

func (this *User)GetRow(where string) UserDataRet {
	where =" where 1=1 "+where
	datalist:=UserDataRet{
		Total:0,
		Data:User{},
	}
	o:=orm.NewOrm()

	err:=o.Raw("SELECT * from user"+where).QueryRow(&datalist.Data)

	if err!=nil {
		datalist.Error=err
	}

	return datalist
}
