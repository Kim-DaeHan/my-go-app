package main

import (
	"fmt"
	"log"
	"os"
)

// go 에러
// go는 내장타입으로 error라는 인터페이스 타입을 갖는다
// type error interface {
// 	Error() string
// }

func main() {
	// go 에러 처리
	f, err := os.Open("C:\\temp\\1.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(f.Name())

	// 또 다른 에러처리
	// _, err2 := otherFunc()
	// switch err2.(type) {
	// default: // no error
	// 	fmt.Println("ok")
	// case MyError:
	// 	log.Print("Log my error")
	// case error:
	// 	log.Fatal(err.Error())
	// }
}
