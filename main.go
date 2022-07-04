package main

import (
	"fmt"

	"github.com/aio39/study_go_basic/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"hello": "world"}
	dictionary["hello"] = "world"
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition)
}
