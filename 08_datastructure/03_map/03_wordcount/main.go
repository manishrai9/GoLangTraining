package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/2701-h/2701-h.htm")
	if err != nil {
		log.Fatalln(err)
	}

	sc := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	sc.Split(bufio.ScanWords)
	buckets := make([]int, 12)
	for sc.Scan() {
		n := hashBucket(sc.Text())
		buckets[n]++
	}

	fmt.Println(buckets)
}

func hashBucket(word string) int {
	return int(word[0]) % 12 //for some random distribution
}
