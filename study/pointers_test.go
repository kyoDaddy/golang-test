package study

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	a := 2
	b := a  // copy value
	c := &a // copy memory address
	a = 10
	// & : 메모리 주소 보는 방법
	fmt.Println(&a, &b, &c)
	// * : 무엇인가를 메모리 주소에 접근하도록 하는 방법
	fmt.Println(a, b, *c)
	*c = 20
	fmt.Println(a, b, *c)

}
