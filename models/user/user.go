package user

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int    `"orm:column(id);auto;type(uint);pk"`
	Name       string `"orm:column(name);size(30);type(string);unique;"`
	Password   string `"orm:column(password);size(30);type(string);"`
	Role       string `"orm:column(role);size(50);type(string)"`
	Status     int    `"orm:column(status);size(4);type(tinyint)"`
	CreateTime string    `"orm:column(createtime);size(30);type(string)"`
}

type UserDataRet struct {
	Total int
	Error error
	Data  []User
}

type Respose struct {
	Msg    string
	Status int
	Error  error
}

func init() {
	orm.RegisterModel(new(User))
}

func (this *User) GetRow(where string) (row User, err error) {
	where = " where 1=1 " + where
	row = User{}
	o := orm.NewOrm()

	err = o.Raw("SELECT * from user" + where).QueryRow(&row)

	if err != nil {
		return row, err
	}
	return row, err
}

func (this *User) GetUsers(where string, sqlOrderLimit string) UserDataRet {

	datalist := UserDataRet{
		Total: 0,
		Data:  []User{},
	}

	mywhere := " where 1=1 "
	if where != "" {
		mywhere += "and " + where
	}

	o := orm.NewOrm()
	err := o.Raw("select count(*) from user " + mywhere).QueryRow(&datalist.Total)

	if err == nil || datalist.Total > 0 {
		_, err = o.Raw("SELECT * FROM user " + mywhere + sqlOrderLimit).QueryRows(&datalist.Data)
		datalist.Error = err
	} else {
		datalist.Error = err
	}

	return datalist
}

func (this *User) Del(id int) Respose {
	o := orm.NewOrm()

	_, err := o.Delete(&User{Id: id})

	if err != nil {
		return Respose{
			Msg:    "删除失败",
			Error:  err,
			Status: 0,
		}
	} else {
		return Respose{
			Msg:    "删除成功",
			Error:  err,
			Status: 1,
		}
	}

}

func (this *User) Add(user *User) Respose {

	o := orm.NewOrm()

	_, err := o.Insert(user)

	if err != nil {
		return Respose{
			Msg:    "添加失败",
			Error:  err,
			Status: 0,
		}
	} else {
		return Respose{
			Msg:    "添加成功",
			Error:  err,
			Status: 1,
		}
	}
}

func (this *User) Update(user *User) Respose {
	o := orm.NewOrm()

	_, err := o.Update(user,"Name","Role","Status","Password")

	if err != nil {
		return Respose{
			Msg:    "更新失败",
			Error:  err,
			Status: 0,
		}
	} else {
		return Respose{
			Msg:    "更新成功",
			Error:  err,
			Status: 1,
		}
	}
}
