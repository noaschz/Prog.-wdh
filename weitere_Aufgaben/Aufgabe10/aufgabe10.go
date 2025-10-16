package aufgabe10

//SumElements: Erstellen Sie eine Funktion SumElements,
//die an jeder Position die Summe der beiden Elemente enthält..

func SumElements(a, b []int) []int {


    if len(a) == 0 {
        return b
    }

    if len(b) == 0 {
        return a
    }

    return append([]int{a[0]+ b[0]}, SumElements(a[1:], b[1:])...)
}


/*
maxLength := len(a)
    if len(b) > maxLength {
        maxLength = len(b)
    }

    result := make([]int, maxLength)

    for i := 0; i < maxLength; i++ {
        if i < len(a) {
            result[i] += a[i]
        }
        if i < len(b) {
            result[i] += b[i]
        }
    }


    return result
*/
