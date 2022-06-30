package main

import "fmt"

func main() {
	// a := 2
	// b := a
	// a = 10
	// fmt.Println(a, b) // 10 , 2

	// a := 2
	// b := &a
	// fmt.Println(&a, b, *b) // 0xc0000140a0 0xc0000140a8  , 0x -> 16진법

	a := 2
	b := &a
	*b = 20
	fmt.Println(a, b, *b) // 20, 주소 , 20

}
