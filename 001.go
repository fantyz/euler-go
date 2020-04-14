package main

import "fmt"

/*

Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

*/

func SumDivisibleBy(n, max int) int {
	sum := 0
	for i := n; i < max; i += n {
		sum += i
	}
	return sum
}

func MultiplesOf3And5(below int) int {
	return SumDivisibleBy(3, below) + SumDivisibleBy(5, below) - SumDivisibleBy(15, below)
}

func Problem001() {
	fmt.Println("Problem 1 - Sum of all multiples of 3 and 5 below 1000:", MultiplesOf3And5(1000))
}
