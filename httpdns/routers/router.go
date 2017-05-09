package routers

import (
	"github.com/astaxie/beego"
	"httpdns/controllers"
)

// func init() {
// 	beego.Router("/", &controllers.MainController{})
// 	beego.Router("/login", &controllers.LoginController{})
// 	beego.Router("/admin", &controllers.AdminController{})
// 	beego.Router("/register", &controllers.AdminController{})
// }

func init() {
	beego.Router("/", &controllers.MainController{})
	api :=
		beego.NewNamespace("/api",
			beego.NSInclude(&controllers.ApiController{}),
		)
	userns :=
		beego.NewNamespace("/user",
			beego.NSRouter("/login", &controllers.LoginController{}),
			beego.NSInclude(&controllers.LoginController{}),
		)

	httpdnsns :=
		beego.NewNamespace("/httpdns",
			beego.NSInclude(&controllers.AdminController{}),
		)

	errorns :=
		beego.NewNamespace("/error",
			beego.NSInclude(&controllers.LoginController{}),
		)

	beego.AddNamespace(api)
	beego.AddNamespace(userns)
	beego.AddNamespace(httpdnsns)
	beego.AddNamespace(errorns)
}
