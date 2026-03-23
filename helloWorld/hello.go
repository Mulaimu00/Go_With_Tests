package hello

import (
	"fmt"
)

const (
	spanish = "Spanish"
	swahili = "Swahili"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	swahiliHelloPrefix = "Jambo, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return  greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case swahili:
		prefix = swahiliHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func main() {
	greeting := Hello("Mesh","Swahili")
	fmt.Println(greeting)
}