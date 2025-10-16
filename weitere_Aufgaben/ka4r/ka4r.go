package aufgabe9

// SumEven: Erstellen Sie eine Funktion SumEven,
// die die Summe aller geraden Zahlen in der Liste berechnet.
// Diese Version ist rekursiv.
func SumEven(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }

    oddnumbers := numbers[0]
    tail := numbers[1:]
    
    if oddnumbers%2 == 0 {
        return oddnumbers + SumEven(tail)
    } else {
        return SumEven(tail)
    }
}


/*

if len(numbers) == 0 {
        return 0
    }

    head := numbers[0]
    tail := numbers[1:]

    if head%2 == 0 {
        return head + SumEven(tail)
    } else {
        return SumEven(tail)
    }
*/