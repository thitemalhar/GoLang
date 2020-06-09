package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Add(1)
	go foo1()
	bar1()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo1() {
	for i := 0; i < 3; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar1() {
	for i := 0; i < 3; i++ {
		fmt.Println("bar:", i)
	}
}