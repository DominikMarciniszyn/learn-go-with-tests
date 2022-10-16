package greet

import "fmt"

const englishHelloPrefix = "Hello,"

func Hello() string {
	return fmt.Sprintf("%s world!", englishHelloPrefix)
}

func Greet(name string) string {
	return fmt.Sprintf("%s %s", englishHelloPrefix, name)
}
