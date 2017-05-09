package dataopt

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Ip         string `form:"-"`
	Domainname string `form:"domainname"`
}

type Ipinfo struct {
	Iplist []string
}

type Weightinfo struct {
	Domainname string
	Weight     Weightfin
}

type Weightfin struct {
	Ip              []string
	Weight          []string
	CurrentWeight   []string
	EffectiveWeight []string
}

var Records [][]string
var lenRecords int
var W []Weightinfo

//把IP库读到内存常驻
func Cache_ip() {

	inputFile, inputError := os.Open("conf/ip.csv")
	if inputError != nil {
		fmt.Println(inputError)
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	r := csv.NewReader(inputReader)

	Records, inputError = r.ReadAll()
	if inputError != nil {
		fmt.Println(inputError)
	}

	lenRecords = len(Records) - 1
}

//在IP库中查找用户IP的信息
func Search_ip(ipname string) []string {
	var (
		lo       int = 0
		mid      int = 0
		midsmall int = 0
		ip       int = Stringip_int(ipname)
		tmp      []string
	)
	if ipname != "" {
		for lo < lenRecords {
			mid = lo + (lenRecords-lo)/2
			midsmall, _ = strconv.Atoi(Records[mid][0])
			if ip-midsmall >= 0 && ip-midsmall <= 255 {
				return Records[mid]
			} else if ip < midsmall {
				lenRecords = mid - 1
			} else if ip > midsmall {
				lo = mid + 1
			}
		}
	}

	return tmp
}

//字符串按照逗号分隔为字符串数组
func Str_tolist(s string) []string {
	return strings.Split(s, ",")
}

//字符串按照逗号分隔为整形数组
func Str_tointlist(s string) []int {
	var (
		strlist []string
		intlist []int
		i       int = 0
	)
	strlist = strings.Split(s, ",")
	for ; i < len(strlist); i++ {
		tmp, err := strconv.Atoi(strlist[i])
		if err == nil {
			intlist = append(intlist, tmp)
		}
	}
	return intlist
}

func Ipku_isp(ipSearch []string) []string {
	var (
		tmp []string
		i   int = 4
	)
	if len(ipSearch) != 0 {
		for ; i < len(ipSearch); i++ {
			tmp = append(tmp, ipSearch[i])
		}
	}

	return tmp
}

//把ISP的sql信息分开
func Strisp_tolist(s string) [][]string {
	var (
		a []string
		b [][]string
		i int = 0
	)
	a = strings.Split(s, ",")

	for ; i < len(a); i++ {
		b = append(b, strings.Split(a[i], "，"))
	}
	return b
}
