package study

import (
	"fmt"
	"os"
	"testing"
)

/* 지연실행 : defer
1. 특정 문장 혹은 함수를 나중에(defer 호출하는 함수가 리턴하기 직전에) 실행하게 한다.
2. javs에서 finally 블록처럼 마지막에 clean-up 작업을 위해 사용
3. 차후 문장에서 어떤 에러가 발생하더라도 항상 파일을 Close 할 수 있도록 한다.

*/
func TestDefer(t *testing.T) {
	f, err := os.Open("1.txt")
	if err != nil {
		panic(err)
	}

	// main 마지막에 파일 close 실행
	defer f.Close()

	// 파일 읽기
	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))
}

/* panic
1. 현재 함수를 즉시 멈추고 현재 함수에 defer 함수들을 모두 실행한 후 즉시 리턴한다.
2. 이러한 panic 모드 실행 방식은 다시 상위함수에도 똑같이 적용되고, 계속 콜스택을 타고 올라가며 적용된다. 그리고 마지막에 프로그램 에러를 내고 종료
*/
func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func TestPanic(t *testing.T) {
	// 잘못된 파일명을 넣음
	openFile("Invalid.txt")

	// openFile() 안에서 panic이 실행되면 아래 println 문장은 실행 안됨
	println("Done")

}

/* recover
1. panic 함수에 의한 패닉상태를 다시 정상상태로 되돌리는 함수
2. panic -> defer (recover()) -> panic 상태 제거하고 다음 문장 호출
*/
func openFile2(fn string) {
	// defer 함수, panic 호출시 실행됨
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func TestRecover(t *testing.T) {
	// 잘못된 파일명을 넣음
	openFile2("Invalid.txt")

	// openFile() 안에서 panic이 실행되면 아래 println 문장은 실행 안됨
	println("Done")

}
