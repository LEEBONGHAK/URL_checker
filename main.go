package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		hitURL(url)
	}
}

// hit: 인터넷 웹 서버의 파일 1개에 접속하는 것
func hitURL(url string) error {
	fmt.Println("Checking ", url)
	//http.Ger(): http request를 한 결과를 가저오는 함수
	response, err := http.Get(url)
	if err != nil || response.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
