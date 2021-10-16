package main

import (
	"fmt"
	"time"
)

func main() {
	// Channel 생성
	channel := make(chan string)
	people := []string{"lee", "kim", "a", "b", "c"}
	for _, person := range people {
		go isSexy(person, channel)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-channel)
	}
}

func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 5)
	channel <- person + " is sexy"
}
