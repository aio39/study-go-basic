package main

import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // variable expression , if 문에서만 쓰이는 if scope 변수를 선언함으로서 가독성 향상
		return false
	} // else{}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
