package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntPow(t *testing.T) {
	testCases := map[string]struct {
		a, b int
		exp  int
	}{
		"a=2,b=4": {2, 4, 16},
		"a=1,b=4": {1, 4, 1},
		"a=2,b=1": {2, 1, 2},
		"a=0,b=4": {0, 4, 0},
		"a=2,b=0": {2, 0, 1},
		"a=0,b=0": {0, 0, 1},
	}

	for name, c := range testCases {
		assert.Equal(t, c.exp, IntPow(c.a, c.b), name)
	}
}

func TestDigits(t *testing.T) {
	testCases := map[string]struct {
		n   int
		exp int
	}{
		"0":     {0, 1},
		"1":     {1, 1},
		"9":     {9, 1},
		"49":    {49, 2},
		"9958":  {9958, 4},
		"-9958": {-9958, 4},
	}

	for name, c := range testCases {
		assert.Equal(t, c.exp, Digits(c.n), name)
	}
}

func TestNthDigit(t *testing.T) {
	testCases := map[string]struct {
		i, n int
		exp  int
	}{
		"i=49,n=-1": {49, -1, 0},
		"i=49,n=0":  {49, 0, 0},
		"i=49,n=1":  {49, 1, 9},
		"i=49,n=2":  {49, 2, 4},
		"i=49,n=3":  {49, 3, 0},
	}

	for name, c := range testCases {
		assert.Equal(t, c.exp, NthDigit(c.i, c.n), name)
	}
}
