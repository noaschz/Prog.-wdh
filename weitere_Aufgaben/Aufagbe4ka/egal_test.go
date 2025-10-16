package aufagbe8

import "fmt"

func ExampleProductOdd() {
    l1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
    l2 := []int{1, 3, 5, 7, 9}
    l3 := []int{2, 4, 6, 8, 10}
    l4 := []int{5}

    fmt.Println(ProductOdd(l1))
    fmt.Println(ProductOdd(l2))
    fmt.Println(ProductOdd(l3))
    fmt.Println(ProductOdd(l4))

    // Output:
    // 105
    // 945
    // 1
    // 5
}