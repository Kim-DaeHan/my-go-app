package main

import "fmt"

func main() {

	k := 1

	if k == 1 {
		println("One")
	} else if k == 2 { //같은 라인
		println("Two")
	} else { //같은 라인
		println("Other")
	}

	i := 2
	max := 5
	if val := i * 2; val < max {
		println(val)
	}

	// 아래 처럼 사용하면 Scope 벗어나 에러
	// val++

	var name string
	var category = 1

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)

	grade(95)
	check(2)

}

// switch 뒤에 조건변수 혹은 Expression을 적지 않는 용법 => 각 case 조건문들을 순서대로 검사하여 조건에 맞는 경우 해당 case 블럭을 실행하고 switch문을 빠져 나온다.
func grade(score int) {
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}
}

// go 컴파일러는 case문 마지막에 항상 break문을 자동으로 추가한다. 계속 밑의 문장들을 실행하게 하려면 fallthrough문을 명시해 주면 된다.
func check(val int) {
	switch val {
	case 1:
		fmt.Println("1 이하")
		fallthrough
	case 2:
		fmt.Println("2 이하")
		fallthrough
	case 3:
		fmt.Println("3 이하")
		fallthrough
	default:
		fmt.Println("default 도달")
	}
}
