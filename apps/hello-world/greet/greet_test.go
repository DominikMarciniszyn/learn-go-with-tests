package greet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Hello(t *testing.T) {
	result := Hello()
	expected := "Hello, world!"

	assert.Equal(t, expected, result)
}

func Test_Greet(t *testing.T) {
	result := Greet("Tom")
	expected := "Hello, Tom"

	assert.Equal(t, expected, result)
}
