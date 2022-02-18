package main

import (
	"fmt"
	"math/rand"
	"time"
)

// the boring function retuns a 'receive only' channel to communicate with it
func boring(msg string) <-chan string {
	c := make(chan string)

	// spawn go routine inside function 
	go func() {

		// this for loop simulates an infinite sender
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

		// close(c)
	}()

	// return a channel to the caller
	return c
}

func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)

	go func(){
		// infinite loop to read value from channel
		for {
			// read value from c1. This line will block until value is received
			v1 := <-c1
			c <- v1
		}

	}()
	go func(){
    for {
			// read value from c2 and post value to c channel
			c <- <-c2
		}
	}()

	return c
}

func main() {
	// merge 2 channels into 1 channel
	c := fanIn(boring("Joe"), boring("Sue"))

	for i := 0; i < 15; i++ {
		// now we can read from 1 channel
		fmt.Println(<-c)
	}

	fmt.Println("You're boring. I'm leaving.")
}
