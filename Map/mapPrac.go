package main

import "fmt"

func main() {
	m := map[string]int{
		"Ojas":   1,
		"Avyukt": 2,
		"Tanush": 3,
	}
	fmt.Println(m)
	fmt.Println(m["Ojas"])

	m["Malhar"] = 4 //Add key value (entry) to the map
	fmt.Println(m["Malhar"])

	if v, ok := m["Tanush"]; ok {
		fmt.Printf("The value is %d\n", v)
	}

	for k, v := range m {
		fmt.Printf("%v : %d\n", k, v)
	}

	fmt.Println("Length of map is", len(m))

	if v, ok := m["Malhar"]; ok {
		fmt.Println("Found key Malhar and deleting it which has value ", v)
		delete(m, "Malhar") // Deleting a Key from the Map "delete(mapName, "key")
	}

	fmt.Println(m)
}
