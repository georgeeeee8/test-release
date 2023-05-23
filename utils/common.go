package utils

// Add 相加
func Add(a int, b int) int {
	return a + b
}

// AddMultiple 多數值相加
func AddMultiple(values ...int) int {
	var result int
	for _, v := range values {
		result += v
	}

	return result
}
