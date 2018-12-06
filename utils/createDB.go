package utils

import (
	_ "wyweb/models/user"
	"github.com/astaxie/beego/orm"
	"fmt"
	"os"
	"github.com/astaxie/beego"
)

func CreateDB() {
	file := "logs/lock.log"

	if FileExist(file) {
		return
	}

	fb, err := os.Create(file)
	defer fb.Close()

	if err != nil {
		beego.Error(err)
		return
	}

	orm.RunCommand()

	// 数据库别名
	name := "default"

	// drop table 后再建表
	force := true

	// 打印执行过程
	verbose := true

	// 遇到错误立即返回
	err1 := orm.RunSyncdb(name, force, verbose)
	if err1 != nil {
		fmt.Println(err1)
	}
}

func FileExist(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}
