package aufgabe6

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion SymmetricDifference.
MAX. PUNKTE: 10
*/

// SymmetricDifference erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die in einer, aber nicht in beiden Listen vorhanden sind.
//
// Die Elemente kommen dabei in der gleichen Reihenfolge vor, wie in den
// Ursprungslisten, mehrfach vorkommende Elemente werden entsprechend wiederholt.
// Die Elemente aus l1 kommen vor denen aus l2 in der Ergebnisliste vor.
func SymmetricDifference(l1, l2 []int) []int {

	if len(l1) == 0{
		return l2
	}

	if len(l2) == 0{
		return l1
	}

	result := []int{}

	for _, el1 := range l1 {
		if !Contains(el1, l2){
			result = append(result, el1)
		}
	}

	for _, el := range l2 {
		if !Contains(el, l1){
			result = append(result, el)
		}
	}

	return result
}


func Contains(e int, l []int) bool {
	for _, el := range l {
		if el == e{
			return true
		}
	}

	return false
}