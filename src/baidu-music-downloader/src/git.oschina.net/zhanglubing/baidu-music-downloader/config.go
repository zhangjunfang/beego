package main

import (
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
)

type Config struct {
	path     string   // 下载文件格式
	min_rate int      // 最小比特率
	cookie   string   // Cookie
	artists  []string // 歌手
	album    []string // 专辑
	top      []string // 榜单
	tag      []string // 标签、风格
	songlist []string // 网友歌单
}

func (c *Config) read(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panic("Config file missing.")
		return err
	}
	json, err := simplejson.NewJson(data)
	if err != nil {
		log.Panic("Invalid config file.")
		return err
	}

	settings := json.Get("settings")
	c.path, _ = settings.Get("path").String()
	c.min_rate, _ = settings.Get("min_rate").Int()
	c.cookie, _ = json.Get("cookie").String()
	c.artists, _ = json.Get("artists").StringArray()
	c.album, _ = json.Get("album").StringArray()
	c.top, _ = json.Get("top").StringArray()
	c.tag, _ = json.Get("tag").StringArray()
	c.songlist, _ = json.Get("songlist").StringArray()

	return nil
}

func (c *Config) getAllURLList() []string {
	index := 0

	size := len(c.artists) + len(c.album) + len(c.top) + len(c.tag) + len(c.songlist)
	res := make([]string, size)

	copy(res, c.artists)
	index += len(c.artists)
	copy(res[index:], c.album)
	index += len(c.album)
	copy(res[index:], c.top)
	index += len(c.top)
	copy(res[index:], c.tag)
	index += len(c.tag)
	copy(res[index:], c.songlist)
	index += len(c.songlist)

	return res
}
