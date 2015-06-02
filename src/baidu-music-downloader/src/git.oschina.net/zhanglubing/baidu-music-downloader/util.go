package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Fetch(url string) string {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.6,en;q=0.4")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("DNT", "1")
	request.Header.Set("Host", "music.baidu.com")
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.66 Safari/537.36")
	request.Header.Set("Cookie", config.cookie)
	client := &http.Client{}
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		return string(body)
	}
	return ""
}

func _fetchSongInfo(s *Song) {
	if s == nil || strings.TrimSpace(s.id) == "" {
		return
	}

	url := "http://music.baidu.com/song/" + s.id + "/download?__o=%2Fsong%2F" + s.id
	html := Fetch(url)

	re_login := regexp.MustCompile("登录|验证码")
	login := re_login.FindString(html)
	if len(login) >= 2 {
		log.Fatalln("The cookie has expired.")
		return
	}

	// link
	re_download_link := regexp.MustCompile("link=(.*\\.(\\w*)\\?xcode=[\\w\\d]*)\" +id=\"(\\d*)\"")
	links := re_download_link.FindAllStringSubmatch(html, -1)
	size := len(links)
	if size >= 1 {
		s.link = links[size-1][1]
		s.suffix = links[size-1][2]
		s.rate, _ = strconv.Atoi(links[size-1][3])
	}

	s.name = _getMetaInfo(html, "song_title: \"(.*)\",")
	s.author = _getMetaInfo(html, "singer_name: \"(.*)\",")
	s.album = _getMetaInfo(html, "album_name: \"(.*)\",")
}

func _getMetaInfo(html string, regex string) string {
	re := regexp.MustCompile(regex)
	str := re.FindStringSubmatch(html)
	if len(str) == 2 {
		return str[1]
	}
	return "Unkown"
}

func ParseSongList(html string) []string {
	re_song_list := regexp.MustCompile("<li +data-songitem([\\s\\S]+?)</li>")
	return re_song_list.FindAllString(html, -1)
}

func ParseSongId(html string) Song {
	song := Song{}
	id := regexp.MustCompile("('|&quot;)sid('|&quot;): *('|&quot;)*(\\d+)('|&quot;)*").FindStringSubmatch(html)
	if len(id) == 6 {
		song.id = id[4]
	} else {
		log.Println("Failed to parse song info.")
	}
	return song
}

func Download(c chan string, song Song) {

	_fetchSongInfo(&song)

	file := song.getFileName(config.path)

	defer func() {
		c <- song.id
	}()

	array := strings.Split(file, "/")
	path := ""
	name := ""
	if len(array) >= 2 {
		path = strings.Join(array[:len(array)-1], "/")
		name = strings.Join(array[len(array)-1:], "")
	} else {
		path = ""
		name = file
	}

	log.Println("Downloading:", name)

	if IsFileExist(file) {
		log.Println("File already exist.", name)
		return
	}

	if strings.TrimSpace(song.link) == "" {
		log.Println("Link is empty.", name)
		return
	}
	if song.rate < config.min_rate {
		log.Println("DO not download the rateless version.", song.rate, "kbps")
		return
	}

	resp, err := http.Get(song.link)
	defer resp.Body.Close()
	if err != nil {
		log.Println("Failed to download.", err)
		return
	}

	e1 := os.MkdirAll(path, 0777)
	if e1 != nil {
		log.Println("Failed to create directory.", e1)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	e2 := ioutil.WriteFile(file, data, 0644)
	if e2 != nil {
		log.Println("Failed to write file.", e2)
	}

	log.Println("Finished: ", name)

}

func IsFileExist(fpath string) bool {
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		return false
	}
	return true
}
