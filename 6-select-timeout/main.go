package main

import (
	"fmt"
	"math/rand"
	"time"
)

// the boring function retuns a 'receive only' channel to communicate with it
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		}
	}()

	return c
}

func main() {
	c := boring("Joe")
	// timeout limits time select will block for.
	timeout := time.After(5 * time.Second)

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}
