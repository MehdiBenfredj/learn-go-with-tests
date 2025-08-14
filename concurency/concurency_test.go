package concurency

import (
	"reflect"
	"testing"
	"time"
)

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func TestCheckWebsites(t *testing.T) {
	t.Run("testing with mock", func(t *testing.T) {
		want := map[string]bool{
			"google.com": false,
			"mehdi.com":  true,
			"test.fr":    false,
		}

		got := CheckWebsites(mockWebSiteChecker, []string{"google.com",
			"mehdi.com",
			"test.fr"})

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("wanted %v, got %v", want, got)
		}

	})
}

var m = map[string]bool{
	"google.com": false,
	"mehdi.com":  true,
	"test.fr":    false,
}

func mockWebSiteChecker(url string) bool {
	return m[url]
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
