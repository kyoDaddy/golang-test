package study

import (
	"fmt"
	"testing"
)

/* Struct
Go에는 Class/Object/Inheitancew 개념이 없다.
1. field만을 가진다.
*/
type person struct {
	name string
	age  int
}

func TestStruct(t *testing.T) {
	// person 객체 생성
	p0 := person{}
	var p1 person
	p1 = person{"k", 100}
	p2 := person{"l", 25}

	// person 객체 포인터 생성
	p3 := new(person)

	// 필드값 설정
	p0.name = "gg"
	p0.age = 22

	fmt.Println(p0)
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

}
