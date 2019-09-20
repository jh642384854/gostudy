package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	runtime.GOMAXPROCS(1)
	wg.Add(2)
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finsh")
	wg.Wait()
	fmt.Println("Terminating Program")
}

func printPrime(str string)  {
	defer wg.Done()
next:
	for outer := 2; outer < 500; outer++  {
		for inner := 2; inner< outer ;inner++  {
			if outer % inner == 0{
				continue next
			}
		}
		fmt.Printf("%s:%d \n",str,outer)
	}
	fmt.Println("Completed",str)
}