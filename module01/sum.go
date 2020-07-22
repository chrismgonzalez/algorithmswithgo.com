package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	result := 0
	for _, v := range numbers {
		result += v
	}
	return result
}

// todo - implement a recursive solution
