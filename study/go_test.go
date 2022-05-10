package study

/**
golang은 테스트 프레임워크를 내장하고 있다.
- 파일명은 _test.go
- 함수명은 Test 로 시작한다.
- 매개변수는 t*testing,T 를 받는다.
- 실패지점에서 t.Fail() 을 호출한다.
*/
import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {

	fmt.Println("test!")
	t.Errorf("error!")

	//result := something()

	//assert.Equal(t, true, result)
	// same: assert.True(t, result)
}

func TestTwo(t *testing.T) {
	t.Log("로그를 보려면 -v 플래그가 필요해요")
}

/* 테스트 코드 실행하기
go test [build/test flags] [packages] [build/test flags & test binary flags]
*/
