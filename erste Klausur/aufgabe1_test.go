package own

import "fmt"

func ExampleLongestAbc_standard_cases() {
	list1 := []string{"xyz", "abcde", "abcdef", "ab", "abc"}
	list2 := []string{"abcxyz", "abcdef", "a", "abc"}

	fmt.Println(LongestAbc(list1))
	fmt.Println(LongestAbc(list2))

	// Output:
	// abcdef
	// abcdef

}

func ExampleLongestAbc_special_cases() {
	list1 := []string{"xyz", "def", "hello"}
	list2 := []string{"abc", "abc", "abc"}

	fmt.Println(LongestAbc(list1)) // Kein "abc" vorhanden, sollte ""
	fmt.Println(LongestAbc(list2)) // Alle gleich lang, soll "abc" zurückgeben

	// Output:
	//
	// abc
}
