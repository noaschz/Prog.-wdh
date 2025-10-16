package aufgabe9

//MinElements: Erstellen Sie eine Funktion MinElements, 
//die an jeder Position das kleinere der beiden Elemente enthält.

func MinElements(a, b []int) []int {

    if len(a) == 0 {
        return b
    }

    if len(b) == 0 {
        return a
    }


    less := a[0]

    if a[0] > b[0]{
        less = b[0]
    }

    return append([]int{less}, MinElements(a[1:], b[1:])...)
}





/*
if len(a) == 0{
    return b
}
   
if len(b) == 0{
    return a
}


    less := a[0]

    if less >= b[0]{
        less = b[0]
    }

    return append([]int{less}, MinElements(a[1:], b[1:])...)
*/