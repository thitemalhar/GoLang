// https://play.golang.org/p/T3Ep0ccvbPi

package main

import "fmt"

func main() {
	defer foo() //defer will run that statement at the very end - eg: closing a file after opening
	xi := []int{1, 2, 3, 4, 5}
	add := sum(xi...) //Cannot pass in the slice of int to the sum(int) parameter..
	//by adding "..." it's calling furling and unfurling
	fmt.Println(`The total is`, add)

	func() {
		fmt.Println("this is anonymous func")
	}()

	f := func() {
		fmt.Println("This is f func")
	}
	f()

	numOfSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evenNum := even(sum, numOfSlice...)
	fmt.Println("Sum of even number in numOfSum is ", evenNum)

	fmt.Println("******************************")
	fmt.Println("Example for Recurssion")
	a := factorial(4)
	fmt.Println("\tFactorial for the 4 is :", a)
	fmt.Println("******************************")

}

func factorial(fact int) int {
	if fact == 0 {
		return 1
	}
	return fact * factorial(fact-1)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func foo() {
	fmt.Println("this is foo()")
}

func even(f func(x ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
