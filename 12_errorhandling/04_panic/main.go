package main

import "os"

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("error:", err)
		// log.Println("error:", err)
		//	log.Fatalln("error:", err)
		panic(err)
	}
}
