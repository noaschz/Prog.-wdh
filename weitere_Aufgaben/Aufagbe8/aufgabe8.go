package aufagbe8

//ProductOdd: Erstellen Sie eine Funktion ProductOdd, 
//die das Produkt aller ungeraden Zahlen in der Liste berechnet.
// ka nr4 aber mussrekursiv sein 
func ProductOdd(numbers []int) int {
	
	result := 1
	
	for _, i := range numbers {

		if i%2 != 0 {
			result *=   i
		}
	}

	return result
}


/*
result := 1
	
	for _, i := range numbers {

		if i%2 != 0 {
			result *=   i
		}
	}
*/