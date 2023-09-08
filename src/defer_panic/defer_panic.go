package main

import (
	"fmt"
	"os"
)

// 지연실행 defer
// go 언어의 defer 키워드는 특정 문장 혹은 함수를 나중에 (defer를 호출하는 함수가 리턴하기 직전에) 실행하게 한다.
// 일반적으로 defer는 C#, Java 같은 언어에서의 finally 블럭처럼 마지막에 Clean-up 작업을 위해 사용된다.

// panic 함수
// go 내장함수인 panic() 함수는 현재 함수를 즉시 멈추고 현재 함수에 defer 함수들을 모두 실행한 후 즉시 리턴한다.
// 이러한 panic 모드 실행 방식은 다시 상위함수에도 똑같이 적용되고, 계속 콜스택을 타고 올라가며 적용된다.
// 그리고 마지막에는 프로그램이 에러를 내고 종료하게 된다.

// recover 함수
// go 내장함수인 recover() 함수는 panic 함수에 의한 패닉상태를 다시 정상상태로 되돌리는 함수이다.
// recover() 함수를 사용하면 panic 상태를 제거하고 openFile()의 다음 문장인 println()을 호출하게 된다.

func main() {

	// f, err := os.Open("1.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// main 마지막에 파일 close 실행 => 어떤 에러가 발생하더라도 항상 파일을 Close 할 수 있도록한다.
	// defer f.Close()

	// 파일 읽기
	// bytes := make([]byte, 1024)
	// f.Read(bytes)
	// fmt.Println(len(bytes))

	openFile("Invalid.txt")

	fmt.Println("Done")
}

func openFile(fn string) {
	// defer 함수. pacin 호출시 실행됨
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Open Error", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}
