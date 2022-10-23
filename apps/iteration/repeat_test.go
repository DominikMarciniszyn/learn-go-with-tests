package iteration

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a", 5)
	expected := "aaaaa"

	assert.Equal(t, expected, result)
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)

	// Output: aaaaa
}

//func BenchmarkRepeat(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Repeat("a")
//	}
//}
