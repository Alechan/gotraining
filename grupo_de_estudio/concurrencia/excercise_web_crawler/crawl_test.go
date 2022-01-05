package main

import (
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

func TestCrawl_CreatesGoRoutines(t *testing.T) {
	timeOut := 500 * time.Millisecond

	fetcher := &fakeFetcherConcurrentCounter{
		firstCallFlag: false,
		wg:            sync.WaitGroup{},
		urlNodes:      urlNodesWithoutRecursion,
	}

	// Set the wait group to wait for 3 (the test expects 3 concurrent calls)
	fetcher.wg.Add(3)

	t.Run("Should make fetches concurrently", func(t *testing.T) {
		// Call Crawl in a goroutine and wait for it to finish
		// If it doesn't finish, it didn't create 3 go routines and the one(s) created
		// are blocked in the waitgroup forever.
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

func TestCrawl_NoDuplicateFetches(t *testing.T) {
	fetcher := &fakeFetcherConcurrentNoDuplicates{
		urlsCounter: map[string]int{},
		urlNodes:    urlNodesWithRecursion,
	}

	t.Run("Should not make duplicate fetches", func(t *testing.T) {
		Crawl("https://root/", 3, fetcher)

		actualUrlsCounter := fetcher.urlsCounter
		require.Equal(t, 1, actualUrlsCounter["https://root/"])
		require.Equal(t, 1, actualUrlsCounter["https://root/a"])
		require.Equal(t, 1, actualUrlsCounter["https://root/b"])
		require.Equal(t, 1, actualUrlsCounter["https://root/c"])
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

var urlNodesWithRecursion = map[string][]string{
	"https://root/": {
		"https://root/a",
		"https://root/b",
		"https://root/c",
	},
	"https://root/a": {
		"https://root/",
	},
	"https://root/b": {
		"https://root/",
	},
	"https://root/c": {
		"https://root/",
	},
}
