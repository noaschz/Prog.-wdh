package main

func ReverseNumber(x int) int { // rekursiv
	return Helper(x, 0)
}

/*
func Helper(x, prev int) int {
	if x == 0 {
		return prev
	}
	return Helper(x/10, prev*10+x%10)
*/
