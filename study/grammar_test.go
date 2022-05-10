package study

import (
	"fmt"
	"math"
	"testing"
)

var k = 1

func TestVariable(t *testing.T) {
	// var <변수명> <type> = <초기값>
	// 타입 추론이 가능하기 때문에 타입을 생략할 수도 있다. (Go가 initial value를 통해 타입을 추론해서 타입을 지정하기 때문..)
	var a int
	var f float32 = 11
	var i, j, k int = 1, 2, 3

	var d = 5

	fmt.Println(a, f, i, j, k, d)

	/* 변수를 선언하는 또 다른 방식으로 Short Assignment Statement ( := ) 를 사용할 수 있다.
	즉, var i= 1 을 쓰는 대신 i := 1 이라고 var 를 생략하고 사용할 수 있다.
	하지만 이러한 표현은 함수(func) 내에서만 사용할 수 있으며, 함수 밖에서는 var를 사용해야 한다.
	Go에서 변수와 상수는 함수 밖에서도 사용할 수 있다.
	*/
	x := 1
	fmt.Println(k, x)

}

func TestConstants(t *testing.T) {
	// const 키워드를 사용하여 상수를 선언
	const hi = "Hi"
	// 여러개 묶어서 한번에 const 지정할 수 있음
	const (
		Sky   = "Blue"
		Rose  = "Red"
		Gyuri = "Awesome"
	)

	// iota 사용하면 상수값을 순차적으로 0부터 부여할 수 있다
	const (
		apple = iota
		grape
		orange
	)

	fmt.Println(Sky)
	fmt.Println(apple, grape, orange)
}

func TestDataType(t *testing.T) {

	// String
	// Back Quote(``)와 Double Quote("")로 선언할 수 있다. (single quote ('') 안 씀)
	// immutable type : 한번 생성되면 수정할 수가 없다.
	rawLiteral := `메롱메롱\n`
	interLiteral := "얼레리꼴레리\n"
	fmt.Println(rawLiteral, interLiteral)

}

func TestTypeConversion(t *testing.T) {
	// 암묵적인 type conversion이 이뤄지지 않는다. 항상 명시적으로 지정해 주어야 한다.
	// type conversion 없을 경우 runtime err가 발생한다.
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)

}

