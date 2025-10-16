package main

import "fmt"

func ExampleSumOfEvenDigits() {
	fmt.Println(SumOfEvenDigits(123456))
	fmt.Println(SumOfEvenDigits(13579))
	fmt.Println(SumOfEvenDigits(2468))

	// Output:
	// 12
	// 0
	// 20
}
