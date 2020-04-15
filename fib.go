package main

// Fibonacci generates the fibonacci sequence starting with 1, 2, 3, 5, ... (as opposed to 1, 1, 2 ...)
type Fibonacci struct {
	a, b int
}

func NewFibonacci() *Fibonacci {
	return &Fibonacci{0, 1}
}

func (f *Fibonacci) Next() int {
	f.a, f.b = f.b, f.a+f.b
	return f.b
}
