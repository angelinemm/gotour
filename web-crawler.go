package main

import (
	"fmt"
	"sync"
)

// FetchedUrls is a map (safe with mutex) of urls 
// that have already been fetched
type FetchedUrls struct {
	v   map[string]bool
	mux sync.Mutex
}

// Add adds a url to the maps of already fetched urls
func (urls FetchedUrls) Add(url string) {
    urls.mux.Lock()
	urls.v[url] = true
	urls.mux.Unlock()
}

// AlreadyFetched checks if a url has already been fetched
func (urls FetchedUrls) AlreadyFetched(url string) bool {
    urls.mux.Lock()
	_, ok := urls.v[url]
	urls.mux.Unlock()
	return ok
}

// Fetcher returns the body of URL and
// a slice of URLs found on that page.
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, fetched FetchedUrls) {
	if depth <= 0 || fetched.AlreadyFetched(url) {
		return
	} 
	fetched.Add(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher, fetched)
	}
	return
}

func main() {
    urls := FetchedUrls{v: make(map[string]bool)}
	Crawl("http://golang.org/", 4, fetcher, urls)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
