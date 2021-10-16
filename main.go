package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request Failed")

func main() {
	// map을 만들면서 초기화 하는 방법
	results := make(map[string]string) // results := map[string]string{}

	// Channel 생성
	ch := make(chan result)

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
		go hitURL(url, ch)
	}

}

// hit: 인터넷 웹 서버의 파일 1개에 접속하는 것
// hitURL: url이 접속 가능한지 판단 함수
func hitURL(url string, ch chan<- result) { // chan<- : 이 채널은 보내는 것만 가능하다는 것을 말하는 것
	fmt.Println("Checking ", url)
	//http.Ger(): http request를 한 결과를 가저오는 함수
	response, err := http.Get(url)
	status := "OK"
	if err != nil || response.StatusCode >= 400 {
		status = "FAILED"
	}
	ch <- result{url: url, status: status}
}
