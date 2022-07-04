package main

import (
	"fmt"

	"github.com/aio39/study_go_basic/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "こんにちは"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, err := dictionary.Search(word)
	fmt.Println(hello)
}
