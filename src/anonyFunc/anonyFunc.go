package main

import "fmt"

func main() {
	// 익명함수: 함수명을 갖지 않는 함수
	sum := func(n ...int) int { // 익명함수 정의
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5) // 익명함수 호출
	fmt.Println(result)

	// 변수 add에 익명함수 할당
	add := func(i int, j int) int {
		return i + j
	}

	// add 함수 전달
	r1 := calc(add, 10, 20)
	fmt.Println(r1)

	// 직접 첫번째 파라미터에 익명함수를 정의함
	r2 := calc2(func(x int, y int) int { return x - y }, 10, 20)
	fmt.Println(r2)
}

// 일급 함수
// 함수의 입력 파라미터나 리턴 파라미터로서 함수 자체가 사용 될 수 있다
func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

// type문을 사용한 함수 원형 정의
//  type문은 구조체, 인터페이스 등 Custom Type을 정의하기 위해 사용
// 또 한 함수 원형을 정의하는데 사용될 수 있다.
// 즉 위 예제에서 함수 원형이 계속 반복됨을 볼 수 있는데, 이 경우 type 문을 정의함으로써 해당 함수의 원형을 간단히 표현 가능하다

// 원형 정의
type calculator func(int, int) int

// 원형 사용
func calc2(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}
