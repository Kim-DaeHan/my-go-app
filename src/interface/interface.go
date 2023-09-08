package main

import (
	"fmt"
	"math"
)

// go 인터페이스
// 구조체가 필드들의 집합체라면, interface는 메서드들의 집합체이다
// 인터페이스는 타입이 구현해야 하는 메서드 원형들을 정의

// 인터페이스 정의
type Shape interface {
	area() float64
	perimeter() float64
}

// 인터페이스를 구현하기 위해서는 해당 타입이 그 인터페이스의 메서드들을 모두 구현하면된다.

// Rect 정의
type Rect struct {
	width, height float64
}

// Circle 정의
type Circle struct {
	radius float64
}

// Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// 인터페이스 타입
// 빈 인터페이스를 인터페이스 타입으로 부른다
// Empty interface는 메서드를 전혀 갖지 않는 빈인터페이스로서, go의 모든 type은 적오도 0개의 메서드를 구현하므로
// 흔히 go 에서 모든 type을 나타내기 위해 빈 인터페이스를 사용한다.
// 즉 빈 인터페이스는 어떠한 타입도 담을 수 있는 컨테이너라고 볼 수 있으며, 여러 다른 언어에서 흔히 일컫는 Dynamic Type(java의 Object)이라고 볼 수 있다.
// func Marshal(v interface{}) ([]byte, error)
// func Println(a ...interface{}) (n int, err error)

// Type Assertion
// x.(T)로 표현했을 때, 이는 x가 nil이 아니며, x는 T타입에 속한다는 점을 확인(assert)하는 것
// 만약 x가 nil 이거나 x의 타입이 T가 아니라면, 런타임 에러가 발생할 것이고 x가 T 타입인 경우는 T 타입의 x를 리턴한다

func main() {
	// 인터페이스 사용
	r := Rect{10, 20}
	c := Circle{10}

	showArea(r, c)

	var x interface{}
	x = 1
	x = "Tom"

	printIt(x)

	var a interface{} = 1

	i := a       // a와 i 는 dynamic type, 값은 1
	j := a.(int) // j는 int 타입, 값은 1

	printIt(i) // 포인터 주소 출력
	printIt(j) // 1 출력
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area() // 인터페이스 메서드 호출
		fmt.Println(a)
	}
}

func printIt(v interface{}) {
	fmt.Println(v) // Tome
}
