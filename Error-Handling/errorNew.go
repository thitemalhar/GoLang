package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	//_, err := sqrt(-10.63)
	//if err != nil {
	//	log.Fatal(err)
	//}

	_, err := sqaure(-12)
	if err != nil {
		log.Fatal(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("Square root of negative number is unknown")
	}
	return 42, nil
}

func sqaure(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("This number cannot be squared: %v", f)
	}
	return 1, nil
}
