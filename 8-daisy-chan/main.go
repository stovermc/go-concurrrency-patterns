package main

import "fmt"

func daisy(left, right chan int) {
	// get the value from right and add 1 to it
	left <- 1+ <-right
}

func main() {
	const n = 1000

	leftmost := make(chan int)

	left := leftmost
	right := leftmost

	for i := 0; i < n; i++ {
		go daisy(left, right)
		left = right
	}

	go func (c chan int) {
		c <- 1
	}(right)

	fmt.Println(<-leftmost)
}