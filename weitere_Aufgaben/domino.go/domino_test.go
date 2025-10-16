package aufgabe10

import "fmt"

func ExampleIsValidDominoChain() {
	d1 := []Domino{{1, 2}, {2, 3}, {3, 4}}
	d2 := []Domino{{1, 2}, {3, 4}, {4, 5}}
	d3 := []Domino{{1, 1}, {1, 1}, {1, 1}}
	d4 := []Domino{{1, 2}}

	fmt.Println(IsValidDominoChain(d1))
	fmt.Println(IsValidDominoChain(d2))
	fmt.Println(IsValidDominoChain(d3))
	fmt.Println(IsValidDominoChain(d4))

	// Output:
	// true
	// false
	// true
	// true
}
