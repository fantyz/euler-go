package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPrimes(t *testing.T) {
	expSeq := []int{2, 3, 5, 7}

	primes := NewPrimes()
	seq := make([]int, 0, len(expSeq))

	for i := 0; i < len(expSeq); i++ {
		seq = append(seq, primes.Next())
	}

	assert.Equal(t, expSeq, seq)
}

func TestNaivePrimes(t *testing.T) {
	expSeq := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}

	primes := &naivePrimes{n: 1}

	for i := 0; i < 2; i++ {
		seq := make([]int, 0, len(expSeq))

		for j := 0; j < len(expSeq); j++ {
			seq = append(seq, primes.Next())
		}

		assert.Equal(t, expSeq, seq)

		primes.Restart()
	}

}

func TestPrimeFactors(t *testing.T) {
	testCases := map[string]struct {
		n          int
		expFactors []int
	}{
		"2":     {2, []int{2}},
		"9":     {9, []int{3, 3}},
		"13195": {13195, []int{5, 7, 13, 29}},
	}

	for name, c := range testCases {
		assert.Equal(t, c.expFactors, PrimeFactors(c.n), name)
	}
}

func benchmarkNaivePrimes(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		primes := &naivePrimes{n: 1}
		for j := 0; j < n; j++ {
			primes.Next()
		}
	}
}

func BenchmarkNaivePrimes10(b *testing.B)    { benchmarkNaivePrimes(10, b) }
func BenchmarkNaivePrimes100(b *testing.B)   { benchmarkNaivePrimes(100, b) }
func BenchmarkNaivePrimes1000(b *testing.B)  { benchmarkNaivePrimes(1000, b) }
func BenchmarkNaivePrimes10000(b *testing.B) { benchmarkNaivePrimes(10000, b) }
