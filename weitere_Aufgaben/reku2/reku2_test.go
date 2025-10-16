package reku2

import "fmt"

func ExampleRecursiveExcludeStringsBetween() {
	l1 := []string{"abc", "BEGIN", "def", "END", "ghi"}
	l2 := []string{"BEGIN", "def", "END", "ghi"}
	l3 := []string{"BEGIN", "def", "ghi", "END"}
	l4 := []string{"END", "def", "BEGIN", "ghi"}
	l5 := []string{"abc", "def", "ghi"}

	fmt.Println(RecursiveExcludeStringsBetween(l1, "BEGIN", "END"))
	fmt.Println(RecursiveExcludeStringsBetween(l2, "BEGIN", "END"))
	fmt.Println(RecursiveExcludeStringsBetween(l3, "BEGIN", "END"))
	fmt.Println(RecursiveExcludeStringsBetween(l4, "BEGIN", "END"))
	fmt.Println(RecursiveExcludeStringsBetween(l5, "BEGIN", "END"))

	// Output:
	// [abc ghi]
	// [ghi]
	// []
	// []
	// []
}