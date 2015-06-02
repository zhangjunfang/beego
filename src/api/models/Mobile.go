package models

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var logger *log.Logger

type Mobile struct {
	//所查号码
	TxtMobile string
	//归属省份
	TxtProvince string
	//归属城市
	TxtCity string
	//城市区号
	TxtAreaCode string
	//城市邮编
	TxtPostCode string
	//运营商
	TxtVNO string
}
type Ip struct {
	Root     xml.Name `xml:ArrayOfString`
	Eelement []*Item
}
type Item struct {
	Root xml.Name `xml:string`
}

//<ArrayOfString xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns="http://WebXml.com.cn/">
//  <string>102.1.0.50</string>
//  <string>非洲地区</string>
//</ArrayOfString>
func init() {
	logger = log.New(os.Stderr, "sst: ", log.Ldate|log.Ltime|log.Lshortfile)
}
func GetMobileDesc(phone string) *Mobile {
	client := &http.Client{}
	fmt.Println("phone", phone)
	resp, err := client.Get(fmt.Sprintf("http://api.showji.com/Locating/www.showji.c.o.m.aspx?m=%s&output=json&callback=querycallback", phone))
	if err != nil {
		logger.Println(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Println(err.Error())
	}
	defer resp.Body.Close()
	html := string(body)
	html = html[strings.Index(html, "{")+1 : strings.Index(html, "}")]
	return parseMobile(strings.Replace(html, "\"", "", -1))
}
func GetIPDesc(ip string) *Ip {
	client := &http.Client{}
	resp, err := client.Get(fmt.Sprintf("http://webservice.webxml.com.cn/WebServices/IpAddressSearchWebService.asmx/getCountryCityByIp?theIpAddress=%s", ip))
	if err != nil {
		logger.Println(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Println(err.Error())
	}
	defer resp.Body.Close()
	fmt.Println("ip::::::", string(body))
	return parseIP(body)
}
func parseIP(ip []byte) *Ip {
	ips := &Ip{}
	xml.Unmarshal(ip, ips)
	return ips
}
func parseMobile(json string) *Mobile {
	mob := &Mobile{}
	for _, v := range strings.Split(json, ",") {
		item := strings.Split(v, ":")
		if item[0] == "Mobile" {
			mob.TxtMobile = item[1]
		}
		if strings.Contains(item[0], "Corp") {
			mob.TxtVNO = item[1]
		}
		if strings.Contains(item[0], "Province") {
			mob.TxtProvince = item[1]
		}
		if strings.Contains(item[0], "City") {
			mob.TxtCity = item[1]
		}
		if strings.Contains(item[0], "AreaCode") {
			mob.TxtAreaCode = item[1]
		}
		if strings.Contains(item[0], "PostCode") {
			mob.TxtPostCode = item[1]
		}
	}
	return mob
}
