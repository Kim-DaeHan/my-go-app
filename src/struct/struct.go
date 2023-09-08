package main

import "fmt"

// struct(구조체)
// go에서 struct는 Custom Data Type을 표현하는데 사용되는데 go의 struct는 필드들의 집합체이며 필드들의 컨테이너이다.
// go에서 struct는 필드 데이터만을 가지며, (행위를 표현하는) 메서드를 갖지 않는다.
// go언어는 객체지향 프로그래밍을 고유의 방식으로 지원한다. 즉 go에는 전통적인 OOP 언어가 가지는 클래스, 객체, 상속 개념이 없다.
// 전통적인 OOP의 클래스가 필드와 메서드를 함께 갖는 것과 달리 go 언어의 struct는 필드만을 가지며, 메서드는 별도로 분리하여 정의한다.

// struct 선언(public으로 선언할려면 Person)
type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

// 생성자 함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d // 포인터 전달
}

func main() {
	// person 객체 생성(인스턴스 생성)
	p := person{}

	// 필드값 설정
	p.name = "Lee"
	p.age = 10

	fmt.Println(p)

	// person 인스턴스 초기값 할당
	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}
	p3 := person{"Kim", 30}
	fmt.Println(p1, p2, p3)

	// new() 사용 : 모든 필드를 Zero value로 초기화하고 person 객체의 포인터(*person)를 리턴한다.
	p4 := new(person)
	p4.name = "Hans" // p4가 포인터라도 . 을 사용한다.
	fmt.Println(p4)

	// go에서 struct는 기본적으로 mutable 개체로서 필드값이 변화할 경우(별도로 새 개체를 만들지 않고) 해당 개체 메모리에서 직접 변경된다.
	// 하지만 struct 개체를 다른 함수의 파라미터로 넘긴다면, Pass by Value에 따라 객체를 복사해서 전달하게 된다.
	// 그래서 만약 Pass by Reference로 struct를 전달하고자 한다면, struct의 포인터를 전달해야한다.

	// 생성자 함수
	// 생성자 함수는 struct를 리턴하는 함수로서 그 함수 본문에서 필요한 필드를 초기화한다.
	dic := newDict() // 생성자 호출
	dic.data[1] = "A"
	fmt.Println(dic)
}
