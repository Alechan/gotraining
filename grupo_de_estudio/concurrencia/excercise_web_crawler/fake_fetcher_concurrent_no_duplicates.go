package main

import (
	"sync"
)

type fakeFetcherConcurrentNoDuplicates struct {
	urlNodes map[string][]string

	urlsCounterMutex sync.Mutex
	urlsCounter      map[string]int
}

func (f *fakeFetcherConcurrentNoDuplicates) Fetch(url string) (string, []string, error) {
	// Add 1 to the urls counter for asserts in test
	f.urlsCounterMutex.Lock()
	{
		if _, in := f.urlsCounter[url]; !in {
			f.urlsCounter[url] = 0
		}
		f.urlsCounter[url]++
	}
	f.urlsCounterMutex.Unlock()

	return "", f.urlNodes[url], nil
}
