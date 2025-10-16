package own



func CountX(list []int, x int) int { // rekursiv


	if len(list) == 0{
		return 0
	}

	head := list[0]

	if head == x {
		return 1+ CountX(list[1:], x)
	}


	return CountX(list[1:], x)
}

/*if len(list) == 0 {
	return 0
}
if list[0] == x {
	return CountX(list[1:], x) + 1
}
return CountX(list[1:], x)
*/