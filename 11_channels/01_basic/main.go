package main

import (
	"fmt"
	"time"
)

func main() {
	//unbuffered channel, similar to relay race where next runner waits till first reaches
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			//places int in channel and stops until it is taken off, blocking
			c <- i
		}
	}()

	go func() {
		for {
			//code stops unless something is put on channel, blocking
			fmt.Println(<-c)
		}
	}()
	time.Sleep(time.Second) //need to wait, otherwise main function will exit
}

//go run -race main.go
