package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintSungjuk() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) InputSungjuk(class string, grade string) {
	s.class = class
	s.grade = grade
}

func main() {
	var s Student = Student{name: "Ryan", age: 29, class: "수학", grade: "A+"}

	s.InputSungjuk("과학", "c")
	s.PrintSungjuk()
}

// func Increase(x *int) {
// 	// x++
// 	*x = *x + 1

// }

/**
var p *int
p = &a  a 변수의 메모리주소를 p에 할당한다.

fmt.Println(*p)   // p포인터에 할당된 a의 주소를 통해 값을 가져온다.
*/
