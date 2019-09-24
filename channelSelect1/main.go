package main

import (
	"fmt"
	"time"
)

func main() {
	c1:= make(chan int) // int형 채널을 생성
	c2:= make(chan string) // string형 채널을 생성
	
	go func() {
		for{
			c1<-10 // 채널 c1에 10을 보낸뒤
			time.Sleep(100*time.Millisecond) //100밀리초 대기
		}
	}()
	
	go func(){
		for{
			c2<-"Hello, world!" // 채널 c2에 Hello,world를 보낸뒤
			time.Sleep(500*time.Millisecond) // 500밀리초 대기
		}
	}()
	
	go func() {
		for{
			select{
			case i:= <-c1: //채널 c1에 값이 들어왔다면 값을 꺼내서 i에 대입
				fmt.Println("c1 : ", i) // i값을 출력
				
			case s:= <-c2: // 채널 c2에 값이 들어왔다면 값을 꺼내서 s에 대입
				fmt.Println("c2: ", s)	
			}
		}
	}()
	
	time.Sleep(10*time.Second) // 10초동안 프로그램 실행
}