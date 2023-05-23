package utils

import "fmt"

func ExampleAdd() {
	// 將數值做相加運算
	result := Add(2, 3)

	// Output: 5
	fmt.Println(result)
}

func ExampleAddMultiple() {
	// 將多數值做相加運算
	sum := AddMultiple([]int{1, 2, 3}...)

	// Output: 6
	fmt.Println(sum)
}
