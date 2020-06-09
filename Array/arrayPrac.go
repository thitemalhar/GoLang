package main

import (
	"fmt"
)

func main() {

	var a [5]int
	fmt.Println(a)

	a[2] = 2
	fmt.Println(a)

	b := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(b)

	var twoD [2][3]int
	fmt.Println(twoD)

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

}
