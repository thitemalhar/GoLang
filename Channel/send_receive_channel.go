package main

import "fmt"

func main() {
	c := make(chan int)
	r := make(chan int)
	//send
	go foo(c)
	//receive
	bar(c)

	fmt.Println("Time to exit")

	// send - this for the range example
	go func(){
		for i:= 0; i< 5; i++ {
			r <- i
		}
		close(r)		//Need to close the channel after loop is complete
	}()

	//receive
	for v := range r {
		fmt.Println(v)
	}
}

//send
func foo(c chan<- int)  {
	c <- 21
}
//receive
func bar(c <-chan int)  {
	fmt.Println(<-c)
}
