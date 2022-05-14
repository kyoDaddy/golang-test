package study

import (
	"fmt"
	_ "fmt" // 경우에 따라 패키지를 import 하면서 단지 그 패키지 안의 init() 함수만을 호출하고자 하는 케이스가 있다. 이런 경우는 패키지 import 시 _ 라는 alias 를 지정한다
	"testing"
)

func TestCollction(t *testing.T) {

}

func TestArray(t *testing.T) {
	// zero base
	// 선언은 var <변수 이름> [배열크기]<type>
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3

	var b = [3]int{1, 2, 3}
	var c = [...]int{1, 2, 3} // 배열크기 자동으로

	var d = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, // 끝에 콤마 추가
	}

	fmt.Printf("%v\r\n", a)
	fmt.Printf("%v\r\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)

}

func TestSlice(t *testing.T) {
	// 선언은 배열과 똑같이 하는데, 대신 [] 안에 크기는 지정해주지 않는다.
	// var <변수 이름> <type>

	// var s []int
	s := []int{0, 1, 2, 3, 4, 5}
	var a = []int{1, 2, 3}

	if s == nil {
		println("Nil Slice")
	}

	// slice with len :5, capacity : 10
	a = make([]int, 5, 10)

	// 파이썬처럼 index를 활용해 sub slice를 만들어 줄 수 있다.
	// slice[start idx : end idx +1 ]
	s = s[2:5]
	fmt.Printf("%v\n", s)
	s = s[1:]
	fmt.Printf("%v\n", s)

	println("a->", len(a), cap(a))

	// slice와 slice는 append()를 사용해 연결할 수 있다.
	// 주의할 점은 붙일 slice 중 두번째에 ellipsis(...)를 붙여야 한다는 것.
	// Go에서의 ellipsis는 여러 의미로 사용되는데, 여기서는 "unpacking" 의 의미이다. 해당 slice의 컬렉션을 표현하는 것( slice의 모든 element들의 집합)이래서 뭔가 했는데, 대충 slice에 있는 element들을 unpack해서 append같은 variadic function에 넣어준다는 것 같음.
	sliceA := []int{1, 2, 3}
	sliceA = append(sliceA, 10)
	sliceB := []int{4, 5, 6}
	sliceC := append(sliceA, sliceB...)
	fmt.Println("sliceC", sliceC)

	// copy()를 활용하여 복사할 수도 있다.
	sliceD := make([]int, len(sliceC), cap(sliceC)*2)
	copy(sliceD, sliceC)
	fmt.Printf("sliceD %v\n", sliceD)

}

func TestMap(t *testing.T) {
	/* hashtable
	- map[key type]value type
	1. var idmap map[int]string
	2. idmap = make(map[int]string)
	*/

	tickers := map[string]string{
		"GOOD": "Gooogle Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
		"AMZN": "Amazon",
	}

	// map 키 체크
	val, exists := tickers["MSFT"]
	if !exists {
		println("No MSFT ticker")
	}
	// Map은 unordered인 hash이므로 순서가 무작위, 따라서 for range를 이용했을 때 모든 매 요소가 매번 다른 순서로 반복문에 돌려진다.
	for index, name := range tickers {
		println(index, name)
	}

	println(val, exists)

	kyoMap := map[string]string{"name": "kyo", "age": "12"}
	fmt.Println(kyoMap)

}
