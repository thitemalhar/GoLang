/**
	Area of square and rectangle (struct, interface, func - https://play.golang.org/p/VFvjMpuAYka
 */

package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, " and I am ", p.age, "years old")
}

func main() {
	ii1 := []int{1,2,3,4,5,6}
	a1 := foo1(ii1...)
	fmt.Printf("Sum of %d is : %v\n ", ii1, a1)

	ii2 := []int{1,2,3,4,5,6}
	a2 := bar1(ii2)
	fmt.Printf("Sum of %d is : %v\n ", ii2, a2)

	p1 := person{
		first: "Malhar",
		last:  "Thite",
		age:   32,
	}
	p1.speak()
}

func foo1(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar1(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
