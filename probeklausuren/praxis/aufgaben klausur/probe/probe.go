package probe

import "fmt"

func CountDown(n int) {
	if n <= 0 { 	// sagt wann man abbricht
		return
	}

	fmt.Println(n)
	CountDown(n-1)

}





