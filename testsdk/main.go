package main

import (
	"fmt"
	"httpdnssdk"
)

// func main() {
// 	var user httpdnssdk.User
// 	user.LocalIp = "113.134.77.93"
// 	user.Domain = "www.baidu.com"

// 	ipinfo := httpdnssdk.Getid(user)
// 	fmt.Println(ipinfo)
// }

// func main() {
// 	domain := httpdnssdk.Getalldomain()
// 	fmt.Println(domain)
// }

func main() {
	domain := httpdnssdk.Getonedomain("www.baidu.com")
	fmt.Println(domain)
}
