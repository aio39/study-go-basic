package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
} // go에는 __init__, super(), constructor() 등이 없

func main() {
	aio := map[string]string{"name": "aio", "age": "17"} // key - value

	fmt.Println(aio)

	for key, value := range aio {
		fmt.Println(key, value)
	}
	// --------------------------

	favFood := []string{"chicken", "chocolateMintIce"}
	// aio2 := person{"aio", 17, favFood} -> 가능하지만 가독성이 나쁨,  두가지 방법 Mix해서 사용 불가
	aio2 := person{name: "aio", age: 17, favFood: favFood}
	fmt.Println(aio2, aio2.age)

}
