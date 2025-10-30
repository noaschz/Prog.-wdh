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
	shorteststring := 100
	shortestposition := 100

	for i := 0; i < len(list); i++ {
		currentstring := len(list[i])
		currentposition := i
		if currentstring >= 3 && list[currentposition][:3] == "abc" {
			if currentstring < shorteststring {
				shorteststring = currentstring
				shortestposition = currentposition
			}
		}
	}

	if shorteststring != 100 {
		return list[shortestposition]
	}
	
	return ""
}
