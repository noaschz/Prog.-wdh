package aufgabe7

// SumOdd: Erstellen Sie eine Funktion SumOdd,
// die die Summe aller ungeraden Zahlen in der Liste berechnet.
// Diese Version ist rekursiv.
func SumOdd(numbers []int) int {
    
    if len(numbers) == 0 {
        return 0}

head := numbers[0]
tail := numbers[1:]

if head%2 != 0 {
    return head + SumOdd(tail)
}
  return SumOdd(tail)

}




/*
if len(numbers) == 0 {
        return 0
    }

    head := numbers[0]
    tail := numbers[1:]

    if head%2 != 0 {
        return head + SumOdd(tail)
    } else {
        return SumOdd(tail)
    }
*/