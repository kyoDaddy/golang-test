package study

import (
	"fmt"
	"testing"
)

/* Method
1. struct가 field만 가지기 때문에 method를 따로 정의해야 한다.
2. method 정의 시 이 method가 어떤 struct에 대한 method인지 명시해 주어야 함
2. func 키워드 바로 뒤에 parentheses ( () ) 안에다가 써준다.
func (<var name> struct type) <func name>() <ret type> { }
*/

// Rect - struct 정의
type Rect struct {
	width, height int
}

/* Rect의 area() 메소드
코드의 (r Rect)처럼 함수명 앞에 타입과 변수명이 붙어 있는 것을 Receiver라고 한다.
요런 reciever가 장착된 함수는 이제 함수가 아니라 메소드가 된다.

이 Receiver도 value receiver가 있고 pointer receiver가 있다.
pointer receiver는 뭐 생각하는 그대로~ 변수 값을 복사해서 가져오는 게 아니라 주소값을 가져오는 거기 때문에 struct 내의 필드 값이 실제로 변경될 수 있다.
근데 재밌는 건, go에서는 value receiver와 point receiver 사이의 전환을 잘 핸들링 해준다는 것.
pointer var가 아니어도 point receiver를 사용할 수 있다.
*/
func (r Rect) area() int {
	return r.width * r.height
}

func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func TestMethod(t *testing.T) {
	rect := Rect{10, 20}
	area := rect.area()   // 메서드 호출
	area2 := rect.area2() // 메서드 호출
	fmt.Println(area)
	fmt.Println(area2)
}
