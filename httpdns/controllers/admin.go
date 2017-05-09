package controllers

import (
	"github.com/astaxie/beego"
	"httpdns/models"
	"httpdns/untils"
	"httpdns/untils/dataopt"
)

type AdminController struct {
	beego.Controller
}

type TplInfo struct {
	Id         string
	Domainname string
	Ipname     []string
	Weight     []string
	Isp        []string
}

// @router /admin [get]
func (this *AdminController) Get() {
	var (
		domaininfo []models.DomainInfo
		tplinfo    []TplInfo
		err        error
	)
	domaininfo, err = models.Domain_searchall()
	if err == nil {
		for i, _ := range domaininfo {
			var tmp TplInfo
			tmp.Id = domaininfo[i].Id
			tmp.Domainname = domaininfo[i].Domainname
			tmp.Ipname = append(tmp.Ipname, dataopt.Str_tolist(domaininfo[i].Ipname)...)
			tmp.Weight = append(tmp.Weight, dataopt.Str_tolist(domaininfo[i].Weight)...)
			tmp.Isp = append(tmp.Isp, dataopt.Str_tolist(domaininfo[i].Isp)...)

			tplinfo = append(tplinfo, tmp)
		}

		this.Data["domainInfo"] = tplinfo
	} else {
		this.Abort("506")
	}

	this.TplName = "admin.html"
}

// @router /admin/update/:ipname [get]
func (this *AdminController) Updateinfo() {
	var ipname string

	ipname = this.Ctx.Input.Param(":ipname")
	domaininfo, err := models.Domain_select(ipname)
	if err == nil {
		this.Data["domainInfo"] = domaininfo
		this.TplName = "update.html"
	} else {
		this.Abort("506")

	}
}

// @router /admin/updateinfo/:ipname [post]
func (this *AdminController) Update() {
	var (
		updateinfo models.DomainInfo
		ipname     string
	)
	beego.SetLogger("file", untils.Get_adminlogpath())

	ipname = this.Ctx.Input.Param(":ipname")
	updateinfo.Domainname = this.GetString("domainname")
	updateinfo.Ipname = this.GetString("ipname")
	updateinfo.Weight = this.GetString("weight")
	updateinfo.Isp = this.GetString("isp")

	err := models.Domain_update(updateinfo, ipname)

	if err == nil {
		Init_weightlist()
		beego.Info("update " + updateinfo.Domainname + " success")
		this.Redirect("/httpdns/admin", 302)
	} else {
		beego.Error("update " + updateinfo.Domainname + " error")
		this.Abort("506")
	}
}

// @router /admin/delete [post]
func (this *AdminController) Delete_domain() {
	var domainname string
	beego.SetLogger("file", untils.Get_adminlogpath())

	domainname = this.GetString("domainname")
	err := models.Domain_delete(domainname)
	if err == nil {
		Init_weightlist()
		beego.Info("delete domain " + domainname + " success")
		this.Redirect("/httpdns/admin", 302)
	} else {
		beego.Error("delete domain " + domainname + " error")
		this.Abort("506")
	}
}

// @router /admin/deleteip [get]
func (this *AdminController) Delete_ip() {
	var ipname string
	beego.SetLogger("file", untils.Get_adminlogpath())

	ipname = this.Input().Get("ipname")
	err := models.Domain_delete_ip(ipname)
	if err == nil {
		Init_weightlist()
		beego.Info("delete ip " + ipname + " success")
		this.Redirect("/httpdns/admin", 302)
	} else {
		beego.Error("delete ip " + ipname + " error")
		this.Abort("506")
	}
}

// @router /admin/add [get]
func (this *AdminController) Add_domainhtml() {
	this.TplName = "add.html"
}

// @router /admin/add [post]
func (this *AdminController) Add_Newdomain() {
	var domaininfo models.DomainInfo
	beego.SetLogger("file", untils.Get_adminlogpath())

	domaininfo.Domainname = this.GetString("domainname")
	domaininfo.Ipname = this.GetString("ipname")
	domaininfo.Weight = this.GetString("weight")
	domaininfo.Isp = this.GetString("isp")

	err := models.Domain_add(domaininfo)

	if err == nil {
		Init_weightlist()
		beego.Info("add domain " + domaininfo.Domainname + " success")
		this.Redirect("/httpdns/admin", 302)
	} else {
		beego.Error("add domain " + domaininfo.Domainname + " error")
		this.Abort("506")
	}
}

// @router /admin/addip/:domainname [get]
func (this *AdminController) Add_iphtml() {
	var domainname string
	domainname = this.Ctx.Input.Param(":domainname")

	this.Data["domainname"] = domainname
	this.TplName = "addip.html"
}

// @router /admin/addip [post]
func (this *AdminController) Add_Newip() {
	var domaininfo models.DomainInfo
	beego.SetLogger("file", untils.Get_adminlogpath())

	domaininfo.Domainname = this.GetString("domainname")
	domaininfo.Ipname = this.GetString("ipname")
	domaininfo.Weight = this.GetString("weight")
	domaininfo.Isp = this.GetString("isp")

	err := models.Domain_addip(domaininfo)

	if err == nil {
		Init_weightlist()
		beego.Info("add ip " + domaininfo.Ipname + " to " + domaininfo.Domainname + " success")
		this.Redirect("/httpdns/admin", 302)
	} else {
		beego.Error("add ip " + domaininfo.Ipname + " to " + domaininfo.Domainname + " error")
		this.Abort("506")
	}
}
