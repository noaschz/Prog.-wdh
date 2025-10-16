package aufgabe2

import "fmt"

func ExampleCountAbc() {
	l1 := []string{"abcd", "ab", "de", "abcde", "defabcedfg", "kj"}
	l2 := []string{"abcde", "ab", "de", "abc", "defabcedfg", "kj"}
	l3 := []string{"dabc", "ab", "de", "cabcde", "defabcedfg", "kj"}
	l4 := []string{"abc"}

	fmt.Println(CountAbc(l1))
	fmt.Println(CountAbc(l2))
	fmt.Println(CountAbc(l3))
	fmt.Println(CountAbc(l4))

	// Output:
	// 2
	// 2
	// 0
	// 1
}