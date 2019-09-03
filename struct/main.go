package main

import "fmt"

// 구조체는 어떠한 개념을 한곳에 모아놓은것이다.
type Person struct {
	name string
	age  int
}

//Person구조체의 메서드이다. p.printName으로 접근
func (p Person) printName() {
	fmt.Println(p.name)
}

func main() {
	var p Person
	p1 := Person{"개똥이", 15}
	p2 := Person{name: "소똥이", age: 21}
	p3 := Person{name: "Carson"}
	p4 := Person{}

	fmt.Println(p, p1, p2, p3, p4)

	p.name = "smith"
	p.age = 24
	fmt.Println(p)
	p2.printName()
}
