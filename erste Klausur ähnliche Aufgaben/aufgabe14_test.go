package main

import "fmt"

func ExampleCountVowels() {
	fmt.Println(CountVowels("hello world"))
	fmt.Println(CountVowels("abcdefg"))
	fmt.Println(CountVowels("aeiouAEIOU"))
	fmt.Println(CountVowels("xyz"))

	// Output:
	// 3
	// 2
	// 10
	// 0
}
