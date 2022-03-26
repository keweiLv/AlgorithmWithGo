package main

import "sync"

type httppkg struct {
}

func (httppkg) Get(url string) {

}

var http httppkg

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			http.Get(url)
		}(url)
	}
	wg.Wait()
}
