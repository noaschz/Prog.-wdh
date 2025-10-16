package functional

import "fmt"

func IsPrime(n int) bool {

	return n != 1 && !DividByAny(n, n-1)
}

func DividByAny(n,c int) bool{
	if c <= 1 {
		return false
	}
	
	if n%c == 0 {
		return true
	}
	return DividByAny(n, c -1)
}

func ExampleIsPrime() {
	fmt.Println(IsPrime(1))
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(3))
	fmt.Println(IsPrime(4))
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(6))

	//Output:
	//false
	//true
	//true
	//false
	//true
	//false
}
