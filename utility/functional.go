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

func incrementBy(incrementBy int) (startsWith func(int) func() int) {
	return func(start int) func() int {
		curr := start
		return func() int {
			curr += incrementBy
			return start
		}
	}
}
