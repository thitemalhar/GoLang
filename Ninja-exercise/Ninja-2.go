package main

import "fmt"

const (
	a1 = 2020 + iota
	a2 = 2020 + iota
	a3 = 2020 + iota
	a4 = 2020 + iota
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	x := 42
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)

	a := 42 == 42
	b := 42 <= 43
	c := 42 >= 43
	d := 42 != 43
	e := 42 < 43
	f := 42 > 43
	fmt.Println(a, b, c, d, e, f)

	z := 55
	fmt.Printf("%d\t%b\t%#x\n", z, z, z)
	y := z << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)

	fmt.Println("")
	fmt.Println("Print the values of the iota year")
	fmt.Println(a1, a2, a3, a4)

	fmt.Println("")
	fmt.Println("Decial and Hexdecimal value of the KB, MB and GB are: ")
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b", gb, gb)
}
