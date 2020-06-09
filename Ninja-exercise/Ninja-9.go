package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
func main() {
	wg.Add(2)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	go firstFunc()
	go secondFunc()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("About to exit")
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func firstFunc()  {
	fmt.Println("This message is from firstFunc")
	wg.Done()
}

func secondFunc()  {
	fmt.Println("This message is from secondFunc")
	wg.Done()
}