package main

import "fmt"

func main() {

	// 슬라이스는 배열과 달리 고정된 크기를 미리 지정하지 않을 수 있고, 차후 그 크기를 동적으로 변경할 수 도 있고, 또한 부분 배열을 발췌할 수 있음
	// 배열 선언하듯이 var v[]T처럼 하는데 크기는 지정하지 않는다.

	var a []int // 슬라이스 변수 선언
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	// make함수 사용 : 슬라이스의 길이와 용량을 지정할 수 있다
	// 첫번째 파라미터에 타입 지정, 두번째는 길이, 세번째는 내부 배열의 최대 길이
	// 만약 세번째 Capacity(내부 배열의 최대 길이) 생략하면 두번째 길이와 같은 값을 갖는다

	s := make([]int, 5, 10)
	fmt.Println(len(s), cap(s))

	// 별도의 길이와 용량을 지정하지 않으면 길이와 용량이 0인 슬라이스를 만들고 이를 Nil Slice라고 하고, nil 과 비교하면 참을 리턴
	var s2 []int
	if s2 == nil {
		fmt.Println("nil slice")
	}

	fmt.Println(len(s2), cap(s2))

	// 부분 슬라이스
	s3 := []int{0, 1, 2, 3, 4, 5}
	s3 = s3[2:5]
	fmt.Println(s3)

	// 슬라이스 추가, 병합(append)과 복사(copy)
	s4 := []int{0, 1}

	// 하나 확장
	s4 = append(s4, 2)
	// 복수 요소들 확장
	s4 = append(s4, 3, 4, 5)
	fmt.Println(s4)

	// len=0 cap=3 인 슬라이스
	sliceA := make([]int, 0, 3)

	// 계속 한 요소씩 추가
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		// 슬라이스 길이와 용량 확인
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA) // 1 부터 15 까지 숫자 출력

	sliceB := []int{1, 2, 3}
	sliceC := []int{4, 5, 6}

	sliceB = append(sliceB, sliceC...)

	fmt.Println(sliceB)

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source) // target에 source 내용 복사
	fmt.Println(target)
	fmt.Println(len(target), cap(target))
}
