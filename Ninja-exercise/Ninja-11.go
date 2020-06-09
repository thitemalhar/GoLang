package main

import "fmt"

type CustomerErr struct {
	info string
}

func (ce CustomerErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", ce.info)
}

func main() {
	c1 := CustomerErr{
		info: "Need more coffee",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println("Here is the foo error", e, "\n", e.(CustomerErr).info)
}