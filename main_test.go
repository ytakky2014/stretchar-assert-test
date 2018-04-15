package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	actual := sum(10, 20)
	expected := 50

	assert.Equal(t, actual, expected, "got %v\nwant %v", actual, expected)
}
