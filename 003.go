package main

import "fmt"

/*

Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/

func Problem003() {
	factors := PrimeFactors(600851475143)
	max := 0
	for _, f := range factors {
		if f > max {
			max = f
		}
	}
	fmt.Println("Problem 3 - Largest prime factor of the number 600851475143:", max)
}
