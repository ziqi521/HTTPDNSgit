package httpdnssdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Info struct {
	Ip   string
	Port string
	Uri  string
	Url  string
}

type User struct {
	LocalIp string
	Domain  string
}

type Ipinfo struct {
	Iplist []string
}

type Apidomain struct {
	Domain string
	Ip     []string
}

func Getid(user User) Ipinfo {
	var (
		info   Info
		ipinfo Ipinfo
	)
	info.Ip = "localhost"
	info.Port = "8080"
	info.Uri = "/api/getid"

	info.Url = "http://" + info.Ip + ":" + info.Port + info.Uri
	body, err1 := json.Marshal(user)
	if err1 == nil {
		resp, err3 := http.Post(info.Url, "application/json;charset=utf-8", bytes.NewBuffer(body))
		result, err2 := ioutil.ReadAll(resp.Body)
		if err3 == nil {
			if err2 == nil {
				errparse := json.Unmarshal(result, &ipinfo)
				if errparse == nil {
					return ipinfo
				} else {
					return ipinfo
				}
			} else {
				return ipinfo
			}
		} else {
			return ipinfo
		}

		return ipinfo
	} else {
		fmt.Println(err1)
		return ipinfo
	}

}

func Getalldomain() []Apidomain {
	var (
		info   Info
		domain []Apidomain
	)
	info.Ip = "localhost"
	info.Port = "8080"
	info.Uri = "/api/getalldomain"

	info.Url = "http://" + info.Ip + ":" + info.Port + info.Uri

	resp, err1 := http.Get(info.Url)
	if err1 == nil {
		body, err2 := ioutil.ReadAll(resp.Body)
		if err2 == nil {
			err3 := json.Unmarshal(body, &domain)
			if err3 == nil {
				return domain
			}
		}
	}
	return domain
}

func Getonedomain(domainname string) Apidomain {
	var (
		info   Info
		domain Apidomain
	)
	info.Ip = "localhost"
	info.Port = "8080"
	info.Uri = "/api/getdomain/" + domainname

	info.Url = "http://" + info.Ip + ":" + info.Port + info.Uri

	info.Url = "http://" + info.Ip + ":" + info.Port + info.Uri

	resp, err1 := http.Get(info.Url)
	if err1 == nil {
		body, err2 := ioutil.ReadAll(resp.Body)
		if err2 == nil {
			err3 := json.Unmarshal(body, &domain)
			if err3 == nil {
				return domain
			}
		}
	}
	return domain
}
