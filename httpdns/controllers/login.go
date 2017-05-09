package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"httpdns/models"
	"httpdns/untils"
	"time"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	var (
		user models.User
		err  error
	)
	beego.SetLogger("file", untils.Get_adminlogpath())

	getuser := this.GetString("username")
	getpass := this.GetString("password")

	sess := this.StartSession()

	user, err = models.User_search(getuser)

	if err == nil {
		if getpass == user.Password {
			sess.Set("admin", getuser+","+getpass)

			user.Lasttime = time.Now()

			err = models.User_update(user)
			if err == nil {
				beego.Info("user: " + getuser + " is ok")
				this.Redirect("/httpdns/admin", 302)
			} else {
				beego.Error("user: " + getuser + " update time error")
				this.Redirect("/error/loginerror", 302)
			}
		} else {
			beego.Error("user: " + getuser + " pass is error")
			this.Redirect("/error/loginerror", 302)
		}
	} else {
		beego.Error("user: " + getuser + " is not exist")
		this.Redirect("/error/loginerror", 302)
	}
}

// @router /register [get]
func (this *LoginController) Registerget() {
	this.TplName = "register.html"
}

// @router /register [post]
func (this *LoginController) Registerpost() {
	var user models.User
	beego.SetLogger("file", untils.Get_adminlogpath())

	sess := this.StartSession()

	user.Username = this.GetString("username")
	user.Password = this.GetString("password")
	user.Phonenumber = this.GetString("phonenumber")
	user.Lasttime = time.Now()

	fmt.Println(user)

	err := models.User_insert(user)

	if err != nil {
		this.Abort("507")
		beego.Error("user register error")
	} else {
		beego.Info("user resgister success")
		sess.Set("admin", user.Username+user.Password)

		this.Redirect("/httpdns/admin", 302)
	}
}

// @router /loginerror [get]
func (this *LoginController) Loginerror() {
	this.TplName = "loginerror.html"
}
