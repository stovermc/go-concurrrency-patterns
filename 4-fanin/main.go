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

// fanIn combines values from c1 and c2 into a third channel c.
// order of values in c is dependent on the order in which values are received from c1 and c2
func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			// read value from c1
			c <- <-c1
		}

	}()
	go func() {
		for {
			// read value from c2
			c <- <-c2
		}
	}()

	return c
}

func fanInSelect(c1, c2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-c1:
				c <- s
			case s := <-c2:
				c <- s
			}
		}
	}()

	return c
}

func main() {
	// merge 2 channels into 1 channel
	// c := fanIn(boring("Joe"), boring("Sue"))
	c := fanInSelect(boring("Joe"), boring("Sue"))

	for i := 0; i < 15; i++ {
		// now we can read from 1 channel
		fmt.Println(<-c)
	}

	fmt.Println("You're boring. I'm leaving.")
}
