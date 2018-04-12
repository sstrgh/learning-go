package main

import (
	"fmt"
	"strings"
)

const helloPrefixEnglish = "Hello, "
const helloPrefixSpanish = "Hola!, "
const helloPrefixFrench = "Bonjour, "

func Hello(name string, language string) string {
	language = strings.ToLower(language)

	if name == "" {
		name = "World"
	}

	switch language {
	case "french":
		return frenchHello(name)
	case "spanish":
		return spanishHello(name)
	default:
		return englishHello(name)
	}
}

func spanishHello(name string) string {
	return helloPrefixSpanish + name
}

func frenchHello(name string) string {
	return helloPrefixFrench + name
}

func englishHello(name string) string {
	return helloPrefixEnglish + name
}

func main() {
	fmt.Println(Hello("Writer", ""))
}
