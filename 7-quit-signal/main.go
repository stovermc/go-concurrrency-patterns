package main

import (
	"fmt"
	"math/rand"
	"time"
)

// the boring function retuns a 'receive only' channel to communicate with it
func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			for {
				select {
				case c <- fmt.Sprintf("%s %d", msg, i):
					// do nothing
				case <-quit:
					fmt.Println("Clean up.")
					quit <- "see you later!"
					return
				}

				time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			}
		}
	}()

	return c
}

func main() {
	quit := make(chan string)
	c := boring("Joe", quit)

	for i := 3; i >= 0; i-- {
		fmt.Println(<-c)
	}

	quit <- "Bye"
	fmt.Println("Joe said:", <-quit)

}
