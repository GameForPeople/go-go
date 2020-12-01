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
//// Go에는 Enum이 존재하지 않고, 상수로 이를 모두 처리한다.
const (
	cb   bool    = false // 상수는 zeroValue를 허락하지 않는다
	ci8  int8    = 1
	cf32 float32 = 10.5
	cs   string  = "ABCD" // 스트링도 상수가 된다;;
)

//// 구조체의 선언 ================================================================
type Position struct {
	X int
	Y int
}

//// 타입의 선언 ==================================================================
type Nickname string
type Level uint16

// C++ 과 다른 점은 , using 혹은 typedef로 정의할 경우, 단순히 동등 - 치환으로 생각하는데,
// Golang의 Type은 말 그대로 다른 Type으로 서로 동등하지 않다.

func main() {

	fmt.Println()

	// 기본 출력문은 다음과 같이 처리한다. =================================================================
	{
		fmt.Println("Hello Go!")
		fmt.Println("The time is", time.Now())
	}

	fmt.Println()

	// 배열의 선언은 다음과 같이 처리합니다.
	{
		var temp [2]int
		temp[0] = 100
		temp[1] = 150
		fmt.Println("temp:", temp)

		var temp2 = [3]string{"ABC", "XYZ", "WSY"}
		fmt.Println("temp:", temp2)

		temp2ptr := &temp2[2]
		*temp2ptr = "NEO"
		fmt.Println("temp:", temp2)
	}

	fmt.Println()

	// C++ 에는 없는 슬라이스...? 가변 길이의 배열이라는데 벡터인 것인가 리스트인 것인가는 뭔가 배열의 레퍼런스인가...?
	{
		// 기존의 배열을 레퍼런스 화한다.
		{
			temp := [3]int{10, 20, 30}
			tempSlice := temp[0:3]
			// 0번쨰 인덱스부터 3보다 작은 인덱스 2까지로 슬라이스를 만든다.
			// temp [ 0:3] == temp [ 0:] == temp [:3] == temp[:]

			fmt.Println("tempSlice : ", tempSlice)
			tempSlice[0] = 100

			fmt.Println("tempSlice change 'temp' : ", temp)
			tempSliceSlice := tempSlice[2:3]
			tempSliceSlice[0] = 1000

			fmt.Println("tempSliceSlice change 'temp' : ", temp)
		}

		// 슬라이스 자체를 바로 만든다.
		{
			tempSlice := []int{1, 2, 3, 4, 5}
			fmt.Println("tempSlice: ", tempSlice)

			// 슬라이스를 만드는데 세상에 구조체를 바로만들어서 슬라이스를 만든다...?
			tempStructSlice := []struct {
				name string
				age  uint8
			}{
				{"", 1},
				{"", 2}
			}
		}
	}

	fmt.Println()

	// for문은 다음과 같이 합니다. (엥 괄호가 없네) ( go에는 while문을 지원하지 않습니다. ) =====================
	{
		for i := 0; i < 3; i++ {
			fmt.Println("for Loop : ", i)
		}

		//범위 기반 for문은 다음과 같이 돌립니다.
		var temp = [2]int{10, 20}
		for i := range temp {
			fmt.Println("for Loop Range : ", temp[i])
		}
	}

	fmt.Println()

	// if문은 다음과 같이 합니다. 여기도 괄호가 없지만, C++17처럼 If-Init이 가능함 =====================
	{
		if var1 := 1; var1 == 0 {
			fmt.Println("if var == 0 : ", var1)
		} else {
			fmt.Println("for var != 0 : ", var1)
		}
	}

	fmt.Println()

	// switch문은 기존과 다르게 별도의 break없이도 멈추며, fallthrough를 명시하지 않으면 해당 Case만 실행됩니다.
	{
		temp := 0
		switch temp {
		case 0:
			temp++
		case 1:
			temp++
		case 2:
			temp++
		}
		fmt.Println("switch not fallthrough : ", temp)

		temp = 0

		switch temp {
		case 0:
			temp++
			fallthrough
		case 1:
			temp++
			fallthrough
		case 2:
			temp++
			fallthrough
		default:
		}
		fmt.Println("swith with fallthrough : ", temp)
	}

	fmt.Println()

	// omg! golang pointer! ====================================================================
	{
		temp1 := 10
		temp2 := 20

		// pointer var
		pTemp := &temp1

		fmt.Println("ptemp :", pTemp, ", &pTemp :", &pTemp, ", *pTemp :", *pTemp)
		fmt.Println("temp1 :", temp1, ", &temp1 :", &temp1)

		//  change value to p-Var
		*pTemp = 100

		fmt.Println("ptemp :", pTemp, ", &pTemp :", &pTemp, ", *pTemp :", *pTemp)
		fmt.Println("temp1 :", temp1, ", &temp1 :", &temp1)

		//
		pTemp = &temp2

		fmt.Println("ptemp :", pTemp, ", &pTemp :", &pTemp, ", *pTemp :", *pTemp)
		fmt.Println("temp2 :", temp2, ", &temp2 :", &temp2)
	}

	//
	{

	}
}
