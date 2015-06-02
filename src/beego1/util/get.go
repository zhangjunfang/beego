package util

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var (
	client *http.Client
)

func init() {
	//20 秒连接, 5 秒IO时间
	client = clientConstructor(20, 25)
}

//以秒为单位,设置连接时间,和连接+读取总时间的超时
//return http.clien
func clientConstructor(connectTimeOut int, totalTimeOut int) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				deadline := time.Now().Add(time.Second * time.Duration(totalTimeOut))
				c, err := net.DialTimeout(network, addr, time.Second*time.Duration(connectTimeOut))
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
}

func Get(url string) ([]byte, error) {
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	ret, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
