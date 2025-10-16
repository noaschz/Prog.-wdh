package aufgabe6

import "fmt"

func ExampleExcludeStringsOutside() {
	l1 := []string{"abc", "BEGIN", "def", "END", "ghi"}
	l2 := []string{"BEGIN", "def", "END", "ghi"}
	l3 := []string{"BEGIN", "def", "ghi", "END"}
	l4 := []string{"END", "def", "BEGIN", "ghi"}
	l5 := []string{"abc", "def", "ghi"}

	fmt.Println(ExcludeStringsOutside(l1, "BEGIN", "END"))
	fmt.Println(ExcludeStringsOutside(l2, "BEGIN", "END"))
	fmt.Println(ExcludeStringsOutside(l3, "BEGIN", "END"))
	fmt.Println(ExcludeStringsOutside(l4, "BEGIN", "END"))
	fmt.Println(ExcludeStringsOutside(l5, "BEGIN", "END"))

	// Output:
	// [abc ghi]
	// [ghi]
	// []
	// []
	// []
}