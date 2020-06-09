package main

import "fmt"

func main() {
	a := 98
	fmt.Println(a)
	fmt.Println(&a)
	pointer(&a)
	fmt.Println(a)
	//b := &a
	//fmt.Println(*b)
	//fmt.Println(*&a)
}

func pointer(y *int) {
	fmt.Println(y)
	*y = 101
	fmt.Println(y)
}
