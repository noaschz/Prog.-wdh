package praxisaufgaben

// AUFGABENSTELLUNG:
// Reparieren Sie die folgende Funktion.
// Es müssen alle Tests grün werden.
// Zusatzanforderung: Die Funktion muss rekursiv bleiben.

// Berechnet den ganzzahligen Quotienten x / y.
// D.h. die Funktion ignoriert Nachkommastellen bzw. den Rest.

import "fmt"

func ExampleDiv() {

	fmt.Println(Div(3, 2))
	fmt.Println(Div(2, 3))
	fmt.Println(Div(20, 2))
	fmt.Println(Div(-3, 2))
	fmt.Println(Div(3, -2))
	fmt.Println(Div(-3, -2))

	// Output:
	// 1
	// 0
	// 10
	// -1
	// -1
	// 1
}
