package main

import (
	"fmt"
	"sort"
)

type agent struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := agent{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := agent{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := agent{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []agent{u1, u2, u3}

	fmt.Println(users)

	sort.Slice(users, func(i, j int) bool {
		return users[i].Last < users[j].Last
	})

	fmt.Println("***********Sort by Last Name***********")
	for i, person := range users {
		fmt.Println("Person number:", i)
		fmt.Println("\t", person.First, person.Last, person.Age)
		for _, v := range person.Sayings {
			fmt.Println("\t\t", v)
		}
	}

	fmt.Println("***********Sort by Age***********")
	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	for i, person := range users {
		fmt.Println("Person number:", i)
		fmt.Println("\t", person.First, person.Last, person.Age)
		for _, v := range person.Sayings {
			fmt.Println("\t\t", v)
		}
	}
}
