package aufgabe3

// CountX erwartet eine Liste von Zahlen sowie eine Zahl x.
// CountX liefert die Anzahl der Vorkommen von x in der Liste.
func CountX(list []int, x int) int {

if len(list) == 0{
	return 0
}

head := list[0]

if head == x {
	return 1+ CountX(list[1:], x)
}
	return CountX(list[1:], x)
}



/*
	if len(list) == 0 {
		return 0
	}
	head, result := list[0], CountX (list[1:])
	if head = x {
		result ++
	}
return result
*/
