package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go workerProcess(c)
	go workerProcess(c)
	go workerProcess(c)
	go publisher(c)
	go publisher(c)
	go publisher(c)
	go publisher(c)
	time.Sleep(10 * time.Millisecond)
}

var workerID int
var publisherID int

//NO FAN-IN
func publisher(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is publishing data\n", dataID)
		data := fmt.Sprintf("data from publisher %d. Data is %d", thisID, dataID)
		out <- data

	}
}

//FAN OUT
func workerProcess(in <-chan string) {
	workerID++
	thisID := workerID
	for {
		fmt.Printf("%d waiting for input ", thisID)
		input := <-in
		fmt.Println("Input is : ", input)
	}

}
