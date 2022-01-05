package main

import (
	"sync"
)

type fakeFetcherConcurrentCounter struct {
	urlNodes map[string][]string

	// For fist call
	firstCallFlag  bool
	firstCallMutex sync.Mutex

	// For nested calls
	wg sync.WaitGroup
}

func (f *fakeFetcherConcurrentCounter) Fetch(url string) (string, []string, error) {
	f.firstCallMutex.Lock()
	{
		if !f.firstCallFlag {
			// The first call is not required to be run in parallel so don't wait for waitgroup
			f.firstCallFlag = true
			f.firstCallMutex.Unlock()
			return "", f.urlNodes[url], nil
		}
	}
	f.firstCallMutex.Unlock()

	// We are not in the first call, so wait until all other sub-urls have been called
	f.wg.Done()
	f.wg.Wait()

	return "", f.urlNodes[url], nil
}
