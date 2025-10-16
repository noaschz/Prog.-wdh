package main

import "fmt"

func ExampleShortestXYZ() {
	words := []string{"abcdef", "xyzhello", "xyz", "xyzworld", "xyzshortest"}
	fmt.Println(ShortestXYZ(words))
	fmt.Println(ShortestXYZ([]string{"hello", "world"}))
	fmt.Println(ShortestXYZ([]string{"xyzabc", "xyz", "xyzdefg"}))

	// Output:
	// xyz
	//
	// xyz
}
