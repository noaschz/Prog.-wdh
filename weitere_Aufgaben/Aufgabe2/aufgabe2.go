package aufgabe2

//CountAbc: Erstellen Sie eine Funktion CountAbc, 
//die die Anzahl der Elemente zählt, die mit "abc" beginnen

func CountAbc(list []string) int {

	result := 0

	for _, val := range list {
		if len(val) >=3  && val[:3] == "abc" {
			result++
		}
	}

	return result
}






/*
result := 0

	for _, val := range list {

		if len(val) >= 3 && val[:3] == "abc" {
			result++
		}
	}

return result
*/