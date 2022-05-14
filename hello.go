// 패키지명을 꼭 붙어야함!
package main

import (
	"fmt"
	"golang-test/something"
	"golang-test/testlib"
)

// main package와 그 안에 있는 main function을 먼저 찾고 실행시킴
// main은 컴파을 위해 필요한것!
// terminal > go run hello.go
func main() {
	// GOROOT 검색 후 GOPATH 검색하여 겨로를 사용
	song := testlib.GetMusic("Alicia Keys")
	fmt.Println(song)
	//fmt.Println(quote.Go())
	something.SayHello()

}
