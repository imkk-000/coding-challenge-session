package main_test

import (
	. "coding"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDimensionInputListShouldReturn0(t *testing.T) {
	expectedDimension := 0
	inputMatrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	actualDimension := FindDimension(inputMatrix)

	assert.Equal(t, expectedDimension, actualDimension)
}
