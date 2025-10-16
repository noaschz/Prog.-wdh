package aufgabe3

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion CountOdd.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein.
*/

// CountOdd erwartet eine Liste von Zahlen und liefert die Anzahl der ungeraden Zahlen darin.
func CountOdd(list []int) int {

	if len(list) == 0{
		return 0
	}

	head := list[0]

	if head%2 != 0{
		return 1+ CountOdd(list[1:])
	}

	return CountOdd(list[1:])
}
