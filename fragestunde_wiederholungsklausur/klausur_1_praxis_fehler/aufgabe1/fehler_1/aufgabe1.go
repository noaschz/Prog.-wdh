package aufgabe1

// LongestAbc erwartet eine Liste von Strings und liefert
// das längste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
func LongestAbc(list []string) string {

	longestlen := -1
	longestpos := -1

	for pos, val := range list {
		curretlen := len(val)

		if curretlen >= 3 && val[:3] == "abc" &&
		curretlen > longestlen {
			longestlen = curretlen
			longestpos = pos
		}
		}

		if longestpos != -1 {
			return list[longestpos]
		}

		return ""
}







/*longestLen := 100 // 100 ist eine "große" Zahl, nicht geeignet als Startwert.
	longestPos := 100

	for pos, val := range list {
		currentLen := len(val)
		if currentLen <= 3 && val[1:] == "abcde" { // ">=", "val[:3]", "abcde"
			if currentLen < longestLen {
				longestLen = currentLen
				longestPos = pos
			}
		}
	}
	if longestPos != 100 {
		return list[longestPos]
	}
		*/