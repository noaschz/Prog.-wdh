package main
import "fmt"

func ExampleLongestPrefixMatch_standard_cases() {
	list := []string{"pretest", "prefixmatch", "pre", "prefix"}
	fmt.Println(LongestPrefixMatch(list, "pre"))
	fmt.Println(LongestPrefixMatch(list, "prefix"))

	// Output:
	// prefixmatch
	// prefixmatch
}
