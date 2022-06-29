package main

import "fmt"

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers { // range는 for하고만 쓰이는 문법
		// fmt.Println(index, number)
		total += number
	}

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	return total
}

func main() {
	total := superAdd(1, 2, 3, 4, 5)
	fmt.Println(total)
}
