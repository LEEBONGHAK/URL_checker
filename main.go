package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	// map을 초기화 하는 방법
	results := make(map[string]string) // results := map[string]string{}

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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

// hit: 인터넷 웹 서버의 파일 1개에 접속하는 것
// hitURL: url이 접속 가능한지 판단 함수
func hitURL(url string) error {
	fmt.Println("Checking ", url)
	//http.Ger(): http request를 한 결과를 가저오는 함수
	response, err := http.Get(url)
	if err != nil || response.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
