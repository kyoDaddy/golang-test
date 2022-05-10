package study

import (
	"fmt"
	"testing"
	"time"
)

/* gochannel
1. 채널을 통하여 데이터를 주고 받는 통로
2. make() 함수를 통해 미리 생성되어야 하며, 채널 연산자 <- 을 통해서 데이터를 보내고 받는다.
3. 채널은 흔히 goroutine 들 사이 데이터를 주고 받는데 사용되는데, 상대편이 준비될 때까지 채널에서 대기함으로써 별도의 lock을 걸지 않고 데이터를 동기화화는데 사용된다.

*/

func TestGoChannel(t *testing.T) {
	// 정수형 채널을 생성한다
	// 채널을 생성할 때는 make() 함수에 어떤 타입의 데이타를 채널에서 주고 받을지를 미리 지정해 주어야 한다.
	ch := make(chan int)

	go func() {
		ch <- 123 // 채널에 123을 보낸다 (채널로 데이타를 보낼 때는 채널명 <- 데이타 와 같이 사용)
	}()

	var i int
	i = <-ch // 채널로부터 123을 받는다 (채널로부터 데이타틀 받을 경우는 <- 채널명 와 같이 사용)
	println(i)
}

func TestGoChannel2(t *testing.T) {
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	// 위의 goroutine 끝날 때까지 대기
	<-done
}

/* gochannel-buffering
1. unbuffered channel : 하나의 수신자가 데이터를 받을 때까지 송신자가 데이터를 보내는 채널에 묶여 있게 된다.
2. buffered channel : 수신자가 받을 준비가 되어 있지 않을 지라도, 지정된 버퍼만큼 데이터를 보내고 계속 다른 일을 수행할 수 있다.
2-1) 버퍼 채널은 make(chan type, N) 함수를 통해 생성, 두번째 파라미터 N에 사용할 버퍼 갯수를 넣는다.
*/
func TestChannelBuffering(t *testing.T) {
	c := make(chan int)
	c <- 1           // 수신루틴이 없으므로 데드락
	fmt.Println(<-c) // 코멘트해도 데드락 (별도의 goroutine이 없기 떄문)
}

func TestChannelBuffering1(t *testing.T) {
	ch := make(chan int, 1)

	// 수신자가 없더라도 보낼 수 있다.
	// 버퍼채널을 사용하면, 수신자가 당장 없더라도 최대버퍼 수까지 데이터를 보낼 수 있으므로, 에러가 발생하지 않는다.
	ch <- 101

	fmt.Println(<-ch)
}

/* channel parameter
1. 송수신을 모두 하는 채널을 일반적으로 전달
2. 해당 채널로 송신만 할것인지 수신만할 것인지 지정가능
2-1) 송신 파라미터 (p chan<- int)
2-2) 수신 파라미터 (p <- chan int)
*/
func TestChannelParameter(t *testing.T) {
	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch)
}

func sendChan(ch chan<- string) {
	ch <- "Data"
	// x := <-ch // 에러발생
}
func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

/* channel close
1. 채널 오픈 후 데이터 송신한 후 , close() 함수로 채널을 닫을 수 있다.
2. 채널 닫게 되면, 해당 채널로는 더이상 송신을 할 수 없지만, 닫힌 이후에도 계속 수신은 가능하다.
*/
func TestChannelClose(t *testing.T) {
	ch := make(chan int, 2)

	// 채널에 송신
	ch <- 1
	ch <- 2

	// 채널을 닫는다
	close(ch)

	// 채널 수신
	println(<-ch)
	println(<-ch)

	if _, success := <-ch; !success {
		println("더이상 데이터 없음")
	}
}

/* channel range
수신자는 임의의 갯수의 데이타를 채널이 닫힐 때까지 계속 수신할 수 있다.
*/
func TestChannelRange(t *testing.T) {
	ch := make(chan int, 2)

	// 채널에 송신
	ch <- 1
	ch <- 2

	// 채널을 닫는다
	close(ch)

	// 방법1
	// 채널이 닫힌 것을 감지할 때까지 계속 수신
	//for {
	//	if i, success := <-ch; success {
	//		println(i)
	//	} else {
	//		break
	//	}
	//}

	// 방법2
	// 방법1과 동일한 채널 range 문
	for i := range ch {
		println(i)
	}
}

/* channel select

 */
func TestChannelSelect(t *testing.T) {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 완료")

		case <-done2:
			println("run2 완료")
			break EXIT
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}
func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
