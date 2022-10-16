package greet

import "fmt"

const (
	Spanish = "Spanish"
	French  = "French"
)

const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const frenchHelloPrefix = "Bonjour,"

func Hello() string {
	return fmt.Sprintf("%s world!", englishHelloPrefix)
}

func Greet(name string, language string) string {
	return fmt.Sprintf("%s %s", greetPrefix(language), name)
}

func greetPrefix(language string) (prefix string) {
	switch language {
	case French:
		prefix = frenchHelloPrefix
	case Spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
