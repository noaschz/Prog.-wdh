package main

func FindUniqueElements(list []int) []int {

	result := []int{}

	for i := 0; i < len(list); i++ {
		counter := 0
		for j := 0; j < len(list); j++ {
			if list[i] == list[j] {
				counter++
			}
		}
		if counter == 1 {
			result = append(result, list[i])
		}
	}
	return result
}
