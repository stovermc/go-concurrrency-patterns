package main

import (
	"fmt"
	"math/rand"
	"time"
)

// pass in an additional parameter that is a channel.
func boring(msg string, c chan string) {
	for i := 0; ; i++ {

		// send value to channel
		// channel waits for reciever to be ready
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	// init our channel
	c := make(chan string)

	go boring("boring!", c)

	for i := 0; i < 5; i++ {
		// <-c read the value from boring function
		// <-c blocks and waits for a value to be sent
		fmt.Printf("You say: %q\n", <-c)
	}

	fmt.Println("You're boring. I'm leaving.")
}
