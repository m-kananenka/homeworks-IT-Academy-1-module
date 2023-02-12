package main

import (
	"fmt"
	"sync"

	"time"
)

func main() {

	defer func(t time.Time) {
		fmt.Println(time.Since(t))
	}(time.Now())

	for i := 0; i <= 50000; i++ {
		Fibonacci(i) //293.1508ms
	}

	var wg sync.WaitGroup
	for i := 0; i <= 50000; i++ {
		wg.Add(1)
		go func(i int) {
			SyncFibonacci(i, &wg) //124.0717ms
		}(i)

	}
	wg.Wait()

	//семафор

	sem := make(chan struct{}, 10) // где 10 - это количество одновременно работающий горутин
	for i := 0; i <= 50000; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int) {
			SyncFibonacci(i, &wg) //74.5442ms
			<-sem
		}(i)
	}
	wg.Wait()
}

func Fibonacci(n int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

}

func SyncFibonacci(n int, wg *sync.WaitGroup) {
	Fibonacci(n)
	wg.Done()
}
