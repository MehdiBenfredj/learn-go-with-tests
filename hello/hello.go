package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	math_test "github.com/mehdibenfredj/learn-go-with-tests/math"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {

	for range 30000 {
		//fmt.Printf("writing %d", now.Second())
		file, err := os.OpenFile("/Users/mehdi/code/go/learn-go-with-tests/math/test_clock.svg", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Errorf("Could not open file", err.Error())
		}
		defer file.Close()
		math_test.WriteTime(file, time.Now())
		time.Sleep(time.Second)

	}
	//<-time.After(30 * time.Second)
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
