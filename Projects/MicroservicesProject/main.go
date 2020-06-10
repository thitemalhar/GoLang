package main

import (
	"context"
	"github.com/thitemalhar/GoLang/Projects/MicroservicesProject/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gb)

	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}

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
