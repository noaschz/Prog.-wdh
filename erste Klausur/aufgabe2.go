package own

func FilterBetween(list []int, m int, n int) []int {

	result := []int{}

	for _, val := range list{
		if val > m && val < n {
			result = append(result, val)
		}
	}

	return result
}










/*
result := []int{}

	for i := 0; i < len(list); i++ {
		if list[i] > m && list[i] < n {
			result = append(result, list[i])
		}
	}
	return result
*/