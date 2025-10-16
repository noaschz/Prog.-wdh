package main

func CommonElements(l1 []int, l2 []int) []int {

	return []int{}
}


/*

	result := []int{}

	if len(l1) == 0 && len(l2) == 0 {
		return []int{}
	}
	for i := 0; i < len(l1); i++ {
		counter := 0
		for j := 0; j < len(l2); j++ {
			if l1[i] == l2[j] {
				counter++
			}
		}
		if counter == 1 {
			result = append(result, l1[i])
		}
	}
	return result
*/