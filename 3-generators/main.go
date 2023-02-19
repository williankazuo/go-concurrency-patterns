package main

import "fmt"

func main() {
	receivedValues := generateNumbers(1, 400)
	for v := range receivedValues {
		fmt.Printf("Value: %v\n", v)
	}
}

func generateNumbers(start, final int) <-chan int {
	out := make(chan int)
	go func() {
		for i := start; i < final; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}
