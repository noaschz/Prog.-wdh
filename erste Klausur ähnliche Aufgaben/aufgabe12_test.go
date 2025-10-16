package main

import "fmt"

func ExampleReverseNumber() {
	fmt.Println(ReverseNumber(1234))
	fmt.Println(ReverseNumber(98765))
	fmt.Println(ReverseNumber(100))
	fmt.Println(ReverseNumber(7))

	// Output:
	// 4321
	// 56789
	// 1
	// 7
}
