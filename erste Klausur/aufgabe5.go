package own


func Duplicates(list []int) []int { // nehme alle Dinge einmal in die Liste auf, die zwei oder mehrmals vorkommen!

	if len(list) == 0 {
		return []int{}
	}

	result := []int{}

	for i := 0; i < len(list); i++{
		counter := 0
		for j := 0; j < len(list); j++{
			if list[i] == list[j]{
				counter++
			}
			if counter > 1 && !Contains(list[i], result) {
				result = append(result, list[i])
			}
		}
		
	}
		return result
}


func Contains(e int, l []int) bool {

	for _, el := range l {
		if el == e {
			return true
		}
	}

	return false

}




/*result := []int{}
	for i := 0; i < len(list); i++ {
		counter := 0
		for j := 0; j < len(list); j++ {
			if list[i] == list[j] {
				counter++
			}
		}
		if counter > 1 && !Contains(result, list[i]) {
			result = append(result, list[i])
		}
	}
	return result
}

func Contains(list []int, x int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == x {
			return true
		}
	}
	return false*/