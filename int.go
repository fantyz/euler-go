package main

func IntPow(a, b int) int {
	v := 1
	for i := 0; i < b; i++ {
		v *= a
	}
	return v
}

func Digits(n int) int {
	if n < 0 {
		n = n * -1
	}

	digits := 1
	v := 10
	for {
		if v > n {
			return digits
		}
		v *= 10
		digits++
	}
}

func NthDigit(i, n int) int {
	if n < 1 {
		return 0
	}
	return (i / IntPow(10, n-1)) % 10
}
