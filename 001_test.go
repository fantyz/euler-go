package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumDivisbleBy(t *testing.T) {
	testCases := map[string]struct {
		N     int
		Below int
		Exp   int
	}{
		"n=3, below=10": {3, 10, 18},
		"n=4, below=10": {4, 10, 12},
		"n=5, below=10": {5, 10, 5},
	}

	for n, c := range testCases {
		assert.Equal(t, c.Exp, SumDivisibleBy(c.N, c.Below), n)
	}
}

func TestSumDivisbleByMultiplesOf3And5(t *testing.T) {
	testCases := map[string]struct {
		Below int
		Exp   int
	}{
		"below=10": {10, 23},
	}

	for n, c := range testCases {
		assert.Equal(t, c.Exp, MultiplesOf3And5(c.Below), n)
	}

}
