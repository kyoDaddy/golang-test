package main

import (
	"fmt"
	"golang-test/testlib"
)

// terminal > go run hello.go
func main() {
	// GOROOT 검색 후 GOPATH 검색하여 겨로를 사용
	song := testlib.GetMusic("Alicia Keys")
	fmt.Println(song)
	//fmt.Println(quote.Go())
}
