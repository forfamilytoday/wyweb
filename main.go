package main

import (
	_ "wyweb/routers"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"strconv"
	HELPS "wyweb/utils"
	"github.com/robfig/cron"
	"fmt"
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//获取数据库配置信息
	conn := beego.AppConfig.String("conn")
	orm.RegisterDataBase("default", "mysql", conn, 30)

	//设置日志

	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	level := beego.AppConfig.String("level")

	if level == "level" {
		beego.SetLevel(beego.LevelError)
	} else {
		num, _ := strconv.Atoi(level)
		beego.SetLevel(num)
	}
	//首次执行创建数据库，并生生一个锁文件
	HELPS.CreateDB()

	//开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600

}

func main() {
	c := cron.New()

	spec := "0 0 23 * * ?"

	c.AddFunc(spec, func() {
		fmt.Println(11111111111)
		HELPS.Goquery()
	})
	c.Start()
	//开启爬虫
	beego.Run()
}
