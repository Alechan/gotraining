package main

import (
	"sync"
)

type fakeFetcherConcurrentCounter struct {
	urlNodes map[string][]string

	// For fist call
	firstCallPassedFlag bool
	firstCallMutex      sync.Mutex

	// For nested calls
	wg sync.WaitGroup
}

func (f *fakeFetcherConcurrentCounter) Fetch(url string) (string, []string, error) {
	// Check if this is the first call to fetch (that shouldn't be made concurrently)
	f.firstCallMutex.Lock()
	{
		if !f.firstCallPassedFlag {
			f.firstCallPassedFlag = true
			nextUrls := f.urlNodes[url]

			return "", nextUrls, nil
		}

	}
	f.firstCallMutex.Unlock()

	// We are not in the first call, so wait until all other sub-urls have been called
	f.wg.Done()
	f.wg.Wait()

	nextUrls := f.urlNodes[url]

	return "", nextUrls, nil
}
