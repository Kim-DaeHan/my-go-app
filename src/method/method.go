package main

import "fmt"

// go method
// go 언어는 객체지향 프로그래밍 고유의 방식으로 지원한다. 그래서 필드와 메서드를 함께 가질 수 없고 go 언어에서 struct가 필드만을 가지며 메서드는 별도로 분리되어 정의된다.
// go 메서드는 특별한 형태의 func 함수이다. 메서드는 함수 정의에서 func 키워드와 함수명 사이에 어떤 struct를 위한 메서드인지를 표시하게 된다. 흔히 receiver로 불리우는
// 이부분은 메서드가 속한 struct 타입과 struct 변수명을 지정하는데, struct 변수명은 함수 내에서 마치 입력 파라미터처럼 사용된다

// Rect - struct 정의
type Rect struct {
	width, height int
}

// Rect의 area() 메서드
func (r Rect) area() int {
	r.width++
	return r.width * r.height
}

// Value vs 포인터 receiver
// Value receiver는 struct의 데이터를 복사하여 전달하며, 포인터 receiver는 struct의 포인터만을 전달한다.
// Value receiver의 경우 만약 메서드내에서 그 struct의 필드값이 변경되더라도 호출자의 데이터는 변경되지 않는 반면
// 포인터 receiver는 메서드 내의 필드값 변경이 그대로 호출자에서 반영된다
// 위의 Rect.area() 메서드는 value receiver를 표현한 것으로 만약 area() 메서드 내에서 width나 height가 변경되더라도 main() 함수의 rect 구조체의 필드값에는 변화가 없다.
// main 함수 의 rect 구조체의 필드값에는 변화가 없지만 메서드 내에서는 변화한 값을 기반으로 계산되어 리턴된다.

// 포인터 Receiver
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()           // 메서드 호출
	fmt.Println(rect.width, area) // 10, 220

	rect2 := Rect{10, 20}
	area2 := rect2.area2()          // 메서드 호출
	fmt.Println(rect2.width, area2) // 11 220

}
