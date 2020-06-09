package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person // inheriting
	ltk    bool
}

func (s secretAgent) speak() { //Methods in GO
	fmt.Println("I am ", s.first, s.last, " and my age is ", s.age)
	if s.ltk == true {
		fmt.Println("\t and I have license to kill")
	}
}

func main() {

	//composite data structure or aggregate data structure
	p1 := person{
		first: "Malhar",
		last:  "Thite",
		age:   32,
	}

	p2 := person{
		first: "Kavita",
		last:  "Todkar",
		age:   35,
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)

	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)

	sa1.speak()

}
