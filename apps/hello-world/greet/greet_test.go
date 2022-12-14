package greet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_HelloAndGreet(t *testing.T) {
	t.Run("saying hello world", func(t *testing.T) {
		result := Hello()
		expected := "Hello, world!"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("saying hello to people", func(t *testing.T) {
		result := Greet("Tom", "English")
		expected := "Hello, Tom"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("saying hello to people in spanish", func(t *testing.T) {
		result := Greet("Tom", "Spanish")
		expected := "Hola, Tom"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("saying hello to people in french", func(t *testing.T) {
		result := Greet("Tom", "French")
		expected := "Bonjour, Tom"

		assertCorrectMessage(t, result, expected)
	})
}

func assertCorrectMessage(t testing.TB, actual, expected string) {
	t.Helper()

	assert.Equal(t, expected, actual)
}
