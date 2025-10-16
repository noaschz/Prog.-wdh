package aufgabe12

import "fmt"

func ExampleCard_EqualTo() {
	c1 := Card{2, 2}
	c2 := Card{2, 3}
	c3 := Card{3, 3}
	c4 := Card{2, 2}

	fmt.Println(c1.EqualTo(c2))
	fmt.Println(c2.EqualTo(c1))
	fmt.Println(c1.EqualTo(c3))
	fmt.Println(c3.EqualTo(c1))
	fmt.Println(c1.EqualTo(c4))

	// Output:
	// false
	// false
	// false
	// false
	// true
}