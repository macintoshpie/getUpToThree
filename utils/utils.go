package utils

func IsEqualUpToThree(result, num int) bool {
	if isLargerThanThree(num) {
		// undefined behavior
		return true
	}
	// return result == getUpToThree(num)
	return result == num
}

func isLargerThanThree(num int) bool {
	return num-3 > 0
}
