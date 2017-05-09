package controllers

import (
	"github.com/astaxie/beego"
	"httpdns/models"
)

type ErrController struct {
	beego.Controller
}

func (this *ErrController) Error401() {
	err := models.Error{Status: "401", Msg: "Page not found"}
	this.Data["content"] = &err
	this.TplName = "error.tpl"
}

func (this *ErrController) Error403() {
	err := models.Error{Status: "403", Msg: "Page not found"}
	this.Data["content"] = &err
	this.TplName = "error.tpl"
}

func (this *ErrController) Error404() {
	err := models.Error{Status: "404", Msg: "接口不存在"}
	this.Data["content"] = &err
	this.TplName = "error.tpl"
}

func (this *ErrController) Error500() {
	err := models.Error{Status: "500", Msg: "服务器内部错误，请联系741361303@qq.com进行修复"}
	this.Data["content"] = &err
	this.TplName = "error.tpl"
}

func (this *ErrController) Error503() {
	err := models.Error{Status: "503", Msg: "Page not found"}
	this.Data["content"] = &err
	this.TplName = "error.tpl"
}

func (this *ErrController) Error506() {
	err := models.Error{Status: "506", Msg: "Sql find error"}
	this.Data["content"] = &err
	this.TplName = "error.tpl"
}

func (this *ErrController) Error507() {
	err := models.Error{Status: "507", Msg: "Register Error"}
	this.Data["content"] = &err
	this.TplName = "error.tpl"
}
