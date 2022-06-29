package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int { // a int, b int  // Type이 같은 경우는 한번만 작성하면 됨
	return a * b
}

// func lenAndUpper(name string) ( int, string) {
// 	return len(name), strings.ToUpper(name)
// }

// naked return
func lenAndUpper(name string) (length int, upperCase string) {
	defer fmt.Println("I'm done") // 함수가 끝나고 실
	length = len(name)
	upperCase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// fmt.Println(multiply(2, 3))
	totalLength, upperName := lenAndUpper("Hello") // _로 받으면 ignored value가
	fmt.Println(totalLength, upperName)
	repeatMe("aio", "miku", "ichika", "dare")
}
