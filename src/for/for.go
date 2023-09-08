package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 조건식만 쓰는 for 루프: 다른 언어의 while 루프와 같다
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println(n)

	// for range 문: 컬렉션으로 부터 한 요소씩 가져와 차례로 for 블럭의 문장들을 실행
	names := []string{"홍길동", "이순신", "강감찬"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	// break, continue, goto 문: goto는 임의의 문장으로 이동
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for루프 시작으로
		}
		a++
		if a > 10 {
			break //루프 빠져나옴
		}
	}
	if a == 11 {
		goto END //goto 사용예
	}
	println(a)

END:
	println("End")
}
