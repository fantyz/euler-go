package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEvenlyDivisible(t *testing.T) {
	testCases := []struct {
		max int
		n   int
		exp bool
	}{
		{10, 2520, true},
		{10, 2530, false},
		{3, 6, true},
	}
	for _, c := range testCases {
		assert.Equal(t, c.exp, IsEvenlyDivisible(c.max, c.n), "n=%d,max=%d", c.n, c.max)
	}
}

func TestSmallestEvenlyDivisible(t *testing.T) {
	assert.Equal(t, 2520, SmallestEvenlyDivisible(10))
}
