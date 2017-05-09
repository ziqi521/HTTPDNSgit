package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/mysql"
	"httpdns/controllers"
	_ "httpdns/routers"
	_ "httpdns/untils"
	"httpdns/untils/dataopt"
	"httpdns/untils/encryptor"
)

func init() {
	//初始化数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)

	dbname := encryptor.Decrypt(beego.AppConfig.String("dbname"))
	dbuser := encryptor.Decrypt(beego.AppConfig.String("dbuser"))
	dbpass := encryptor.Decrypt(beego.AppConfig.String("dbpass"))

	orm.RegisterDataBase("default", "mysql", dbuser+":"+dbpass+"@/"+dbname+"?charset=utf8", 30, 30)

	//过滤器
	var FilterUser = func(ctx *context.Context) {
		tmp := ctx.Input.CruSession.Get("admin")
		if tmp == nil && ctx.Request.RequestURI != "/user/login" {
			ctx.Redirect(302, "/user/login")
		} else {
			fmt.Println(tmp)
		}
	}

	beego.InsertFilter("^[^/api].*$", beego.BeforeRouter, FilterUser)

	// 初始化IP库
	dataopt.Cache_ip()
	// 初始化平滑加权轮询数据格式
	controllers.Init_weightlist()
	// 初始化域名映射信息
	controllers.Init_domaininfo()

}

func main() {
	// 注册error处理
	beego.ErrorController(&controllers.ErrController{})
	beego.Run()
}
