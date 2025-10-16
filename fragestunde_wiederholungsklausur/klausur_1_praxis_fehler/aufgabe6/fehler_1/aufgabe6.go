package aufgabe6

// Duplicates erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die mehr als einmal in l1 vorkommen.
// Im Ergebnis kommt jedes Element nur einmal vor.
// Die Elemente stehen im Ergebnis in der Reihenfolge ihres ersten Auftretens.
func Duplicates(list []int) []int {

	result := []int{}
	for i, el1 := range list {
		if !Contains(el1, result) && 
		Contains(el1, list[i+1:]) {
			result = append(result, el1)
		}
	}
	return result
}

	func Contains(e int, l []int) bool {
		for _, el := range l{
			if el == e{
				return true
			}
		}
		return false
	}

/*
result := []int{}
	for _, el1 := range list {
		if Contains(list, el1) {
			result = append(result, el1)
		}
	}
	return result
}

// überpüfen ob el nochmal in liste ist fehlt mir
// Hilfsfunktion Contains
func Contains(l []int, el int) bool {
	for _, e := range l {
		if e == el {

		}
	}
	return false
}
*/
