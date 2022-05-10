package testlib

import "fmt"

/** 사용자 정의 패키지 생성
1. 사용자 정의 라이브러리 패키지는 일반적으로 폴더를 하나 만들고 그 폴더 안에 .go 파일들을 만들어 구성한다.
2. 하나의 서브 폴더안에 있는 .go 파일들은 동일한 패키지명을 가지며, 패키지명은 해당 폴더의 이름과 같게 한다.
3. 해당 폴더에 있는 여러 *.go 파일들은 하나의 패키지로 묶인다.
4. go install : 라이브러를 컴파일 후 cache할 수 있다.
*/
var pop map[string]string

func init() {
	pop = make(map[string]string)
	pop["Adele"] = "Hello"
	pop["Alicia Keys"] = "Falling"
	pop["John Legend"] = "All of Me"
}

// GetMusic : Popular music by singer (외부에서 호출 가능)
func GetMusic(singer string) string {
	return pop[singer]
}

func getKeys() { // 내부에서만 호출 가능
	for _, kv := range pop {
		fmt.Println(kv)
	}
}
