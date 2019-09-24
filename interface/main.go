// package main

// import (
// 	"math"
// )

// /*
// 	interface 는 메서드들의 집합체이다. interface는 타입이 구현해야 하는 메서드 원형(prototype) 을 정의한다.
// 	하나의 사용자 정의타입이 interface를 구현하기 위해서는 단순하게 그 인터페이스가 가지고 있는 모든 메서드들을 구현하면 된다.
// */

// type Shape interface {
// 	area() float64
// 	perimeter() float64
// }

// // Rect 정의
// type Rect struct {
// 	width, height float64
// }

// //Circle 정의
// type Circle struct {
// 	radius float64
// }

// //Rect 타입에 대한 Shape 인터페이스 구현
// func (r Rect) area() float64 {
// 	return r.width * r.height
// }

// func (r Rect) perimeter() float64 {
// 	return 2 * (r.width + r.height)
// }

// //Circle 타입에 대한 Shape 인터페이스 구현
// func (c Circle) area() float64 {
// 	return math.Pi * c.radius
// }

// func (c Circle) perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

// /*
//  	인터페이스를 사용하는 일반적인 예로 함수가 파라미터로 인터페이스를 받아들이는 경우가 있다.
//  	함수 파라미터가 interface 인 경우, 이는 어떤 타입이든 해당 인터페이스를 구현하기만 하면 모두 입력 파라미터로
// 	 사용될 수 있다는 것을 의미한다.

// */

// func main() {
// 	r := Rect{10., 20.}
// 	c := Circle{10}

// 	showArea()
// }

// func showArea(shapes ...Shape) {
// 	for _, s := range shapes {
// 		a := s.area() // 인터페이스 메서드 호출
// 		println(a)
// 	}
// }

// /*
// 	Empty interface는 메서드를 전혀 갖지 않는 빈 인터페이스로서 , go의 모든Type은 적어도 0개의 메서드를 구현하므로
// 	흔히 Go에서 모든 Type을 나타내기 위해 빈 인터페이스를 사용한다. 즉, 빈 인터페이스는 어떠한 타입도 담을 수 있는 컨테이너라고
// 	볼 수 있으며, 흔히 일컫는 Dynamic Type 이라고 볼 수 있다.

// */
