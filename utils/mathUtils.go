package utils

// This function is **exported** because it starts with a capital letter
func Add(a, b int) int {
	return a + b
}

// This is **not exported** outside the package
func Subtract(a, b int) int {
	return a - b
}
