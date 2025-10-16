package aufagbe8

// ProductOdd: Erstellen Sie eine Funktion ProductOdd,
// die das Produkt aller ungeraden Zahlen in der Liste berechnet.
// Die Funktion muss rekursiv sein.
func ProductOdd(numbers []int) int {

    if len(numbers) == 0{
        return 1
    }

    head := numbers[0]

    if head%2 != 0 {
        return numbers[0] * ProductOdd(numbers[1:])
    }
  return ProductOdd(numbers[1:])
}




/*
if len(numbers) == 0 {
        return 1
    }
    if numbers[0]%2 != 0 {
        return numbers[0] * ProductOdd(numbers[1:])
    }
    return ProductOdd(numbers[1:])
}
*/