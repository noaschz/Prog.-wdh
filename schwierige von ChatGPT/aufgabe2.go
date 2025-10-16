package main

func SumOfEvensRecursive(list []int) int {

	if len(list) == 0 {
		return 0
	}
	if list[0]%2 == 0 {
		return SumOfEvensRecursive(list[1:]) + list[0]
	}
	return SumOfEvensRecursive(list[1:])
}
