package main

import (
	"fmt"

	"github.com/aio39/study_go_basic/something"
)

func main() {
	fmt.Println("hello world!!") // Go에서는 대문자 함수를 Export 해줌
	something.SayHello()
}
