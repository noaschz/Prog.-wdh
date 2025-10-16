package aufgabe4

// RecursiveMaxElements erwartet zwei int-Listen und liefert eine Liste, die an jeder Position
// jeweils das größere der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt, gilt dieses Element als das größere.
// Die Funktion muss rekursiv sein.
func RecursiveMaxElements(l1, l2 []int) []int {
	
		if len(l1) == 0{
			return l2
		}

		if len(l2) == 0 {
			return l1
		}

		greater := l1[0]

		if greater < l2[0]{
			greater = l2[0]
		}

	return append([]int{greater}, RecursiveMaxElements(l1[1:], l2[1:])...)
}





/*
	if len(l1) == 0 {
		return l2
	}

	if len(l2) == 0 {
		return l1
	}

	greater := l1[0]
	if greater < l2[0] {
		greater = l2[0]
	}

	return append([]int{greater}, RecursiveMaxElements(l1[1:], l2[1:])...)
*/