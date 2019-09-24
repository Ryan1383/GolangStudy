package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
	r:= rand.Intn(100) // 랜덤 숫자 생성
	time.Sleep(time.Duration(r)) //랜덤 시간 동안 대기
	fmt.Println(n) 
}

func main() {
	for i:= 0; i< 100; i++ {
		go hello(i)
	}
	
	fmt.Scanln()
}
