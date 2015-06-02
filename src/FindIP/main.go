package main

import (
	"compress/gzip"
	//"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	//"sync"
	//"github.com/PuerkitoBio/goquery"
)

type config struct {
	PostUrl           string   `toml:"postUrl"`
	DNS               []string `toml:"dns"`
	Domain            []string
	MainChannelNumber int
}

var (
	mainChannel = make(chan int, 5) // 主线程
	//wg          = sync.WaitGroup{}  // 用于等待所有 goroutine 结束
)

func main() {
	//读取配置文件
	var conf config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		//fmt.Println(err)
		return
	}
	//设置线程数
	mainChannel = make(chan int, conf.MainChannelNumber)
	//设置日志级别
	SetLogInfo()
	//循环
	for _, domain := range conf.Domain {
		for _, dns := range conf.DNS {
			//添加参数
			v := url.Values{}
			v.Add("host", domain)
			v.Add("server", dns)
			v.Add("type", "0")
			v.Add("callback", "jQuery1710689435072010383_1386686057249")
			//多进程
			mainChannel <- 1
			//wg.Add(1)
			//Info("mainChannel", 1, wg)
			go PostUrl(&v, &conf)
		}
		Info("-------------", domain, "--done!------------")
	}

	//等待完成
	//wg.Wait()
	Info("finished!")
}

//发生post消息
func PostUrl(values *url.Values, conf *config) {
	defer func() {
		<-mainChannel
		//wg.Done()
	}()
	client := &http.Client{}
	//fmt.Println("djflkdf===", "POST", conf.PostUrl, strings.NewReader(values.Encode()))
	reqest, err := http.NewRequest("POST", conf.PostUrl, strings.NewReader(values.Encode()))
	if err != nil {
		Error("Error: %s \n", err)
	}
	addReqestHeader(reqest)
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(reqest)
	if err != nil {
		Error("Error: %s \n", err)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		Info(getResponseBody(response))
	} else {
		Warn("  Code :", response.StatusCode, values)
	}
}
func addReqestHeader(request *http.Request) {
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Accept-Charset", "utf-8;q=0.7,*;q=0.3")
	//request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.76 Safari/537.36`)
}
func getResponseBody(resp *http.Response) string {
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

//设置log相关
func SetLogInfo() {
	// debug 1, info 2
	SetLevel(2)
	SetLogger("console", "")
	SetLogger("file", `{"filename":"log.log"}`)
}
