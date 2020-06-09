package main

import "fmt"

type user struct {
	first string
	last string
	icecream []string
}

func main() {

	p1 :=  user {
		first: "Malhar",
		last: "Thite",
		icecream: []string{"Chocolate", "Vanilla", "Orange"},
	}

	p2 :=  user {
		first: "Kavita",
		last: "Todkar",
		icecream: []string{"Vanilla", "Mango", "Butterscotch"},
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last)
	for i, v := range p1.icecream {
		fmt.Println(i, v)
	}

	personMap := map[string]user {
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(personMap)

	for _, v := range personMap {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for _, val := range v.icecream {
			fmt.Println("\t", val)
		}
	}
}
