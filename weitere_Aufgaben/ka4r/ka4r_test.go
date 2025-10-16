package aufgabe9

import "fmt"

func ExampleSumEven() {
    l1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
    l2 := []int{2, 4, 6, 8, 10}
    l3 := []int{1, 3, 5, 7, 9}
    l4 := []int{}

    fmt.Println(SumEven(l1))
    fmt.Println(SumEven(l2))
    fmt.Println(SumEven(l3))
    fmt.Println(SumEven(l4))

    // Output:
    // 20
    // 30
    // 0
    // 0
}