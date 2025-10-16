package aufgabe11

import "fmt"

func ExampleCard_LessThan() {
	c1 := Card{2, 2}
	c2 := Card{2, 3}
	c3 := Card{3, 3}

	fmt.Println(c1.LessThan(c2))
	fmt.Println(c2.LessThan(c1))
	fmt.Println(c1.LessThan(c3))
	fmt.Println(c3.LessThan(c1))
	fmt.Println(c2.LessThan(c3))
	fmt.Println(c3.LessThan(c2))

	// Output:
	// true
	// false
	// false
	// false
	// false
	// false
}