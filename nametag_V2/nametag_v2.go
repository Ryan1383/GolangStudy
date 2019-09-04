package main

import "fmt"

type Person struct {
	id      int
	name    string
	address string
	phone   string
}

func (p Person) printNametag() {
	fmt.Println("------------NameTag-----------")
	fmt.Printf("ID: %d \n이름: %s \n주소: %s \n전화번호: %s \n", p.id, p.name, p.address, p.phone)
	fmt.Println("------------------------------")
}

func main() {
	person1 := Person{1, "이인", "수원", "010-2476-1383"}
	person2 := Person{2, "라이언", "서울", "02-1444-4421"}
	person3 := Person{3, "제임스", "충주", "041-322-2212"}

	var person_slice []Person
	person_slice = append(person_slice, person1, person2, person3)

	for i := 0; i < len(person_slice); i++ {
		person_slice[i].printNametag()
	}
}
