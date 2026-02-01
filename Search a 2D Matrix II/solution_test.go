package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	testCases := []struct {
		target   int
		expected bool
	}{
		{5, true},
		{14, true},
		{20, false},
		{1, true},
		{30, true},
		{0, false},
		{31, false},
	}

	for _, tc := range testCases {
		result := searchMatrix(matrix, tc.target)
		assert.Equal(t, tc.expected, result, "searchMatrix(matrix, %d)", tc.target)
	}
}

func TestSearchMatrixEmptyMatrix(t *testing.T) {
	// Test empty matrix
	emptyMatrix := [][]int{}
	result := searchMatrix(emptyMatrix, 1)
	assert.False(t, result, "searchMatrix should return false for empty matrix")
}

func TestSearchMatrixSingleElement(t *testing.T) {
	// Test single element matrix
	singleMatrix := [][]int{{1}}

	result := searchMatrix(singleMatrix, 1)
	assert.True(t, result, "searchMatrix should find existing element")

	result = searchMatrix(singleMatrix, 2)
	assert.False(t, result, "searchMatrix should return false for non-existing element")
}
