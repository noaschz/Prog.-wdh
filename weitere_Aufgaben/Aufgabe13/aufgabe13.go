package aufgabe13

//Union: Erstellen Sie eine Funktion Union, die die Vereinigung der beiden Listen zurückgibt.

func Union(a, b []int) []int {
	
	if len(a) == 0 {
		return b
	}

	if len(a) == 0 {
		return b
	}
	
	result := []int{}

	for _, el := range a {
		if !Contains(el, result) {
			result = append(result, el)
		}
	}

	for _, el := range b {
		if !Contains(el, result) {
			result = append(result, el)
		}
	}


	return result
}

func Contains(e int , l []int) bool {
	for _, el := range l {
		if el == e {
			return true
		}
	}
	return false
}


/*
result := []int{}


		if len(a) == 0 {
			return b
		}

		if len(a) == 0 {
			return b
		}

		for _, el1 := range a {
			if !Contains(el1, result) {
				result = append(result, el1)
			}
		}
	
		for _, el2 := range b {
			if !Contains(el2, result) {
				result = append(result, el2)
			}
		}

	return result

}


func Contains(e int , l []int) bool {
	for _, el := range l {
		if el == e {
			return true
		}
	}

	return false
*/