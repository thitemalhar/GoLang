package main

import "fmt"

func main() {
	x := []int{2, 5, 44, 3, 4, 5}
	fmt.Println(x)

	s := []int{4, 5, 6, 8, 12, 82}
	fmt.Println(s)

	t := []int{100, 99, 98, 97}
	fmt.Println(t)

	s = append(s, t...) //'...' is added to the end to append 't' to 's'
	fmt.Println(s)

	//delete an element from slice
	s = append(s[:3], s[4:]...)
	fmt.Println(s)

	//create a slice with 'make'
	a := make([]int, 10, 12)
	fmt.Println(a)
	fmt.Println("Length of the slice is %d", len(a))
	fmt.Println("Capacity of the slice is %d", cap(a))

}
