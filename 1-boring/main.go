package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	// Once this go routine is run, the program moves on and doesn't wait for the boring func to finish.
	// This is because the main goroutine is a caller and doesn't wait.
	// Thus, we don't see anything. 
	go boring("boring!")

	/* This can be solved by calling the go routing in an infinite for loop. ie 

	  for {
			go boring("boring!")
		}
	*/

	// This code will execute and print. Because of the sleep, we will see 
	// print messages from boring as well.
  fmt.Println("I'm listenling.")
	time.Sleep(2*time.Second)
	fmt.Println("You're boring. I'm leaving.")

	// Despite using time.Sleep to see messages from boring, the the main 
	// gorouting and the boring goroutine are not communicating. Real conversation
	// requires communication. 
}
