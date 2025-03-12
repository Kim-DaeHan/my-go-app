package main

import "fmt"

// 두 인터페이스 정의
type Walker interface {
	Walk()
}

type Talker interface {
	Talk()
}

// 함수는 Walker만 요구
func TakeAStroll(w Walker) {
	// 기본 기능 수행
	w.Walk()

	// 추가 기능이 있는지 확인
	if talker, ok := w.(Talker); ok {
		// 추가 기능이 있다면 사용
		talker.Talk()
	} else {
		fmt.Println("이 객체는 말할 수 없습니다")
	}
}

// 두 가지 구현체
type Dog struct{}

func (d Dog) Walk() { fmt.Println("개가 걷습니다") }
func (d Dog) Talk() { fmt.Println("멍멍!") }

type Fish struct{}

func (f Fish) Walk() { fmt.Println("물고기가 헤엄칩니다") }

// Fish는 Talk를 구현하지 않음

func main() {
	TakeAStroll(Dog{})  // "개가 걷습니다" "멍멍!"
	TakeAStroll(Fish{}) // "물고기가 헤엄칩니다" "이 객체는 말할 수 없습니다"
}
