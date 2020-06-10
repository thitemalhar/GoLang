package main

import (
	"github.com/thitemalhar/GoLang/Projects/MicroservicesProject/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	//http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	//	log.Println("Hello world!")
	//	d, err := ioutil.ReadAll(r.Body)
	//	if err != nil {
	//		http.Error(rw, "Oops", http.StatusBadRequest)
	//		return
	//	}
	//
	//	fmt.Fprintf(rw, "Hello %s\n", d)
	//})
	//
	//http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
	//	log.Println("Goodbye world!")
	//})
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gb)

	http.ListenAndServe(":9090", sm)
}
