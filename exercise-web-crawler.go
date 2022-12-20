package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Counter struct {
	mu sync.Mutex
	v int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v++
}

func (c *Counter) Dec() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v--
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}

type CacheEntry struct {
	mu sync.Mutex
	done bool
	body string
	urls []string
	err error
}

type Cache struct {
	mu sync.Mutex
	v map[string]*CacheEntry
}

func (cache *Cache) Get(url string) *CacheEntry {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if _, ok := cache.v[url]; ok == false {
		cache.v[url] = &CacheEntry{}
	}
	return cache.v[url]
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cnt *Counter, cache *Cache) {
	cnt.Inc()

	go func() {
		defer cnt.Dec()
		if depth <= 0 {
			return
		}
		entry := cache.Get(url)
		entry.mu.Lock()
		if entry.done == false {
			entry.body, entry.urls, entry.err = fetcher.Fetch(url)
			entry.done = true
		}
		entry.mu.Unlock()
		body, urls, err := entry.body, entry.urls, entry.err
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			Crawl(u, depth-1, fetcher, cnt, cache)
		}
	}()
}

func main() {
	cnt := &Counter{}
	cache := &Cache{v: map[string]*CacheEntry{}}
	Crawl("https://golang.org/", 4, fetcher, cnt, cache)
	for cnt.Value() != 0 {
		time.Sleep(time.Millisecond)
	}
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
