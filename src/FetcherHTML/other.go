package main

import (
	"compress/gzip"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"net/http"
	"net/url"
)

type config struct {
	ThemesUrl  string
	Proxy      bool
	ProxyURL   string
	LogLevel   int
	SaveFolder string
	IndexUrl   string
}

var (
	conf   config
	client = &http.Client{}
)

//设置log相关
func SetLogInfo() {
	// debug 1, info 2
	SetLevel(2)
	SetLogger("console", "")
	SetLogger("file", `{"filename":"log.log"}`)
}

//读取配置文件
func ReadConfig() {
	//读取配置文件
	if _, err := toml.DecodeFile("config.ini", &conf); err != nil {
		fmt.Println("配置文件错误:", err)
		Error(err)
		return
	}
	//再次设置日志级别
	SetLevel(conf.LogLevel)

	//如果有代理，设置代理
	if conf.Proxy {
		pr, err := url.Parse(conf.ProxyURL)
		if err != nil {
			Error("url.Parse(conf.ProxyURL):", err)
			return
		}
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(pr),
			},
		}
	}

	u, err := url.Parse(conf.ThemesUrl)
	if err != nil {
		Error("Parse(conf.ThemesUrl):", err)
	}
	conf.SaveFolder = conf.SaveFolder + u.Host + "/"
}

//添加消息头
func AddReqestHeader(request *http.Request) {
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Accept-Charset", "utf-8;q=0.7,*;q=0.3")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.72 Safari/537.36")
}

//获取响应
func GetResponse(targetUrl string) (resp *http.Response, err error) {
	Debug("GetResponse:", targetUrl)
	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		Error("GetResponse:", targetUrl, err)
		return
	}
	AddReqestHeader(req)
	return client.Do(req)
}

//获取内容
func GetResponseBody(resp *http.Response) string {
	var body []byte
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
		}
		defer reader.Close()
		bodyByte, err := ioutil.ReadAll(reader)
		if err != nil {
		}
		body = bodyByte
	default:
		bodyByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
		}
		body = bodyByte
	}
	return string(body)
}
