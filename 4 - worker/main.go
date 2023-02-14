package main

import "fmt"

func main() {
	in := make(chan int)
	done := make(chan bool)

	go worker(in, done)
	for i := 0; i < 50; i++ {
		in <- i
	}

	close(in)
	<-done
}

func worker(in <-chan int, done chan<- bool) {
	for v := range in {
		fmt.Printf("Processing: %v\n", v)
	}

	done <- true
}
