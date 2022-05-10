package study

import (
	"fmt"
	"testing"
)

var caffeines map[string]string

/* init
1. package에서 init은 별도의 호출없이 package가 로드되면서 실행된다.
2. package를 import하면서 그 안에 다른 건 필요 없고 init()함수만 실행하고 싶어하면 import하면서 _(under bar)라고 alias를 지정해주면 된다.

*/
func init() {
	caffeines = make(map[string]string)
	caffeines["weak"] = "coffee"
	caffeines["strong"] = "snoopy"
	fmt.Printf("init %v\n", caffeines)
}

/* Scope
함수/구조체/인터페이스/메서드 등등...
1. public : 이름 첫글자가 대문자
2. non-public : 그외 (패키지 내부에서만 사용 가능)
*/
func GetCaffeine(intensity string) string { // 외부에서 호출가능
	return caffeines[intensity]
}

func getAllCaffeines() { // 내부에서만 호출 가능
	//for _, k := range caffeines {
	//	fmt.Println(caffeines[k])
	//}
	for index, name := range caffeines {
		println(index, name)
	}

}

func TestScope(t *testing.T) {
	getAllCaffeines()

}
