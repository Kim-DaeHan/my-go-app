package main

import "fmt"

func main() {
	msg := "Hello"
	say(msg)
	fmt.Println(msg) // 변경되지 않은 메시지 출력
	say2(&msg)
	fmt.Println(msg) // 변경된 메시지 출력
	say3("This", "is", "a", "book")
	say3("Hi")
	total := sum(1, 7, 3, 5, 9)
	fmt.Println(total)
	count, total := sum2(1, 7, 3, 5, 9)
	fmt.Println(count, total)
	count, total = sum3(1, 7, 3, 5, 9)
	fmt.Println(count, total)
}

func say(msg string) {
	fmt.Println(msg)
	msg = "ddd" // 메시지 변경
}

// 01. Pass By Value : 위의 예제에서는 msg의 값 "Hello" 문자열이 복사되어 함수 say()에 전달된다.
// 즉 만약 파라미터 msg의 값이 say() 함수 내에서 변경된다하더라도 호출함수 main()에서의 msg변수는 변함이 없다.
// 02. Pass By Reference : 아래의 예제에서처럼 msg 변수앞에 & 부호를 붙이면 msg 변수의 주소를 표시하게 된다.
// 흔히 포인터라 불리우는 이 용법을 사용하면 함수에 msg 변수의 값을 복사하지 않고 msg 변수의 주소를 전달하게 된다.
// 피호출 함수 say2() 에서는 *string과 같이 파라미터가 포인터임을 표시하고 이 때 say2 함수의 msg는 문자열이 아니라 문자열을 갖는 메모리 영역의 주소를 갖게 된다.
// msg 주소에 데이터를 쓰기 위해서는 *msg = ""과 같이 앞에 *를 붙이는데 이를 흔히 Dereferencing이라 한다.

func say2(msg *string) {
	fmt.Println(msg)
	*msg = "Changed" // 메시지 변경
}

// 03. 가변인자함수 : 함수에 고정된 수의 파라미터들을 전달하지 않고 다양한 숫자의 파라미터를 전달하고자 할 때 가변 파라미터를 나타내는 ... 을 사용
// 즉 문자열 가변 파라미터를 나타내기 위해서 ...string과 같이 표현
func say3(msg ...string) {
	for _, s := range msg {
		fmt.Println(s)
	}
}

// 04. 함수 리턴값 : 리턴값이 없을 수도, 리턴값이 하나 일 수도, 또 는 복수 개일 수도 있다.
// Named Return Parameter라는 기능을 제공하는데, 이는 리턴되는 값들을 (함수에 정의된) 리턴 파라미터들에 할당할 수 있는 기능
func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func sum2(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func sum3(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
