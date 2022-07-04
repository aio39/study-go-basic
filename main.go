package main

import (
	"fmt"

	"github.com/aio39/study_go_basic/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	// dictionary.Delete(baseWord)

	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}

	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

}
