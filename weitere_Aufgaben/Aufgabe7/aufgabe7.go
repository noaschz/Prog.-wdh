package aufgabe7


//SumOdd: Erstellen Sie eine Funktion SumOdd,
//die die Summe aller ungeraden Zahlen in der Liste berechnet
// KA NR3 aber nicht rekursiv
func SumOdd(numbers []int) int {

	if len(numbers) == 0 {
		return 0
	}

head := numbers[0]


if head%2 != 0{
	return head + SumOdd(numbers[1:])
}

return SumOdd(numbers[1:])

}

/*
result := 0
	
	for _, i := range numbers {

		if i%2 != 0 {
			result = result + i
		}
	}

return result
*/