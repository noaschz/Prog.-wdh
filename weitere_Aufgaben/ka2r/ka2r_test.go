package aufgabe7

import "fmt"

func ExampleSumOdd() {
    l1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
    l2 := []int{1, 3, 5, 7, 9}
    l3 := []int{2, 4, 6, 8, 10}
    l4 := []int{}

    fmt.Println(SumOdd(l1))
    fmt.Println(SumOdd(l2))
    fmt.Println(SumOdd(l3))
    fmt.Println(SumOdd(l4))

    // Output:
    // 16
    // 25
    // 0
    // 0
}