package main

import (
	"fmt"
	"time"
)


func main() {
	done := make(chan bool) // 동기채널 생성
	count := 3 //반복 횟수
	
	go func() {
		for i:= 0; i<count; i++{
			done <-true  // 고루틴에 true를 보냄, 값을 꺼낼때까지 대기
			fmt.Println("고루틴 : ", i) //반복문의 변수출력
			time.Sleep(1 * time.Second) //1초대기
		}
	}()
	
	for i:=0; i<count; i++ {
		<-done // 채널에 값이 들어올때가지 대기 , 값을 꺼냄
		fmt.Println("메인함수 : ", i) //반복문의 변수출력
		
	}
	
}