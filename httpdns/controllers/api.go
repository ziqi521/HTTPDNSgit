package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"httpdns/models"
	"httpdns/untils"
	"httpdns/untils/dataopt"
)

type ApiController struct {
	beego.Controller
}

// 用户需要发送本身的IP和域名
type User struct {
	Ip         string `form:"ip"`
	Domainname string `form:"domainname"`
}

type Ipinfo struct {
	Iplist []string
}

type Weightinfo struct {
	Domainname string
	Weight     Weightcount
}

type Weightcount struct {
	Ip              []string
	Weight          []int
	CurrentWeight   []int
	EffectiveWeight []int
}

type Parsedomain struct {
	Domain string
	Ip     []string
	Weight []string
	Isp    [][]string
}

type Apidomain struct {
	Domain string
	Ip     []string
}

var W []Weightinfo

var DomainParse []Parsedomain

func Init_weightlist() {
	var getweightlist []models.GetWeightList
	var tmp Weightinfo
	// fmt.Println("初始化一个加权列表")
	getweightlist = models.WeightGet()
	for i := 0; i < len(getweightlist); i++ {
		tmp.Domainname = getweightlist[i].Domainname
		tmp.Weight.Ip = dataopt.Str_tolist(getweightlist[i].Ipname)
		tmp.Weight.Weight = dataopt.Str_tointlist(getweightlist[i].Weight)
		tmp.Weight.EffectiveWeight = dataopt.Str_tointlist(getweightlist[i].Weight)
		for j := 0; j < len(tmp.Weight.Ip); j++ {
			tmp.Weight.CurrentWeight = append(tmp.Weight.CurrentWeight, 0)
		}
		W = append(W, tmp)
	}
	// fmt.Println(W)
}

func Init_domaininfo() {
	var (
		domaininfo []models.DomainInfo
		tmp        Parsedomain
		err        error
		i          int = 0
	)
	domaininfo, err = models.Domain_searchall()
	if err == nil {
		for ; i < len(domaininfo); i++ {
			tmp.Domain = domaininfo[i].Domainname
			tmp.Ip = dataopt.Str_tolist(domaininfo[i].Ipname)
			tmp.Weight = dataopt.Str_tolist(domaininfo[i].Weight)
			tmp.Isp = dataopt.Strisp_tolist(domaininfo[i].Isp)
			DomainParse = append(DomainParse, tmp)
		}
	}

}

func get_domaininfo(domainname string) (tmp Parsedomain) {
	var (
		i int = 0
	)

	for ; i < len(DomainParse); i++ {
		if domainname == DomainParse[i].Domain {
			tmp = DomainParse[i]
			return tmp
		}
	}
	return tmp
}

func get_domainseat(domainname string) int {
	for i := 0; i < len(W); i++ {
		if W[i].Domainname == domainname {
			return i
		}
	}
	return 0
}

func set_weight(domain string, ip string) {
	var total int = 0
	for i := 0; i < len(W); i++ {
		if W[i].Domainname == domain {
			for j := 0; j < len(W[i].Weight.Ip); j++ {
				if W[i].Weight.Ip[j] == ip {
					for k := 0; k < len(W[i].Weight.EffectiveWeight); k++ {
						total += W[i].Weight.EffectiveWeight[k]
					}
					W[i].Weight.CurrentWeight[j] -= total
					break
				}
			}
		}
	}
}

func next_weight(we Weightcount) (ip string) {
	var (
		total      int = 0
		i          int = 0
		bestweight int = 0
		tmp        int = 0
	)

	for ; i < len(we.Ip); i++ {

		we.CurrentWeight[i] += we.EffectiveWeight[i]
		total += we.EffectiveWeight[i]
		if we.EffectiveWeight[i] < we.Weight[i] {
			we.EffectiveWeight[i]++
		}

		if ip == "" || we.CurrentWeight[i] > bestweight {
			ip = we.Ip[i]
			bestweight = we.CurrentWeight[i]
			tmp = i
		}
	}
	we.CurrentWeight[tmp] -= total
	return ip
}

