package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	ns, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(ns)
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("error:", err)
		log.Println("error:", err)
		//	log.Fatalln("error:", err)
		//		panic(err)
	}
}
