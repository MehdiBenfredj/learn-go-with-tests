package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mehdibenfredj/learn-go-with-tests/mocking"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	mocking.Countdown(os.Stdout, &mocking.DefaultSleeper{})
}

func Hello(n string, l string) string {
	var helloPrefixDict = map[string]string{
		"Spanish": "Hola, ",
		"English": "Hello, ",
		"French":  "Salut, ",
	}

	var worldDict = map[string]string{
		"Spanish": "Mundo",
		"English": "World",
		"French":  "Monde",
	}

	_, ok := helloPrefixDict[l]

	if !ok || l == "" {
		l = "English"
	}

	if n == "" {
		n = worldDict[l]
	}

	return helloPrefixDict[l] + n
}