// @router /getid [post]
func (this *ApiController) Getid() {
	untils.Statices_tolog()
	var (
		ipinfo    Ipinfo
		user      User
		getdomain Parsedomain
		i         int = 0
		tmp       int = 0
		domaintmp int = 0
		rerr      models.RetMsg
	)
	beego.SetLogger("file", untils.Get_apilogpath())

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)

	if err == nil {
		var (
			ipSearch    []string
			ipSearchIsp []string
		)

		ipSearch = dataopt.Search_ip(user.Ip)
		ipSearchIsp = dataopt.Ipku_isp(ipSearch)

		getdomain = get_domaininfo(user.Domainname)
		domaintmp = get_domainseat(user.Domainname)

		// 判断走什么算法来返回IPlist的首个IP
		if len(ipSearchIsp) != 0 && ipSearchIsp[0] == "中国" {
			for i = 0; i < len(getdomain.Ip); i++ {
				if ipSearchIsp[1] == getdomain.Isp[i][1] && ipSearchIsp[4] == getdomain.Isp[i][3] {
					tmp = 2
					ipinfo.Iplist = append(ipinfo.Iplist, getdomain.Ip[i])
					set_weight(user.Domainname, getdomain.Ip[i])
				} else if ipSearchIsp[4] == getdomain.Isp[i][3] {
					tmp = 2
					ipinfo.Iplist = append(ipinfo.Iplist, getdomain.Ip[i])
					set_weight(user.Domainname, getdomain.Ip[i])
				} else {
					tmp = 1
				}
			}
		} else {
			tmp = 1
		}

		// 走平滑加权轮询
		if tmp == 1 {
			ipinfo.Iplist = append(ipinfo.Iplist, next_weight(W[domaintmp].Weight))
			for i = 0; i < len(W[domaintmp].Weight.Ip); i++ {
				if W[domaintmp].Weight.Ip[i] != ipinfo.Iplist[0] {
					ipinfo.Iplist = append(ipinfo.Iplist, W[domaintmp].Weight.Ip[i])
				}
			}
			beego.Info("/api/getid from " + user.Ip + " get:" + user.Domainname + " wrr")
		} else if tmp == 2 {
			// 根据用户地理位置和运营商返回最合适的IP
			for i = 0; i < len(W[domaintmp].Weight.Ip); i++ {
				if W[domaintmp].Weight.Ip[i] != ipinfo.Iplist[0] {
					ipinfo.Iplist = append(ipinfo.Iplist, W[domaintmp].Weight.Ip[i])
				}
			}
			beego.Info("/api/getid from " + user.Ip + " get:" + user.Domainname + " geographical position")
		}
	} else {
		this.Data["json"] = rerr.NewRetMsg(510, "get json parse error")
		beego.Error("/api/getid json parse error from " + user.Ip + " get:" + user.Domainname)
	}

	if len(ipinfo.Iplist) != 0 {
		this.Data["json"] = &ipinfo
		this.ServeJSON()
	} else {
		this.Data["json"] = rerr.NewRetMsg(400, "没有查到该域名")
		this.ServeJSON()
	}

}

// @router /getalldomain [get]
func (this *ApiController) Getalldomain() {
	var (
		re    []Apidomain
		retmp Apidomain
		i     int = 0
		j     int = 0
	)
	for ; i < len(DomainParse); i++ {
		var tmp []string
		for ; j < len(DomainParse[i].Ip); j++ {
			tmp = append(tmp, DomainParse[i].Ip[j])
		}
		retmp.Domain = DomainParse[i].Domain
		retmp.Ip = tmp
		re = append(re, retmp)
	}
	beego.SetLogger("file", untils.Get_apilogpath())
	beego.Info("/api/getalldomain requested...")

	this.Data["json"] = &re
	this.ServeJSON()
}

// 返回域名对应全部IP不走算法
// @router /getdomain/:domainname [get]
func (this *ApiController) Getonedomain() {
	var (
		domainname string
		err        models.RetMsg
		tmp        Parsedomain
		re         Apidomain
		i          int = 0
	)
	beego.SetLogger("file", untils.Get_apilogpath())

	domainname = this.Ctx.Input.Param(":domainname")
	tmp = get_domaininfo(domainname)
	re.Domain = tmp.Domain
	for ; i < len(tmp.Ip); i++ {
		re.Ip = append(re.Ip, tmp.Ip[i])
	}
	if len(tmp.Ip) != 0 {
		this.Data["json"] = &re
		this.ServeJSON()
	} else {
		this.Data["json"] = err.NewRetMsg(400, "查无此域名")
		this.ServeJSON()
		beego.Error("/api/getdomain/" + domainname + "domain find error")
	}

	beego.Info("/api/getdomain/" + domainname + "requested")
}

// @router /getstatices [get]
func (this *ApiController) Get_Statices() {
	data, err := untils.Get_Statices()
	if err == nil {
		this.Data["json"] = &data
		this.ServeJSON()
	} else {

	}
}
