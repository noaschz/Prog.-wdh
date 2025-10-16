package aufgabe3

import "fmt"

func ExampleAllAbc() {
	l1 := []string{"abcd", "ab", "de", "abcde", "defabcedfg", "kj"}
	l2 := []string{"abcde", "ab", "de", "abc", "defabcedfg", "kj"}
	l3 := []string{"dabc", "ab", "de", "cabcde", "defabcedfg", "kj"}
	l4 := []string{"abc"}

	fmt.Println(AllAbc(l1))
	fmt.Println(AllAbc(l2))
	fmt.Println(AllAbc(l3))
	fmt.Println(AllAbc(l4))

	// Output:
	// [abcd abcde]
	// [abcde abc]
	// []
	// [abc]
}