package main

import "fmt"

type Student struct {
	name    string
	class   int
	sungjuk Sungjuk // Sungjuk 타입을 만들어서 (커스텀) Student의 필드로
}

type Sungjuk struct {
	name  string
	grade string
}

// Student에 속한 메서드
func (s Student) ViewSungJuk() {
	fmt.Println(s.sungjuk)
}

func ViewSungJuk(s Student) { //구조체에 속하지 않은 함수
	fmt.Println(s.sungjuk)
}

func (s Student) InputSungjuk(name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
	fmt.Println(s.sungjuk)

}

func InputSungjuk(s Student, name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
	fmt.Println(s.sungjuk)

}

func main() {
	var s Student
	s.name = "철수"
	s.class = 1

	s.sungjuk.name = "수학"
	s.sungjuk.grade = "C"

	s.ViewSungJuk()
	ViewSungJuk(s)

	s.InputSungjuk("과학", "A")
	s.ViewSungJuk()
	InputSungjuk(s, "과학", "B")

}
