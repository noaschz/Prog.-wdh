package praxisaufgaben

// AUFGABENSTELLUNG:
// Erwartet drei Zahlen m,n,max > 0 und liefert eine Liste mit allen
// gemeinsamen Vielfachen von m und n, die nicht größer als max sind.
func CommonMultiples(m, n, max int) []int {

	var Vielfachen []int

	for i := 1; i <= max; i++ {
		if i%m == 0 && i%n == 0 {
			Vielfachen = append(Vielfachen, i)
		}
	}
	return Vielfachen
}

/*
result := []int{}

	for i := 1; i <= max; i++ {
		if i%n == 0 && i%m == 0 {
			result = append(result, i)
		}
	}
	return result
*/
