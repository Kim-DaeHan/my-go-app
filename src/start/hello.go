package main

import "fmt"

func main() {

	// iota라는 identifier를 사용하면 상수 값을 0부터 순차적으로 부여
	const (
		Apple  = iota // 0
		Grape         // 1
		Orange        // 2
	)

	// 25개의 예약 키워드들은 변수명, 상수명, 함수명 등의 Identifier로 사용 불가능
	//     break        default      func         interface    select
	//     case         defer        go           map          struct
	//     chan         else         goto         package      switch
	//     const        fallthrough  if           range        type
	//     continue     for          import       return       var

	// 형변환
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	fmt.Println(u)
	fmt.Println(f)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	fmt.Println(bytes)
	fmt.Println(str2)

	var a int = 100
	fmt.Println(a)
	a *= 10
	fmt.Println(a)
	a >>= 2
	fmt.Println(a)
	a |= 1
	fmt.Println(a)
	fmt.Println(Apple, Grape, Orange)

	var k int = 10
	var p = &k      // k의 주소를 할당
	fmt.Println(*p) // p가 가리키는 주소에 있는 실제 내용을 출력
}
