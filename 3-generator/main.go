package main

import (
	"fmt"
	"math/rand"
	"time"
)

// now the boring function retuns a 'receive only' channel.
// this will allow for communication with the boring function
func boring(msg string) <-chan string {
	// init our channel
	c := make(chan string)

	// spawn go routine inside function that sends data to our channel
	go func() {

		// this for loop simulates an infinite sender
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

		// the sender should close the channel
		close(c)
	}()

	return c
}

func main() {

	joe := boring("Joe")
	sue := boring("Sue")

	// this loop prints out the values from each channel in sequence. note: <-joe will always block before moving on to sue
	for i := 0; i < 10; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-sue)
	}

	fmt.Println("You're boring. I'm leaving.")
}
