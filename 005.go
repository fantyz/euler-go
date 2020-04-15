package main

import "fmt"

/*

Smallest multiple

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

*/

func IsEvenlyDivisible(max, n int) bool {
	for i := max; i > 0; i-- {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func SmallestEvenlyDivisible(max int) int {
	i := 0
	for {
		i += max
		if IsEvenlyDivisible(max, i) {
			return i
		}
	}
}

func Problem005() {
	fmt.Println("Problem 5 - The smallest positive number that is evenly divisible by all of the numbers from 1 to 20:", SmallestEvenlyDivisible(20))
}
