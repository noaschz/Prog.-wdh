package functional

import "fmt"


func ExamplePrimeFactors(){
	fmt.Println(PrimeFactors(60, 6))
	fmt.Println(PrimeFactors(42, 5))
	fmt.Println(PrimeFactors(23, 2))

	//Output:
	//[2 2 3 5]
	//[2 3 7]
	//[23]
}