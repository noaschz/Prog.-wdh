package praxisaufgaben

// AUFGABENSTELLUNG:
// Liefert eine Liste mit allen Stellen in list,
// an denen x vorkommt.
func FindAll(list []int, x int) []int {
	result := []int{}

	for i := 0; i < len(list); i++ {
		if list[i] == x {
			result = append(result, i)
		}
	}
	return result
	
}



/*
result := []int{}

	for i := 0; i < len(list); i++ {
		if list[i] == x {
			result = append(result, i)
		}
	}
	return result
*/