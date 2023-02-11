package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//Fibonacci(1000) // 8.4734ms
	fmt.Println()

	//go Fibonacci(1000)//0s
	fmt.Println()

	var wg sync.WaitGroup

	for i := 0; i <= 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Fibonacci(i) //999.3µs
		}(i)
		wg.Wait()
	}
}

func Fibonacci(n int) int {
	defer func(t time.Time) {
		fmt.Println(time.Since(t))
	}(time.Now())
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Print(x, " ")
	}
	return x
}
