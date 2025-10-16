package aufgabe4

import "fmt"

func ExampleShortestAbcWithIndex() {
	l1 := []string{"abcd", "ab", "de", "abcde", "defabcedfg", "kj"}
	l2 := []string{"abcde", "ab", "de", "abc", "defabcedfg", "kj"}
	l3 := []string{"dabc", "ab", "de", "cabcde", "defabcedfg", "kj"}
	l4 := []string{"abc"}

	fmt.Println(ShortestAbcWithIndex(l1))
	fmt.Println(ShortestAbcWithIndex(l2))
	fmt.Println(ShortestAbcWithIndex(l3))
	fmt.Println(ShortestAbcWithIndex(l4))

	// Output:
	// abcd 0
	// abc 3
	//  -1
	// abc 0
}