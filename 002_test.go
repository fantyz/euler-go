package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumEvenFibonacciNumbers(t *testing.T) {
	testCases := map[string]struct {
		Below int
		Exp   int
	}{
		"below=10": {10, 10},
	}

	for n, c := range testCases {
		assert.Equal(t, c.Exp, SumEvenFibonacciNumbers(c.Below), n)
	}

}
