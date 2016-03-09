package main

import "fmt"

func main() {
	//unbuffered channel, similar to relay race where next runner waits till first reaches
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			//places int in channel and stops until it is taken off, blocking
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

//go run -race main.go
