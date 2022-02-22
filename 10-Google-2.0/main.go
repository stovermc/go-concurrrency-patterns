package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result 

var (
	Web = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// Google uses goroutines to concurrently issue searches. The fan in pattern 
// is used to collect all the results from a single channel.
func Google(query string) (results []Result) {
	c := make(chan Result)

	go func() {
		c <- Web(query)
	}()

	go func() {
		c <- Image(query)
	}()

	go func() {
		c <- Video(query)
	}()

  for i := 0; i < 3; i++ {
		results = append(results, <-c)
	}

	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}