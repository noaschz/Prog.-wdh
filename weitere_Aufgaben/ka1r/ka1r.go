package aufgabe8


// ProductOdd: Erstellen Sie eine Funktion ProductOdd,
// die das Produkt aller ungeraden Zahlen in der Liste berechnet.
// Diese Version ist rekursiv.
func ProductOdd(numbers []int) int {
    
if len(numbers) == 0{
    return 1
}


    head := numbers[0]
    tail := numbers[1:]

    if head%2 != 0 {
        return head * ProductOdd(tail)
    }

    return ProductOdd(tail)

}





/*
if len(numbers) == 0 {
        return 1
    }

    head := numbers[0]
    tail := numbers[1:]

    if head%2 != 0 {
        return head * ProductOdd(tail)
    } else {
        return ProductOdd(tail)
    }
*/