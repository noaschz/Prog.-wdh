package reku1

import "fmt"

func ExampleRecursiveShortestAbc() {
	l1 := []string{"abcd", "ab", "de", "abcde", "defabcedfg", "kj"}
	l2 := []string{"abcde", "ab", "de", "abc", "defabcedfg", "kj"}
	l3 := []string{"dabc", "ab", "de", "cabcde", "defabcedfg", "kj"}
	l4 := []string{"abc"}

	fmt.Println(RecursiveShortestAbc(l1))
	fmt.Println(RecursiveShortestAbc(l2))
	fmt.Println(RecursiveShortestAbc(l3))
	fmt.Println(RecursiveShortestAbc(l4))

	// Output:
	// abcd
	// abc
	//
	// abc
}