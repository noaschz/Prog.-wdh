package main

func MergeSortedListsRecursive(l1, l2 []int) []int {

	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}
	if len(l1) == 0 && len(l2) == 0 {
		return []int{}
	}
	if l1[0] > l2[0] {
		return append([]int{l2[0]}, MergeSortedListsRecursive(l1, l2[1:])...)
	}
	return append([]int{l1[0]}, MergeSortedListsRecursive(l1[1:], l2)...)
}
