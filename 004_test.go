package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		n   int
		exp bool
	}{
		{1, true},
		{11, true},
		{121, true},
		{112, false},
		{9999, true},
		{9979, false},
		{12345678987654321, true},
	}
	for _, c := range testCases {
		assert.Equal(t, c.exp, IsPalindrome(c.n), c.n)
	}
}

func TestLargestPalindromeProduct(t *testing.T) {
	assert.Equal(t, 9009, LargestPalindromeProduct(2))
}
