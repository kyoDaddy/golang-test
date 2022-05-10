package study

import (
	"fmt"
	_ "fmt"
	"testing"
)

func TestVariadic(t *testing.T) {

	say("This", "is", "a", "book")
	say("Hi")

	count, total := sum(1, 7, 3, 5, 9)
	fmt.Println(count, total)
}

/* variadic function (가변인자함수)
함수에 정해진 수의 인자가 아닌 N개 (N은 0이상의 정수) 의 인자를 전달하고 싶을 때가 있다.
ellipsis 를 통해 가능
*/
func say(msg ...string) {
	//for _, s := range msg {
	for i, s := range msg {
		fmt.Println(i, s)
	}
}

/* 함수 return 지정
func <func name>param1<param type>, ...param2<param type>) <return type> { }
1. 복수개의 값을 리턴할 수도 있다.
2. 근데 함수 사용할 때 함수의 return type에 따라 그 수만큼 변수만들어서 값 가져와야 한다는 것 까먹지 말자.

*/
func sum(nums ...int) (int, int) {
	s := 0     // 합계
	count := 0 // 요소 갯수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

/*
3. Named Return Parameter : 특이하게 return variable를 미리 지정해 놓을 수도 있다.
3-1) return statement가 어떠한 argument를 가지지 않고 named return value를 리턴하는데, 이를 "naked" return이라고 한다.
3-2) 어쨌든 함수가 끝나면 반환되는 값은 있기 때문에 return을 어딘가는 꼭 써줘야 함.
*/
func sum2(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}

/* 익명항수
1. 함수명을 갖지 않는 함수를 말한다. function literal 이라고 한다.
2. 보통 함수 전체가 변수에 할당되거나, 다른 함수의 parameter에 직접 정의되곤 한다.
*/
func TestAnonymous(t *testing.T) {
	sum := func(n ...int) int { // 익명함수 정의
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5) // 익명함수 호출
	println(result)
}

/* Functions as type
1. 함수의 원형을 하나의 type으로 만들 수 있다.
2. delegate : 함수의 원형을 정의하고 함수를 타 메서드에 전달하고 리턴받는 기능
*/
// 원형 정의
type calculator func(int, int) int

// calculator 원형 사용
func calc(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

/* Closure
1. 함수 바깥에 있는 변수를 참조하는 function value
2. 함수 밖의 변수를 함수 안에서 읽고 쓸 수 있다.
3. closure를 사용하면 지역 변수가 소멸되지 않고, 함수가 호출될 때마다 계속 그 변수를 사용할 수 있다.
4. 함수가 선언될 때의 환경을 유지하여, 프로그램의 흐름을 변수에 저장할 수 있게한다.
*/
func TestClosure(t *testing.T) {
	next := nextValue()

	println(next())
	println(next())
	println(next())

	anotherNext := nextValue()
	println(anotherNext())
	println(anotherNext())
}

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
