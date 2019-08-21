package main_test

import (
	. "coding"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounterNameInputListOfNameShouldReturnNewListOfName(t *testing.T) {
	expectedListOfName := []string{"me", "you", "me1", "they", "we", "they1"}

	actualListOfName := CounterName([]string{"me", "you", "me", "they", "we", "they"})

	assert.Equal(t, expectedListOfName, actualListOfName)
}
