package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Student) greeting() {
	fmt.Println("hello")
}

type Student struct {
	Person // 학생 구조체안에 사람 구조체를 필드로 가지고 있음 has-a 관계
	school string
	grade  int
}

func main() {
	var s Student

	s.greeting()
}
