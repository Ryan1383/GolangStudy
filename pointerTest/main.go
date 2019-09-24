// package main

// import "fmt"

// type Rectangle struct {
// 	width, height int
// }

// func rectangleArea(rect *Rectangle) int { // 매개변수로 구조체 포인터를 받음
// 	return rect.width * rect.height
// }

// func rectabgleScaleA(rect *Rectangle, factor int) { //매개변수로 구조체 포인터를 받음
// 	rect.width = rect.width * factor   // 포인터이므로 원래의 값이 변경됨
// 	rect.height = rect.height * factor // 포인터이므로 원래의 값이 변경됨
// }

// func rectangleScaleB(rect Rectangle, factor int) { // 매개변수로 구조체 인스턴스를 받음
// 	rect.width = rect.width * factor   // 값이 복사되었으므로 원래의 값에는 영향이 안감
// 	rect.height = rect.height * factor // 값이 복사되었으므로 원래의 값에는 영향이 안감
// }

// func (rect *Rectangle) area() int {
// 	return rect.width * rect.height
// 	// 리시버 변수를 사용하여 현재 인스턴스에 접근이 가능
// }

// func main() {

// 	rect1 := Rectangle{30, 30}
// 	rectabgleScaleA(&rect1, 10) // 구조체의 포인터를 넘김
// 	fmt.Println(rect1)          // 바뀐값이 들어감

// 	rect2 := Rectangle{30, 30}
// 	rectangleScaleB(rect2, 10) // 구조체 인스턴스 그대로 넘김
// 	fmt.Println(rect2)         //안바뀜

// 	// 보통 rectangleArea(rect *Rectangle) 과 같이 함수의 매개변수에 구조체 포인터를 받는다. 그리고 값을 넘겨준 때도 rect변수에 &를 붙여
// 	// 주소를 넘겨준다
// 	// 함수의 매개변수에 구조체 포인터가 아닌 일반적인 형태(구조체 인스턴스) 로 넘겨주면 값이 모두 복사되므로 주의한다.

// }
