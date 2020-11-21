package utility

// Add two numbers
func Add(x, y int) (res int) {
	res = x + y
	return
}

// Minus two numbers
func Minus(x, y int) (res int) {
	res = x - y
	return res
}

// IncrementBy a number and start with x
func IncrementBy(incrementBy int) (startsWith func(int) func() int) {
	return func(start int) func() int {
		curr := start
		return func() int {
			curr += incrementBy
			return curr
		}
	}
}
