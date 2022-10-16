package operations

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 4)
	expected := 6

	assert.Equal(t, expected, sum)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)

	// Output: 6
}
