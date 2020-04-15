package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	expSeq := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

	gen := NewFibonacci()
	seq := make([]int, 0, len(expSeq))

	for i := 0; i < len(expSeq); i++ {
		seq = append(seq, gen.Next())
	}

	assert.Equal(t, expSeq, seq)
}
