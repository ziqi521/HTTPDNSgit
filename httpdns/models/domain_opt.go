package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	// "httpdns/untils/dataopt"
)

type DomainInfo struct {
	Id         string
	Domainname string
	Ipname     string
	Weight     string
	Isp        string
}

type GetWeightList struct {
	Domainname string
	Ipname     string
	Weight     string
}

func WeightGet() []GetWeightList {
	var weightlist []GetWeightList
	o := orm.NewOrm()

	_, err := o.Raw("select domainname, group_concat(ipname) AS ipname, group_concat(weight) AS weight from ip group by domainname;").QueryRows(&weightlist)
	if err == nil {
		return weightlist
	} else {
		return weightlist
	}
}

func Domain_searchall() ([]DomainInfo, error) {
	var domaininfo []DomainInfo
	o := orm.NewOrm()

	_, err := o.Raw("select domainname,group_concat(ipname) AS ipname,group_concat(weight) AS weight,group_concat(isp) AS isp from ip group by domainname;").QueryRows(&domaininfo)
	if err == nil {
		return domaininfo, nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}

func Domain_search(domainname string) DomainInfo {
	var domaininfo DomainInfo

	o := orm.NewOrm()
	err := o.Raw("select domainname,group_concat(ipname) AS ipname,group_concat(weight) AS weight,group_concat(isp) AS isp from ip where domainname=? group by domainname;", domainname).QueryRow(&domaininfo)
	if err == nil {
		return domaininfo
	} else {
		fmt.Println("Domain_search errors")
		return domaininfo
	}
}

func Domain_select(ipname string) (DomainInfo, error) {
	var domaininfo DomainInfo
	o := orm.NewOrm()
	err := o.Raw("select domainname,ipname,weight,isp from ip where ip.ipname=?", ipname).QueryRow(&domaininfo)
	if err == nil {
		return domaininfo, nil
	} else {
		return domaininfo, err
	}

}

func Domain_add(domaininfo DomainInfo) error {
	o := orm.NewOrm()

	_, err1 := o.Raw("insert into domain (domainname) values(?)", domaininfo.Domainname).Exec()
	_, err2 := o.Raw("insert into ip (ipname, weight, domainname, isp) values(?, ?, ?, ?)", domaininfo.Ipname, domaininfo.Weight, domaininfo.Domainname, domaininfo.Isp).Exec()

	if err1 == nil && err2 == nil {
		return nil
	} else {
		return errors.New("Insert domain error!")
	}
}

func Domain_addip(domaininfo DomainInfo) error {
	o := orm.NewOrm()

	_, err := o.Raw("insert into ip (ipname, weight, domainname, isp) values(?, ?, ?, ?)", domaininfo.Ipname, domaininfo.Weight, domaininfo.Domainname, domaininfo.Isp).Exec()

	if err == nil {
		return nil
	} else {
		return errors.New("Insert ip error!")
	}
}

func Domain_delete(domainName string) error {
	o := orm.NewOrm()

	_, err1 := o.Raw("delete from domain where domainname = ?", domainName).Exec()
	_, err2 := o.Raw("delete from ip where domainname = ?", domainName).Exec()

	if err1 == nil && err2 == nil {
		return nil
	} else if err1 != nil {
		return err1
	} else if err2 != nil {
		return err2
	} else {
		return errors.New("delete error")
	}
}

func Domain_delete_ip(ipName string) error {
	o := orm.NewOrm()

	_, err := o.Raw("delete from ip where ipname = ?", ipName).Exec()
	if err == nil {
		return nil
	} else {
		return errors.New("delete ip error")
	}
}

func Domain_update(domain DomainInfo, ipname string) error {
	o := orm.NewOrm()

	_, err := o.Raw("update ip set ip.ipname=?, ip.weight=?, ip.isp=? where ip.ipname=?", domain.Ipname, domain.Weight, domain.Isp, ipname).Exec()
	if err == nil {
		return nil
	} else {
		return errors.New("update error")
	}
}
