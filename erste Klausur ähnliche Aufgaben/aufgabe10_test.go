package main

import "fmt"

func ExampleSumOfOddDigits() {
	fmt.Println(SumOfOddDigits(12345))
	fmt.Println(SumOfOddDigits(24680))
	fmt.Println(SumOfOddDigits(13579))
	fmt.Println(SumOfOddDigits(9081726354))

	// Output:
	// 9
	// 0
	// 25
	// 25
}
