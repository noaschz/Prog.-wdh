package own

func ProductOfLists(l1, l2 []int) []int { // rekursiv

	if len(l1) == 0 {
		return l2
	}

	if len(l2) == 0{
		return l1
	}


	return append([]int{l1[0] * l2[0]}, ProductOfLists(l1[1:], l2[1:])...)

}








/*if len(l1) == 0 {
	return l2
}
if len(l2) == 0 {
	return l1
}
if len(l1) == 0 && len(l2) == 0 {
	return []int{}
}
return append([]int{l1[0] * l2[0]}, ProductOfLists(l1[1:], l2[1:])...)
}
*/