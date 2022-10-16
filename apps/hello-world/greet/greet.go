package greet

import "fmt"

func Hello() string {
	return "Hello, world!"
}

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
