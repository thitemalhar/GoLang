package main

import (
	"fmt"
	"runtime"
)

var x int = 42
var y string = "James Bond"
var z bool = true

type (
	hotdog int
)

var a hotdog
var b int

func main() {
	fmt.Printf("%v\n", a)

	fmt.Printf("%T\n", a)
	a = 12

	fmt.Println(a)

	b = int(a)
	fmt.Println(b)

	fmt.Printf("%T\n", b)

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
