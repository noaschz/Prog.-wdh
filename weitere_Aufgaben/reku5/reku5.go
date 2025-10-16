package aufgabe3

// RecursiveCountOdd erwartet eine Liste von Zahlen und liefert die Anzahl der ungeraden Zahlen darin.
// Die Funktion muss rekursiv sein.
func RecursiveCountOdd(list []int) int {

if len(list)== 0{
	return 0
}

head := list[0]
tail := list[1:]

count := 0

if head%2 != 0{
	count++
}

return count + RecursiveCountOdd(tail)

}

/*
	if len(list) == 0 {
		return 0
	}

	count := 0
	if list[0]%2 != 0 {
		count = 1
	}

	return count + RecursiveCountOdd(list[1:])
*/