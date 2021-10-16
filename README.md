# URL Checker  
  
making url checker with go  
Lecture : [쉽고 빠른 Go 시작하기](https://nomadcoders.co/go-for-beginners)  
  
## Learned  
  
### Goroutines  
  
Goroutines : 기본적으로 다른 함수와 동시에 실행시키는 함수 (병렬 처리 가능)  
병렬 처리하고 싶은 함수 앞에 `go` 사용  
main()이 실행되는 동안만 유효 즉, main()이 끝나면 Goroutine 소멸 또한 main()은 Goroutine을 기다리지 않음  
```
package main

import (
	"fmt"
	"time"
)

func main() {
	go sexyCount("lee")
	sexyCount("kim")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
```  
  
## Channel  
  
Channel은 파이프 같은 역할 => main() 과 Goroutine과의 소통을 할 수 있게 해줌  
  
Channel 만들기: `(channel name) := make(chan (main()과 주고 받고 싶은 data type))`  
example: `channel := make(chan bool)`  
  
Channel 사용하기 (`<-` : Channel의 데이터를 주고 받기)
```
package main

import (
	"fmt"
	"time"
)

func main() {
	// Channel 생성
	channel := make(chan bool)
	people := [2]string{"lee", "kim"}
	for _, person := range people {
		go isSexy(person, channel)
	}
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func isSexy(person string, channel chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	channel <- true
}
```  