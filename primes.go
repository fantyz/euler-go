package main

func NewPrimes() *naivePrimes {
	return &naivePrimes{n: 1}
}

type naivePrimes struct {
	n      int
	primes []int
}

func (p *naivePrimes) Next() int {
	// Naive and prime generation
	for {
		isPrime := true
		p.n++

		for _, i := range p.primes {
			if p.n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			p.primes = append(p.primes, p.n)
			return p.n
		}
	}
}

func (p *naivePrimes) Restart() {
	p.n = 1
	p.primes = p.primes[:0]
}

func PrimeFactors(n int) []int {
	var factors []int
	primes := NewPrimes()
	for {
		if n == 1 {
			// done
			return factors
		}
		primes.Restart()
		for {
			p := primes.Next()
			if p == n || n%p == 0 {
				factors = append(factors, p)
				n = n / p
				break
			}
		}
	}
}
