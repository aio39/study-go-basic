package main

import "fmt"

func main() {
	const name string = "aio39"
	change := "a" // 타입 추론, 한번 정한 type은 변경 불가능. func 안에서만 가능  ==> var change string = "39"
	change = "b"
	// name = "aio40"
	fmt.Println(name)
	fmt.Println(change)
}
