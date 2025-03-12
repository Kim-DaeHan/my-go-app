package main

import (
	"fmt"
	"my-go-app/src/testlib"
)

func main() {
	fmt.Println("Hello, World!")
	song := testlib.GetMusic("Alicia Keys")
	fmt.Println(song)
}

// package main

// func main() {
// 	println("Test")
// }
