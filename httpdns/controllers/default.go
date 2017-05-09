package controllers

import (
	"github.com/astaxie/beego"
	"httpdns/models"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	var tmp []string
	v := this.GetSession("admin")

	if v == nil {
		this.Redirect("/user/login", 302)
	} else {
		if s, ok := v.(string); ok {
			tmp = strings.Split(s, ",")
		}

		user, err := models.User_search(tmp[0])
		if err == nil {
			if tmp[1] == user.Password {
				this.Redirect("/httpdns/admin", 302)
			} else {
				this.Redirect("/user/login", 302)
			}
		} else {
			this.Redirect("/user/login", 302)
		}
	}
	this.Redirect("/httpdns/admin", 302)
}
