package main

import (
	"sync"
	"testing"
	"time"
)

func TestCrawl_CreatesGoRoutines(t *testing.T) {
	timeOut := 100 * time.Millisecond

	fetcher := &fakeFetcherConcurrentCounter{
		firstCallPassedFlag: false,
		wg:                  sync.WaitGroup{},
		urlNodes:            urlNodesWithoutRecursion,
	}

	// Set the wait group to wait for 3 (the test expects 3 concurrent calls)
	fetcher.wg.Add(3)

	t.Run("Should make fetches concurrently", func(t *testing.T) {
		ch := make(chan struct{})
		go func() {
			Crawl("https://root/", 2, fetcher)
			ch <- struct{}{}
		}()

		select {
		case <-ch:
			// Test passed!
		case <-time.After(timeOut):
			t.Error("The crawler didn't end in time. Are the calls made concurrently?")
		}
	})
}

var urlNodesWithoutRecursion = map[string][]string{
	"https://root/": {
		"https://root/a",
		"https://root/b",
		"https://root/c",
	},
	"https://root/a": {},
	"https://root/b": {},
	"https://root/c": {},
}
