package main

import "fmt"

func main() {
	done := make(chan struct{})
	intStream := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			intStream <- i
		}
		close(done)
	}()
	printIntegers(done, intStream)
}

func printIntegers(done <-chan struct{}, intStream <-chan int) {
	for {
		select {
		case i := <-intStream:
			fmt.Println(i)
		case <-done:
			return
		}
	}
}
