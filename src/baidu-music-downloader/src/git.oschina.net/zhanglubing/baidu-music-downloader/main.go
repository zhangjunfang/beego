package main

import (
	"strconv"
)

var (
	batch_size = 5
	config     = Config{}
)

func main() {

	config.read("config.json")

	c := make(chan string, batch_size)

	for i := 0; i < batch_size; i++ {
		c <- "Init-" + strconv.Itoa(i)
	}

	urls := config.getAllURLList()
	for i := 0; i < len(urls); i++ {
		html := Fetch(urls[i])
		songs := ParseSongList(html)
		for j := 0; j < len(songs); j++ {
			song := ParseSongId(songs[j])
			<-c
			go Download(c, song)
		}
	}

	for i := 0; i < batch_size; i++ {
		<-c
	}

}
