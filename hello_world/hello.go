package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englighHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englighHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func Add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(Hello("World", "English"))
}
