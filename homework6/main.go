package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	//асинхронный запуск
	for i := 0; i < 10; i++ {
		go PrintNum(i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println()

	//синхронный запуск
	c := make(chan int)
	go PrintNumSync(0, c)

	for j := range c {
		fmt.Print(j, " ")

	}
	fmt.Println()

	//Worker Pools
	in := make(chan int)
	out := make(chan int)

	//10 инстансов
	for i := 0; i < 10; i++ {
		go PrintNumWp(in, out)
	}
	for job := 0; job < 10; job++ {
		in <- job
	}
	close(in)

	//собираем результаты- это гарантирует, что все горутины закончились
	for k := 0; k < 10; k++ {
		<-out
	}
}

func PrintNum(num int) {
	sum := (num * num) + (num*num)/int(math.Sqrt(3))
	time.Sleep(2 * time.Second)
	fmt.Print(sum, " ")

}

func PrintNumSync(num int, c chan int) {
	for i := 0; i < 10; i++ {
		num = i
		sum := (num * num) + (num*num)/int(math.Sqrt(3))
		time.Sleep(time.Second * 1)
		c <- sum
	}
	close(c)
}

func PrintNumWp(in <-chan int, out chan<- int) {
	for i := range in {
		sum := (i * i) + (i*i)/int(math.Sqrt(3))
		time.Sleep(time.Second * 1)
		fmt.Print(sum, " ")
		out <- sum
	}

}
