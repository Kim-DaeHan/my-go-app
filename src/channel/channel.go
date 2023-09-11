package main

import (
	"fmt"
	"time"
)

// go 채널
// 데이터를 주고 받는 통로, 채널은 make() 함수를 통해 미리 생성되어야 하며, 채널 연산자 <- 을 통해 데이터를 보내고 받는다.
// 채널은 흔히 고루틴들 사이 데이터를 주고 받는데 사용되는데 상대편이 준비될 때까지 채널에서 대기함으로써 별도의 lock을 걸지 않고 데이터를 동기화하는데 사용된다.
// go 채널은 수신자와 송신자가 서로를 기다리는 속성때문에, 이를 이용하여 고 루틴이 끌날 때 까지 기다리는 기능을 구현할 수 있다.
// 즉, 익명함수를 사용한 한 고루틴에서 어떤 작업이 실행되고 있을 때, 메인 루틴은 <- done에서 계속 수신하며 대기하고 있게 된다.
// 익명함수 고루틴에서 작업이 끝난 후, done 채널에 true를 보내면, 수신자 메인루틴은 이를 받고 프로그램을 끝내게 된다.

func main() {

	// 정수형 채널 생성
	ch := make(chan int)

	go func() {
		ch <- 123 // 채널에 123을 보낸다
	}()

	var i int
	i = <-ch // 채널로부터 123을 받는다
	fmt.Println(i)

	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	// 위의 고루틴이 끝날 때까지 대기
	<-done
	fmt.Println("finish")

	// go 채널 버퍼링
	// Unbuffered Channel, Buffered Channel 존재, 위의 채널은 Unbuffered Channel로서 이 채널에서는 하나의 수신자가 데이터를 받을 때까지 송신자가 데이터를 보내는 채널에 묶여 있게 됨
	// Buffered Channel을 사용하면 비록 수신자가 받을 준비가 되어 있지 않을 지라도 지정된 버퍼만큼 데이터를 보내고 계속 다른 일을 수행할 수 있음
	// 버퍼 채널은 make(chan type, N) 함수를 통해 생성되는데, 두번째 파라미터 N에 사용할 버퍼 갯수를 넣는다
	// 버퍼 채널을 이용하지 않는 경우, 아래와 같은 코드는 에러를 발생시킴 왜냐하면 메인루틴에서 채널에 1을 보내면서 상대편 수신자를 기다리고 있는데, 이 채널을 받는 수신자 고루틴이 없기 때문
	// c := make(chan int)
	// c <- 1           // 수신 루닡이 없으므로 데드락
	// fmt.Println(<-c) // 코멘트해도 데드락(별도의 고루틴 없기 때문)

	// 하지만 아래와 같이 버퍼채널을 사용하면, 수신자가 당장 없더라도 최대버퍼 수까지 데이터를 보낼 수 있으므로 에러 발생 x
	ch1 := make(chan int, 1)

	// 수신자가 없더라도 보낼 수 있음
	ch1 <- 101

	fmt.Println(<-ch1)

	// 채널 파라미터
	// 채널을 함수의 파라미터로 전달할 때, 일반적으로 송수신을 모두 하는 채널을 전달하지만, 특별히 해당 채널로 송신만 할 것인지 혹은 수신만할 것인지를 지정할 수도 있음
	ch2 := make(chan string, 1)
	sendChan(ch2)
	receiveChan(ch2)

	// 채널 닫기
	// close() 함수를 사용하여 채널을 닫을 수 있음
	// 채널을 닫게 되면 해당 채널로는 더이상 송신을 할 수 없지만, 채널이 닫힌 후에도 계속 수신은 가능
	// 채널 수신에 사용되는 <-ch은 두개의 리턴값을 갖는데, 첫째는 채널의 메시지이고, 두번째는 수신이 제대로 되었는가를 나타냄
	// 만약 채널이 닫혔다면 두번째 리턴값은 false를 리턴

	ch3 := make(chan int, 2)

	// 채널에 송신
	ch3 <- 1
	ch3 <- 2

	// 채널을 닫는다.
	close(ch3)

	// 채널 수신
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)

	if _, success := <-ch3; !success {
		fmt.Println("더이상 데이터 없음")
	}

	// 채널 range문
	// 채널에서 송신자가 송신을 한 후, 채널을 닫을 수 있다. 그리고 수신자는 임의의 갯수의 데이터를 채널이 닫힐 때까지 계속 수신할 수 있다.
	ch4 := make(chan int, 2)

	// 채널에 송신
	ch4 <- 1
	ch4 <- 2

	// 채널을 닫는다
	close(ch4)

	// 방법1
	// 채널이 닫힌 것을 감지할 때까지 계속 수신
	// for {
	// 	if i, success := <-ch4; success {
	// 		fmt.Println(i)
	// 	} else {
	// 		break
	// 	}
	// }

	// 방법2
	// 위의 표현과 동일한 채널 range문
	for i := range ch4 {
		fmt.Println(i)
	}

	// 채널 select 문
	// go의 select문은 복수 채널들을 기다리면서 준비된 채널을 실행하는 기능을 제공
	// 즉, select문은 여러 개의 case문에서 각각 다른 채널을 기다리다가 준비가 된 채널 case를 실행하는 것
	// select문은 case 채널들이 준비되지 않으면 계속 대기하게 되고, 가장 먼저 도착한 채널의 case를 실행한다. 만약 복수 채널에 신호가 오면 go 런타임이 랜덤하게 그 중 한 개를 선택한다.
	// 하지만 select문에 default문이 있으면, case 문 채널이 준비되지 않더라도 계속 대기하지 않고 바로 default문을 실행한다.
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			fmt.Println("run1 완료")
		case <-done2:
			fmt.Println("run2 완료")
			break EXIT // go의 break 레이블은 다른 언어에서의 goto문과 다르게 해당 레이블(EXIT)로 이동한 후 자신이 빠져나온 루프 다음 문장을 실행 => 여기서는 for 루프 다음 즉 main() 함수의 끝
		}
	}
}

func sendChan(ch chan<- string) {
	ch <- "Data"
	// x := <-ch // 에러 발생(송신전용 채널에 수신을 시도해서 에러 발생)
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
