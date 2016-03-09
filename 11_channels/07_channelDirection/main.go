package main

import "fmt"

func main() {
	c := incrementor()
	chSum := puller(c)

	for n := range chSum {
		fmt.Println(n)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for num := range c {
			sum += num
		}
		out <- sum
		close(out)
	}()
	return out

}

// the optional <- operator specifies channel direction, send or receive
//https://golang.org/ref/spec#Channel_types
//go run -race main.go
