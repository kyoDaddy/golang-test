package study

import (
	"log"
	"os"
	"testing"
)

func TestError(t *testing.T) {
	f, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	println(f.Name())

	/* case2 : error의 Type을 체크해서 에러 타입별로 별도의 에러 처리를 하는 방식
	switch err.(type) {
	default:
		println("ok")
	case MyError:
		log.Print("Log my error")
	case error:
		log.Fatal(err.Error())
	}
	*/

}
