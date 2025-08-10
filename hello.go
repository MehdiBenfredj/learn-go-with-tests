package main

const englishHelloPrefix = "Hello, "

func Hello(n string) string {
	if n == "" {
		n = "World"
	}
	return englishHelloPrefix + n
}
