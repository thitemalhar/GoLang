package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Last  string
	Age   int
}

func main() {
	u1 := user {
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	u2 := user{
		First: "Iron",
		Last:  "Man",
		Age:   42,
	}

	//fmt.Println(u1, u2)
	people := []user{u1,u2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(string(bs))
}