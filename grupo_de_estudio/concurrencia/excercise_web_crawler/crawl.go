package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var fetchedUrlsSync = struct {
	fetchedUrls map[string]bool
	sync.Mutex
}{fetchedUrls: map[string]bool{}}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	if isDuplicateUrl(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	crawlConcurrently(urls, depth, fetcher)
	return
}

func crawlConcurrently(urls []string, depth int, fetcher Fetcher) {
	wg := sync.WaitGroup{}
	for _, u := range urls {
		wg.Add(1)
		go func(nestedUrl string) {
			Crawl(nestedUrl, depth-1, fetcher)
			wg.Done()
		}(u)
	}
	wg.Wait()
}

func isDuplicateUrl(url string) bool {
	// Get mutex to see if this url has already been fetched
	fetchedUrlsSync.Mutex.Lock()
	{
		if _, in := fetchedUrlsSync.fetchedUrls[url]; in {
			fmt.Printf("Skipped url %v because it has already been fetched\n", url)
			fetchedUrlsSync.Mutex.Unlock()
			return true
		}
		fetchedUrlsSync.fetchedUrls[url] = true
	}
	fetchedUrlsSync.Mutex.Unlock()
	return false
}
