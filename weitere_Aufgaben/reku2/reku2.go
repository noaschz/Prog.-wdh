package reku2

// RecursiveExcludeStringsBetween erwartet eine Liste von Strings und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
// Die Funktion muss rekursiv sein.
func RecursiveExcludeStringsBetween(list []string, first, last string) []string {

	return []string{}
}


/*
if len(list) == 0 {
		return []string{}
	}

	if list[0] == first {
		for i := 1; i < len(list); i++ {
			if list[i] == last {
				return append(list[:1], RecursiveExcludeStringsBetween(list[i+1:], first, last)...)
			}
		}
		return []string{}
	}

	return append([]string{list[0]}, RecursiveExcludeStringsBetween(list[1:], first, last)...)
*/