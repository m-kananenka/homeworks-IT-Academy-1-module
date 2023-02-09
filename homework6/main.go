package main

import (
	"fmt"
	"time"
)

func main() {

	// асинхронный запуск
	for i := 0; i < 10; i++ {
		go Fibonacci(i)
	}
	time.Sleep(time.Second)

	//синхронный запуск
	c := make(chan int, 10)
	go Fib(cap(c), c)
	for i := range c {
		fmt.Println(i)

	}
}

func Fibonacci(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	fmt.Println(x)
	return x

}

// для синхронизации значений пишем в канал с
func Fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	defer close(c)
}
