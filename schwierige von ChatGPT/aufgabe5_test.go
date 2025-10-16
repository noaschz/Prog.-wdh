package main

import "fmt"

func ExampleIsPalindromeRecursive_standard_cases() {
	fmt.Println(IsPalindromeRecursive("racecar"))
	fmt.Println(IsPalindromeRecursive("hello"))
	fmt.Println(IsPalindromeRecursive("a"))
	fmt.Println(IsPalindromeRecursive(""))

	// Output:
	// true
	// false
	// true
	// true
}
