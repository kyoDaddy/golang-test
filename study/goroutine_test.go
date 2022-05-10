package study

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

/* goroutine
1. go 런타임이 관리하는 lightweight 논리적(가상적) 쓰레드
2. go 키워드를 사용하여 함수를 호출하면, 런타임시 새로운 goroutine을 실행 > goroutine은 비동기적으로 함수루틴을 실행하므로, 여러코드를 동시헤 실행하는데 사용된다.
*/
func say1(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func TestGoroutine(t *testing.T) {
	// 함수를 동기적으로 실행
	// 첫번째 동기적 호출은 say() 함수가 완전히 끝났을 때 다음 문장으로 이동
	say1("Sync")

	// 함수를 비동기적으로 실행
	// 다음 3개의 go say() 비동기 호출은 별도의 Go루틴들에서 동작하면서, 메인루틴은 계속 다음 문장(여기서는 time.Sleep)을 실행
	// 여기서 goroutine들은 그 실행순서가 일정하지 않으므로 프로그램 실행시 마다 다른 출력 결과를 나타낼 수 있다.
	go say1("Async1")
	go say1("Async2")
	go say1("Async3")

	// 3초 대기
	time.Sleep(time.Second * 3)

}

/* anonymous function - goroutine
1. go 키워드 뒤에 익명함수를 바로 정의하는 것으로, 익명함수를 비동기로 실행하게 됨
2. sync.WaitGroup을 사용하고 있는데, 이는 기본적으로 여러 Go루틴들이 끝날 때까지 기다리는 역활을 한다.
2-1) WaitGroup을 사용하기 위해서는 먼저 Add() 메소드에 몇 개의 Go루틴을 기다릴 것인지 지정하고, 각 Go루틴에서 Done() 메서드를 호출한다 (여기서는 defer 를 사용하였다)
2-2) 메인루틴에서는 Wait() 메서드를 호출하여, Go루틴들이 모두 끝나기를 기다린다.
*/
func TestAnonymousGoroutine(t *testing.T) {
	// WaitGroup 생성. 2개의 Go루틴을 기다림.
	var wait sync.WaitGroup
	wait.Add(2)

	// 익명함수를 사용한 goroutine
	go func() {
		defer wait.Done() // 끝나면 .Done() 호출
		fmt.Println("Hello")
	}()

	// 익명함수에 파라미터 전달
	go func(msg string) {
		defer wait.Done() // 끝나면 .Done() 호출
		fmt.Println(msg)
	}("Hi")

	wait.Wait() // Go루틴 모두 끝날 때까지 대기
}

/* 다중 CPU 처리
1. go는 디폴트로 1개의 cpu를 사용한다.
2. 즉 여러개의 goroutine을 만들더라도, 1개의 cpu에서 작업을 시분할하여 처리한다(concurrent 처리)
3. 머신이 복수개의 cpu를 가진 경우, go 프로그램을 다중 cpu에서 병렬처리(parallel 처리)하게 할 수 있는데, runtime.GOMAXPROCS(CPU수) 함수를 호춣해야한다. (여기서 cpu수는 logical cpu 수)
4. 동시성, 병렬화
4-1) 프로그래밍에서 동시성은 독립적으로 실행되는 프로세스의 구성이다. 동시성은 한 번에 많은 것을 처리하는 것입니다.
4-2) 병렬화(parallelism)는 병렬화(parallelision) 계산의 동시 실행이다. 병렬화는 동시에 많은 것을 하는 것입니다.
*/

func TestMultiCpu(t *testing.T) {
	// 	4개의 cpu 사용
	runtime.GOMAXPROCS(4)

	names := []string{"kyo", "kk", "latasha", "ain", "h", "m", "k", "j"}

	var wait sync.WaitGroup
	wait.Add(len(names))

	// 함수를 동기적으로 실행
	// 첫번째 동기적 호출은 say() 함수가 완전히 끝났을 때 다음 문장으로 이동
	say1("Sync")

	// 함수를 비동기적으로 실행
	for _, name := range names {
		go func(msg string) {
			defer wait.Done() // 끝나면 .Done() 호출
			say1("Async_" + msg)
		}(name)
	}

	wait.Wait() // Go루틴 모두 끝날 때까지 대기

}
