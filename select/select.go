package select_package

import (
	"errors"
	"net/http"
	"time"
)

func RacerWithoutSelectThatCouldAlsoTakeMultipleURLS(urls ...string) string {
	fastest := make(chan string)

	for _, url := range urls {
		go func() {
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			fastest <- url
		}()
	}

	return <-fastest
}

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", errors.New("Timeout")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
