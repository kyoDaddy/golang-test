package study

import (
	"fmt"
	"testing"
)

/* Interface Type
1. 빈 interface 이며, interface{}와 같이 표현
2. method를 전혀 갖지 않는데, go 모든 type은 적어도 0개의 메서드를 구현하므로, 흔히 Go에서 모든 type을 나타내기 위해 빈 interface는 어떠한 타입도 담을 수 있는 컨테이너라고 볼 수 있다.
3. empty interface -> java에서 object

*/
func printIt(v interface{}) {
	fmt.Println(v)
}

func TestInterfaceType(t *testing.T) {
	var x interface{}
	x = 1
	x = "Tom"

	fmt.Println(x)
}
