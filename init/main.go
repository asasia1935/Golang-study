package main

import (
	"fmt"
	"net/http"
)

// 시작점

func main() {
	http.HandleFunc("/", helloWorld) // 내장 패키지 활용

	// 서버가 죽지 않고 켜두는 용도
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("에러가 떴습니다.")
		panic(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}
