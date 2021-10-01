package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mux sync.Mutex
	v  map[string]int
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c *SafeCounter, wg *sync.WaitGroup) {
    defer wg.Done()

    if depth <= 0 {
        return
    }

    c.mux.Lock()
    c.v[url]++
    c.mux.Unlock()

    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("found: %s %q\n", url, body)
    for _, u := range urls {
        c.mux.Lock()
        i := c.v[u]
        c.mux.Unlock()
        if i == 1 {
            continue
        }
        wg.Add(1)
        go Crawl(u, depth-1, fetcher, c, wg)
    }
    return
}

func main() {
    c := SafeCounter{v: make(map[string]int)}
    var wg sync.WaitGroup
    wg.Add(1)
    Crawl("https://golang.org/", 4, fetcher, &c, &wg)
    wg.Wait()
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