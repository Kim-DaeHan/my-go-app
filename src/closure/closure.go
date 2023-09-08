// go 언어에서 함수는 클로저로서 사용 가능하다
// 클로저는 함수 바깥에 있는 변수를 참조하는 함수 값을 일컫는데, 이때 함수는 바깥의 변수를 마치 함수 안으로 끌어들인 듯이 그 변수를 읽거나 쓸 수 있게 된다.
// 아래 예제에서 nextValue() 함수는 int를 리턴하는 익명함수를 리턴하는 함수이다. go 언어에서 함수는 일급함수로서 다른 함수로부터 리턴되는 리턴값으로 사용될 수 있다.
// 그런데 여기서 이 익명함수가 그 함수 바깥에 있는 변수 i를 참조하고있다. 익명함수 자체가 로컬 변수로 i를 갖는 것이 아니기 때문에 외부 변수 i가 상태를 계속 유지하는 즉 값을 계속 하나씩 증가시키는 기능을 하게된다.
package main

import "fmt"

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := nextValue()

	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3

	anotherNext := nextValue()
	fmt.Println(anotherNext()) // 1 다시 시작
	fmt.Println(anotherNext()) // 2
}
