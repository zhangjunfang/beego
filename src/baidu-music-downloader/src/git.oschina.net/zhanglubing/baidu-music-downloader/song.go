package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Song struct {
	id     string // ID
	name   string // 歌名
	author string // 歌手
	album  string // 专辑
	rate   int    // 比特率
	link   string // 下载连接
	suffix string // 后缀
}

func (s *Song) getFileName(format string) string {
	return s._format(format, s.author+" - "+s.name+"."+s.suffix)
}

func (s *Song) _format(format string, def string) string {
	if strings.TrimSpace(format) == "" {
		return def
	}
	result := strings.Replace(format, "{name}", s.name, -1)
	result = strings.Replace(result, "{author}", s.author, -1)
	result = strings.Replace(result, "{album}", s.album, -1)
	result = strings.Replace(result, "{rate}", strconv.Itoa(s.rate), -1)
	result = strings.Replace(result, "{suffix}", s.suffix, -1)
	return result
}

func (s *Song) show() {
	fmt.Println("{")
	fmt.Printf("\tid: '%s',\n", s.id)
	fmt.Printf("\tname: '%s',\n", s.name)
	fmt.Printf("\tauthor: '%s',\n", s.author)
	fmt.Printf("\talbum: '%s',\n", s.album)
	fmt.Printf("\trate: '%d',\n", s.rate)
	fmt.Printf("\tlink: '%s'\n", s.link)
	fmt.Printf("\tsuffix: '%s'\n", s.suffix)
	fmt.Println("}")
}
