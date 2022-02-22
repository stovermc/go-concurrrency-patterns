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
func Google(query string) []Result {
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

	var results []Result

	// ignore results from server that take longer than 50ms
	timeout := time.After(50 * time.Millisecond)
  for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			results = append(results, r)

		case <-timeout:
			fmt.Println("timeout")
			return results
		}
	}

	return results
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)

	searchReplica := func(i int) {
		c <- replicas[i](query)
	}

	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}