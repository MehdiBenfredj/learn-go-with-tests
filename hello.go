package main

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
