package main

import "fmt"

func main() {
	names := [5]string{"aio", "miku", "rin"} // Array
	names[3] = "3"
	names[4] = "4"

	infinite := []string{} // Slice
	infinite = append(infinite, "aio")

	fmt.Println(names, infinite)

}
