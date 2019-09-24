package main

import (
	"fmt"
	"os"
)

/*
 파일 읽기
 방법 1.
   - os 패키지 사용
*/

/*
	func Create9name string) (file *File, err error) : 기존 파일을 열거나 새 파일을 생성
	func (f *File) Close() error : 열린 파일을 닫음.
	func (f *File) Write(b []byte)(n int, err error) : 파일에 값을 씀, 파일에 쓴 데이터의 길이와 에러값을 리턴

*/
func main() {
	file, err := os.Create("./hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	s := "Hellow World"

	n, err := file.Write([]byte(s)) // s를 []byte 슬라이스로 변환, s를 파일에 저장.
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n, "바이트 저장 완료")
}
