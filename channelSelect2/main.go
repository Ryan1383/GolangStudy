package main

import(
	"fmt"
	"time"
)

func main() {
	c1:= make(chan int) 
	c2:= make(chan string)
	
	go func() {
		for{
			i:= <-c1 // 채널 c1에서 값을 꺼낸 뒤 i에 대입
			fmt.Println("c1 : ", i) // i값을 출력
			time.Sleep(100*time.Millisecond) 
		}
	}()
	
	go func() {
		for{
			c2 <-"Hello, world!" //채널 c2에 Hello, world를 보냄
			time.Sleep(500*time.Millisecond) 
		}
	}()
	
	go func() {
		for{
			select{
			case c1<-10: // 매번 c1에 10을 보낸다
			case s:= <-c2: //c2에 값이 들어왔을 때는 값을 꺼낸뒤 s에 대입
				fmt.Println("c2 :", s)
			}
		}
	}()
	
	time.Sleep(10*time.Second) 
	
}