package reku1

// RecursiveShortestAbc erwartet eine Liste von Strings und liefert
// das kürzeste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
// Die Funktion muss rekursiv sein.
func RecursiveShortestAbc(list []string) string {
    if len(list) == 0 {
        return ""
    }

    if len(list[0]) >= 3 && list[0][:3] == "abc" {
        rest := RecursiveShortestAbc(list[1:])
        if rest == "" || len(list[0]) < len(rest) {
            return list[0]
        }
        return rest
    }

    return RecursiveShortestAbc(list[1:])
}



/*
if len(list) == 0 {
		return ""
	}

	head := list[0]
	tail := RecursiveShortestAbc(list[1:])

	if len(head) >= 3 && head[:3] == "abc" {
		if tail == "" || len(head) < len(tail) {
			return head
		}
	}
*/