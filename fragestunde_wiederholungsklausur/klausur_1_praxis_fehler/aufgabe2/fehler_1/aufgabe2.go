package aufgabe2

// ExcludeBetween erwartet eine Liste und zwei Zahlen m und n.
// Die Funktion liefert eine Liste mit allen Elementen x, für die gilt: m < x < n.
func ExcludeBetween(list []int, m, n int) []int {

	result := []int{}

	for i := range list {
			if list[i] >= m && list[i] <= n {
				result = append(result, list[i])
			}
	}

return result

}

/*
firstpos := -1
	lastpos := -1

	for pos, s := range list {
		if s == first {
			firstpos = pos
		}
		if s == last {
			lastpos = pos
		}
		if lastpos <= firstpos {
			return []string{}
		}
	}

	return append(list[:firstpos], list[lastpos+1:]...)
*/