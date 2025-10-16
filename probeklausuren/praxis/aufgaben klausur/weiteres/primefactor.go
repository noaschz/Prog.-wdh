package functional

// func PrimeFactors(n int) []int {

// 	result := []int{}
// 	c := 2
// 	for n >1 {
// 		if n % c == 0{
// 			result = append(result,  c)
// 			n = n/c
// 		} else {
// 			c++
// 		}
// 	}

// 	return result
// }

//jetzt rekursiv

func PrimeFactors(n, c int) []int {

	if n <= 1 {
		return []int{}
	}

	result := []int{}

	for n > 1 {
		if n%c == 0 {
			result = append(result, c)
			n = n / c

		}

		return append(result, PrimeFactors(n, c)...)
	}
	return []int{}
}
