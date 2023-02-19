package main

import (
	"fmt"
	"sync"
)

func main() {
	out := fanin(
		generateNumbers(1, 30),
		generateNumbers(31, 60),
		generateNumbers(61, 90),
	)

	for value := range out {
		fmt.Printf("value: %v\n", value)
	}
}

func fanin(in ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(in))

	for _, v := range in {
		go func(chIn <-chan int) {
			for value := range chIn {
				out <- value
			}
			wg.Done()
		}(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func generateNumbers(start, final int) <-chan int {
	out := make(chan int)
	go func() {
		for i := start; i <= final; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}
