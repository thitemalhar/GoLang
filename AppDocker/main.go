package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	fmt.Println("This is the Go App and docker tutorial")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hello from docker app")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
