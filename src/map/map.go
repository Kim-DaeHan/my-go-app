package main

import "fmt"

func main() {

	// map: key: value 형태
	var idMap map[int]string // int: string 형태 map
	// 위처럼 선언할 경우 nil 값을 가지는 nil map에는 어떤 데이터를 쓸 수 없다. 초기화하기 위해 make() 함수 사용 가능
	idMap = make(map[int]string)
	fmt.Println(idMap)

	// 리터럴을 사용한 초기화
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
	}

	fmt.Println(tickers)

	var m map[int]string

	m = make(map[int]string)
	// 추가 혹은 갱신
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	// 키에 대한 값 읽기
	str := m[134]
	fmt.Println(str)
	noData := m[999] // 값이 없으면 nil 혹은 zero 리턴
	fmt.Println(noData)

	// 삭제
	delete(m, 777)
	fmt.Println(m)

	// Map 키 체크
	// Go에선 map 변수[키] 읽기를 수행할 때 2개의 리턴값을 리턴한다. 첫번째는 키에 상응하는 값, 두번째는 그 키가 존재하는지 아닌지를 나타내는 bool값
	val, exists := tickers["MSFT"]
	if !exists {
		fmt.Println("No MSFT ticker")
	}

	fmt.Println(val, exists) // Microsoft, true

	// for 루프를 사용한 map 열거
	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}

	// for range 문을 사용하여 모든 맵 요소 출력
	// Map은 unordered 이므로 순서는 무작위
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}
