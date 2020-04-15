package main

import "fmt"

/*

Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

*/

func IsPalindrome(n int) bool {
	i, j := 1, Digits(n)
	for i < j {
		if NthDigit(n, i) != NthDigit(n, j) {
			return false
		}
		i++
		j--
	}
	return true
}

func LargestPalindromeProduct(digits int) int {
	max := IntPow(10, digits) - 1

	largest := 0
	for i := max; i > 0; i-- {
		for j := max; j > 0; j-- {
			n := i * j
			if IsPalindrome(n) {
				if n > largest {
					largest = n
				}
			}
		}
	}

	return largest
}

func Problem004() {
	fmt.Println("Problem 4 - The largest palindrome made from the product of 3-digit numbers:", LargestPalindromeProduct(3))
}
