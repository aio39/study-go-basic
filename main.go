package main

import (
	"fmt"
	"time"
)

func main() {
	go cuteCount("aio")
	cuteCount("miku")
	// go루틴이 실행중이더라도 main 함수가 종료되면 프로그램이 종료
}

func cuteCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is cute", i)
		time.Sleep(time.Second)
	}

}
