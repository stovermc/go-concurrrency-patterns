package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result 

var (
	Web1 = fakeSearch("web1")
	Web2 = fakeSearch("web1")
	Image1 = fakeSearch("image1")
	Image2 = fakeSearch("image2")
	Video1 = fakeSearch("video1")
	Video2 = fakeSearch("video2")
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
		c <- First(query, Web1, Web2)
	}()

	go func() {
		c <- First(query, Image1, Image2)
	}()

	go func() {
		c <- First(query, Video1, Video2)
	}()

	var results []Result

	// ignore results from server that take longer than 50ms
	timeout := time.After(50 * time.Millisecond)
  for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)

		case <-timeout:
			fmt.Println("timeout")
			return results
		}
	}

	return results
}

// First issues multiple requests to the same service but only returns the first results to come back. This adds replication.
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