package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
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
