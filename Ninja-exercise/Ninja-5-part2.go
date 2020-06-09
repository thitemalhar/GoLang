package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	truck1 := truck{
		vehicle:   vehicle{
			doors: 2,
			color: "Red",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "Black",
		},
		luxury:  true,
	}

	fmt.Println(truck1)
	fmt.Println(sedan1)

	fmt.Println(truck1.color, truck1.doors, truck1.fourWheel)
	fmt.Println(sedan1.color, sedan1.doors, sedan1.luxury)
}