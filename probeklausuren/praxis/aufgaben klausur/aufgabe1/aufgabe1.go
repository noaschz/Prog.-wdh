package aufgabe1


/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ShortestAbc.
MAX. PUNKTE: 10
*/

// ShortestAbc erwartet eine Liste von Strings und liefert
// das kürzeste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
//
// Hinweis: Die Funktion muss nur mit kurzen Strings der Länge < 100 funktionieren.
func ShortestAbc(list []string) string {

	shortestlen := 100
	shortestpos := -1

	for pos, val := range list {
			currentlen := len(val)
		if currentlen >= 3 && val[:3] == "abc" &&
		currentlen < shortestlen {
			shortestlen = currentlen
			shortestpos = pos
		}
	}

	if shortestpos != -1 {
		return list[shortestpos]
	}
	return ""
	}





