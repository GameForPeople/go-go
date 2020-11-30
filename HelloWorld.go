// 해당 문서는, Go언어의 전체적인 문법을 정리하여 추후 헷갈릴때마다 확인하기 위함입니다.

package main

//// 임포트 ============================================================================
//0. 다음과 같이 하거나
//import "fmt"
//import "time"

// 1. 아래와 같이한다. <<---- 선택
import (
	"fmt"
	"time"
)

//// 함수 ( 인자 정의 ) ==================================================================
// 0. 함수는 다음과 같이 선언 - 정의한다. <<---- 선택
func func1(x int, y int) int {
	return x + y
}

// 1. 인자가 동일할 경우에는, 생략할 수 있다고하는데 안쓰는게 가독성에 더 좋을 것 같다.
func func2(x, y int) int {
	return x + y
}

//// 함수 ( 반환값 정의 ) ==================================================================
// 0. 반환형을 다음과 같이 노멀하게 정의할 수도 있고
func func3(x int, y int) (int, int) {
	var a = x + y
	var b = x - y

	return a, b
}

// 1. 변수명까지 정의할 수 있다. 이 때에는 함수 정의부에서 추가 선언이 필요 없다. <<--- 선택
func func4(x int, y int) (a int, b int) {
	a = x + y
	b = x - y
	return
}

//// 변수의 선언 ==================================================================
var (
	b    bool // == zero value is false
	by   byte
	i8   int8 // == zero value is 0
	i8_2 int8 = 1
	ui8  uint8
	i16  int16
	ui16 uint16
	i32  int32
	ui32 uint32
	f32  float32
	f64  float64
	s    string    // == zero value is ""
	c64  complex64 // 복소수라니;;;
)

//// 상수의 선언 =================================================================
const (
	cb   bool    = false // 상수는 zeroValue를 허락하지 않는다
	ci8  int8    = 1
	cf32 float32 = 10.5
	cs   string  = "ABCD" // 스트링도 상수가 된다;;
)

func main() {

	// 기본 출력문은 다음과 같이 처리한다. =================================================================
	{
		fmt.Println("Hello Go!")
		fmt.Println("The time is", time.Now())
	}

	// for문은 다음과 같이 합니다. (엥 괄호가 없네) ( go에는 while문을 지원하지 않습니다. ) =====================
	{
		for i := 0; i < 3; i++ {
			fmt.Println("for Loop : ", i)
		}
	}

	// if문은 다음과 같이 합니다. 여기도 괄호가 없지만, C++17처럼 If-Init이 가능함
	{
		if var1 := 1; var1 == 0 {
			fmt.Println("if var == 0 : ", var1)
		} else {
			fmt.Println("for var != 0 : ", var1)
		}
	}

	// switch문은 기존과 다르게 fallthrough
}
