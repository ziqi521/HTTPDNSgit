package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["httpdns/controllers:AdminController"] = append(beego.GlobalControllerRouter["httpdns/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/admin`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:AdminController"] = append(beego.GlobalControllerRouter["httpdns/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Updateinfo",
			Router: `/admin/update/:ipname`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:AdminController"] = append(beego.GlobalControllerRouter["httpdns/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/admin/updateinfo/:ipname`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:AdminController"] = append(beego.GlobalControllerRouter["httpdns/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Delete_domain",
			Router: `/admin/delete`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:AdminController"] = append(beego.GlobalControllerRouter["httpdns/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Delete_ip",
			Router: `/admin/deleteip`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:AdminController"] = append(beego.GlobalControllerRouter["httpdns/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Add_domainhtml",
			Router: `/admin/add`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:AdminController"] = append(beego.GlobalControllerRouter["httpdns/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Add_Newdomain",
			Router: `/admin/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:AdminController"] = append(beego.GlobalControllerRouter["httpdns/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Add_iphtml",
			Router: `/admin/addip/:domainname`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:AdminController"] = append(beego.GlobalControllerRouter["httpdns/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Add_Newip",
			Router: `/admin/addip`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:ApiController"] = append(beego.GlobalControllerRouter["httpdns/controllers:ApiController"],
		beego.ControllerComments{
			Method: "Getid",
			Router: `/getid`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:ApiController"] = append(beego.GlobalControllerRouter["httpdns/controllers:ApiController"],
		beego.ControllerComments{
			Method: "Getalldomain",
			Router: `/getalldomain`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:ApiController"] = append(beego.GlobalControllerRouter["httpdns/controllers:ApiController"],
		beego.ControllerComments{
			Method: "Getonedomain",
			Router: `/getdomain/:domainname`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:ApiController"] = append(beego.GlobalControllerRouter["httpdns/controllers:ApiController"],
		beego.ControllerComments{
			Method: "Get_Statices",
			Router: `/getstatices`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:LoginController"] = append(beego.GlobalControllerRouter["httpdns/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Registerget",
			Router: `/register`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:LoginController"] = append(beego.GlobalControllerRouter["httpdns/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Registerpost",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["httpdns/controllers:LoginController"] = append(beego.GlobalControllerRouter["httpdns/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Loginerror",
			Router: `/loginerror`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
