package main

import (
	"fmt"
	"math/rand"
	"time"
)

//
type Message struct {
	str  string
	wait chan bool
}

// the boring function retuns a 'receive only' channel to communicate with it
func boring(msg string) <-chan Message {
	c := make(chan Message)
	// shared between all Messages
	waitForIt := make(chan bool)
	// spawn go routine inside function
	go func() {
		for i := 0; ; i++ {
			c <- Message{
				str:  fmt.Sprintf("%s %d", msg, i),
				wait: waitForIt,
			}

			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			// everytime the go routine sends a message, this line blocks until a value has been received
			<-waitForIt
		}

		// close(c)
	}()

	// return a channel to the caller
	return c
}

func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i]
		go func() {
			for {
				c <- <-input
			}

		}()
	}

	return c
}

func main() {
	// merge 2 channels into 1 channel
	c := fanIn(boring("Joe"), boring("Sue"))

	for i := 0; i < 15; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)

		// allows boring go routine to send next value to message channel	
		msg1.wait <- true
		msg2.wait <- true
	}

	fmt.Println("You're boring. I'm leaving.")
}
