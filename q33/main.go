package main

import "fmt"

var p int
var q int

func main() {
	p = 20
	q = 30
	swap(&p, &q)
}

func swap(p *int, q *int) {
	fmt.Println(*p)
	fmt.Println(*q)

	tempP := *p
	tempQ := *q

	*p = tempQ
	*q = tempP

	fmt.Println(*p)
	fmt.Println(*q)

}
