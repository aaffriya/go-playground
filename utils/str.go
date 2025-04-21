package utils

// This function is **exported** because it starts with a capital letter
func AddSTR(a, b int) int {
	return a + b
}

// This is **not exported** outside the package
func SubtractSTR(a, b int) int {
	return a - b
}
