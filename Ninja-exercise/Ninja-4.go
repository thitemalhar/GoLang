package main

import "fmt"

func main() {
	x := [5]int{42, 54, 56, 78, 68} //Array

	for i, v := range x {
		fmt.Printf("Index is %v and value is %v\n", i, v)
	}
	fmt.Printf("%T\n", x)

	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}  	//Slice
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", s)

	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:7])
	fmt.Println(s[1:6])

	//Appending values to the slice

	s = append(s, 52)
	fmt.Println(s)

	s = append(s, 53, 54, 55)
	fmt.Println(s)

	y := []int{56, 57, 58, 59, 60}

	s = append(s, y...)
	fmt.Println(s)

	//Deleting the slice
	fmt.Println("Deleting the slice with appending it")
	t := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a := []int {}
	a = append(t[:3])
	a = append(a, t[6:]...)
	fmt.Print(a)
	fmt.Println("\n")

	//Create a slice of US states
	states := make([]string, 1, 500)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	states = append(states, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println("length of states: ", len(states))
	fmt.Println("Capacity of the states slice is: ", cap(states))

	//for i := 0; i < len(states); i++ {
	//	fmt.Println(i, states[i])
	//}

	//Mulidimensional array
	fmt.Println("     ")
	fmt.Println("Mulidimensional array")
	xa := []string{"James", "Bond", "Shaken, not stirred"}
	xb := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	xi := [][]string{xa, xb}
	fmt.Println(xi)

	//MAP exercise
	fmt.Println("     ")
	fmt.Println("MAP exercise")

	maps := map[string][]string {
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(maps)
	for k, v := range maps {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	maps["Malhar"] = []string{`Food`, `Music`, `Movies`}
	for k, v := range maps {
		fmt.Println("This is the record for", k, v)
	}

	if v, ok := maps["Malhar"]; ok {
		fmt.Println("Found key Malhar and deleting it which has value ", v)
		delete(maps, "Malhar")
	}
	for k, v := range maps {
		fmt.Println("This is the record for", k, v)
	}

}