func TestOperator(t *testing.T) {
	a := 6
	b := 5

	/* 사칙연산
	+: 덧셈 (정수, 실수, 문자열)
	-: 뺄셈 (정수, 실수)
	*: 곱셈 (정수, 실수)
	/: 나눗셈 (정수, 실수)
	%: 나머지 (정수만 가능)
	*/
	fmt.Println("Result:", a+b)
	fmt.Println("Result:", a-b)
	fmt.Println("Result:", a*b)
	fmt.Println("Result:", a/b)
	fmt.Println("Result:", a%b)

	/* 비트 연산자
	&: AND 비트연산 (정수)
	|: OR 비트연산 (정수)
	^: XOR 비트연산 (정수)
	&^: 비트 클리어 (정수)
	*/
	fmt.Printf("Result: %08b\n", a&b)
	fmt.Printf("Result: %08b\n", a|b)
	fmt.Printf("Result: %08b\n", a^b)
	fmt.Printf("Result: %08b\n", a&^b)

	/* 시프트 연산자
	<<: 왼쪽 시프트 (양의 정수)
	>>: 오른쪽 시프트 (양의 정수)
	*/
	fmt.Printf("Result: %08b\n", a)
	fmt.Printf("Result: %08b\n", a<<2)
	fmt.Printf("Result: %08b\n", a>>1)

	/* 비교 연산자
	==: 같다
	!=: 다르다
	<: 작다
	>: 크다
	<=: 작거나 같다
	>=: 크거나 같다
	*/
	fmt.Println("Result:", a == b)
	fmt.Println("Result:", a != b)
	fmt.Println("Result:", a < b)
	fmt.Println("Result:", a > b)
	fmt.Println("Result:", a <= b)
	fmt.Println("Result:", a >= b)

	/* 논리 연산자
	&&: AND 연산자
	||: OR 연산자
	!: NOT 연산자
	*/
	c := true
	d := false

	fmt.Println("Result:", c && d)
	fmt.Println("Result:", c || d)
	fmt.Println("Result:", !c)

	// 대입 연산자 : 대입 연산자를 다음과 같이 사용하면, 두 변수의 값을 교환할 수 있다
	var e int = 1
	var f int

	e, f = 1, 2
	e, f = f, e
	println(e, f)

	/* 복합 대입 연산자
	+=
	-=
	*=
	/=
	%=
	&=
	|=
	^=
	<<=
	>>=
	*/
	a += 1
	fmt.Println(a)
	a = a + 1
	fmt.Println(a)

	/* 증감 연산자 : 증감 연산자는 값을 반환하지 않는다.
	++
	--
	*/
	a++
	fmt.Println(a)

	/* 연산자 우선 순위 : 물론, 괄호((,))안에 식이 가장 먼저 실행
	우선 순위	연산자
	1	*, /, %, <<, >>, &, &^
	2	+, -, |, ^
	3	==, !=, <, <=, >, >=
	4	&&
	5	||
	*/

	/* 실수 비교할 때, 실수의 오차를 고려해서 nextafter를 사용해야함!!
	Nextafter : 전달받은 두 수중에서, 앞에 수를 뒤에 수를 향해 1비트만큼 이동한 값을 반환합니다. 따라서, 정확하게 실수의 값을 비교할 수 있습니다.
	(비트로는 실수를 정확히 표현할 수가 없습니다. 그래서 실수를 표현할 때는, 원래 수보다 1비트 크거나 1비트 작은 근접수로 실수를 표현한다.)

	*/
	var x float64 = 0.1
	var y float64 = 0.2
	var z float64 = 0.3

	fmt.Printf("%0.18f\n", z)
	fmt.Printf("%0.18f\n", x+y)
	fmt.Printf("%0.18f == %0.18f (%v)\n", z, x+y, z == x+y)
	fmt.Printf("%0.18f == %0.18f (%v)\n", z, x+y, z == math.Nextafter(x+y, z))

}

func TestCondition(t *testing.T) {

	/* if/else
	1. if 다음에는 반드시 Boolean식으로 표현한다. (0, 1 이런거 못 씀)
	2. () 안 쓰지만 {}는 무조건 써 줌.
	3. else if 혹은 else 를 쓸 때는 반드시 전 조건의 open braces( { )와 같은 라인에 써준다.
	*/
	k := 2
	if k == 1 {
		println("One")
	} else if k == 2 { // 같은 라인
		println("Two")
	} else { // 같은 라인
		println("Other")
	}

	// if문에서 for문 처럼 Optional Statement를 사용할 수 있다. Conditional Statement와는 세미콜론으로 구분해준다.
	i := 3
	max := 10
	if val := i * 2; val < max {
		println(val)
	}

	// switch/case
	// 1. 하나의 case 에 만족한다면 그 블럭을 실행하고 바로 switch문을 빠져 나온다.
	score := 90
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}

	// 2. 다음 case 블럭들에도 들어가게 하고 싶다면 fallthrough 를 써주면 된다.
	val := 1
	switch val {
	case 1:
		fmt.Println("1 이하")
		fallthrough
	case 2:
		fmt.Println("2 이하")
		fallthrough
	case 3:
		fmt.Println("3 이하")
		fallthrough
	default:
		fmt.Println("default 도달")
	}

	// type 검사
	var v interface{}

	switch v.(type) {
	case int:
		println("int")
	case bool:
		println("bool")
	case string:
		println("string")
	default:
		println("unknown")
	}

}

func TestLoop(t *testing.T) {

	// for : while 없고 for문만 있다.
	for i := 1; i <= 10; i++ {
		println(i)
	}
	n := 98
	for n < 100 { // while문 역할
		println(n)
		n++
	}
	/*
		for { // 무한 반복
		}
	*/
	// pyhton 처럼 for-range 문 사용가능
	names := []string{"kyo", "kk", "latasha"}
	for index, name := range names {
		println(index, name)
	}

	// break<Label> 도 존재.. 필요할때 찾아보기...

}
