package main

import (
	"fmt"
	"runtime"
	"sync"
)

// go 루틴은 go 런타임이 관리하는 Lightweight 논리적(혹은 가상적) 쓰레드이다.
// go에서 go 키워드를 사용하여 함수를 호출하면 런타임시 새로운 goroutine을 실행한다. 고루틴은 비동기적으로 함수 루틴을 실행하므로 여러 코드를 동시에 실행하는데 사용된다.
// 고루틴은 OS 쓰레드보다 훨씬 가볍게 비동기 처리를 구현하기 위하여 만든것으로 기본적으로 Go 런타임이 자체 관리
// go 런타임 상에서 관리되는 작업 단위인 고루틴들은 종종 하나의 os 쓰레드 1개로도 실행되곤 한다.
// 즉 고루틴들은 os 쓰레드와 1 대 1로 대응되지 않고 multiplexing 으로 훨씬 적은 os 쓰레드를 사용한다.
// 메모리 측면에서도 os 쓰레드가 1 메가바이트의 스택을 갖는 반면, 고루틴들은 이보다 훨씬 작은 몇 킬로바이트의 스택을 갖는다. go 런타임은 고루틴을 관리하면서
// 고 채널을 통해 고루틴 간의 통신을 쉽게 할 수 있도록 한다.

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

// 익명함수 go루틴
// 고루틴은 익명함수에 대해 사용할 수도 있다. 즉, go 키워드 뒤에 익명함수를 바로 정의하는 것으로 이는 익명함수를 비동기로 실행하게 된다.

// 다중 CPU 처리
// go는 디폴트로 1개의 CPU를 사용한다. 즉 여러 개의 Go 루틴을 만들더라도, 1개의 CPU에서 작업을 시분할하여 처리한다.
// 만약 머신이 복수개의 CPU를 가진 경우, Go 프로그램을 다중 CPU에서 병렬 처리하게 할 수 있는데, 병렬처리를 위해서는 runtime.GOMAXPROCS(CPU수) 함수를 호출하여야 한다

func main() {
	// 4개의 CPU 사용
	runtime.GOMAXPROCS(4)
	// 함수를 동기적으로 실행
	// say("Sync")

	// 함수를 비동기적으로 실행
	// go say("Async1")
	// go say("Async2")
	// go say("Async3")

	// 3초 대기
	// time.Sleep(time.Second * 3)

	// WaitGroup 생성. 2개의 go 루틴을 기다림
	var wait sync.WaitGroup
	wait.Add(2)

	// 익명함수를 사용한 고루틴
	go func() {
		defer wait.Done() // 끝나면 .Done() 호출
		fmt.Println("Hello")
	}()

	// 익명함수에 파라미터 전달
	go func(msg string) {
		defer wait.Done() // 끝나면 .Done() 호출
		fmt.Println(msg)
	}("Hi")

	wait.Wait() // go 루틴 모두 끝날 때 까지 대기
}
