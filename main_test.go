package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlgo(t *testing.T) {
	a, b := 1, 2
	assert.Equal(t, 3, a + b)
}