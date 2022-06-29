package something

import "fmt"

func sayBye() { // private
	fmt.Println("Bye")
}

func SayHello() { // public
	fmt.Println("Hello")
}
