package untils

import (
	"github.com/astaxie/beego"
	"runtime"
)

func Get_apilogpath() string {
	if runtime.GOOS == "windows" {
		var tmp string
		tmp = "{\"filename\":\"" + beego.AppConfig.String("logWinApi") + "\"}"
		return tmp
	} else {
		var tmp string
		tmp = "{\"filename\":\"" + beego.AppConfig.String("logLinuxApi") + "\"}"
		return tmp
	}
}

func Get_adminlogpath() string {
	if runtime.GOOS == "windows" {
		var tmp string
		tmp = "{\"filename\":\"" + beego.AppConfig.String("logWinAdmin") + "\"}"
		return tmp
	} else {
		var tmp string
		tmp = "{\"filename\":\"" + beego.AppConfig.String("logLinuxAdmin") + "\"}"
		return tmp
	}
}
