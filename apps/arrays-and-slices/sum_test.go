package arrays_and_slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		result := Sum(numbers)
		expected := 15

		assert.Equal(t, expected, result)
	})
}
func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	assert.Equal(t, expected, result)
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		result := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		assert.Equal(t, expected, result)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		result := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		assert.Equal(t, expected, result)
	})
}
