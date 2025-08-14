package concurency

import "sync"

type WebsiteChecker func(string) bool

func CheckWebsites1(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	wg := sync.WaitGroup{}
	wg.Add(len(urls))
	mu := sync.RWMutex{}

	for _, url := range urls {
		go func() {
			defer wg.Done()
			res := wc(url)
			mu.Lock()
			results[url] = res
			mu.Unlock()
		}()
	}

	wg.Wait()
	return results
}

type Tuple struct {
	url string
	res bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	wg := sync.WaitGroup{}
	ch := make(chan Tuple, len(urls))

	wg.Add(len(urls))
	t := &Tuple{}
	for _, url := range urls {
		go func() {
			defer wg.Done()
			t.url = url
			t.res = wc(url)
			ch <- *t
		}()
	}
	defer close(ch)

	go func() {
		for r := range ch {
			results[r.url] = r.res
		}
	}()
	wg.Wait()
	return results
}
