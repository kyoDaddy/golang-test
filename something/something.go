package something

import (
	"fmt"
)

// private
func sayBye() {
	fmt.Println("Bye")
}

// public
func SayHello() {
	fmt.Println("Hello")
}
