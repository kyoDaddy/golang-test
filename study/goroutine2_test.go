package study

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine2(t *testing.T) {
	// goroutines는 프로그램이 작동하는 동안만 유효하다.
	// 메인함수가 실행되는 동안만..메인 함수는 goroutine을 기다려주지 않는다.
	c := make(chan string)
	people := [5]string{"kyo", "latasha", "unkwon1", "unkwon2", "unkwon3"}
	for _, person := range people {
		go isSexy(person, c)
	}
	// 채널의 룰은 매우 심플하다.. 먼저 해야하는 걸 먼저 해치운다.
	for i := 0; i < len(people); i++ {
		fmt.Println("Waiting for ", i)
		fmt.Println(<-c)
	}

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}
