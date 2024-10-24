package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	"github.com/macintoshpie/getUpToThree/utils"
)

func main() {
	for i := 0; i <= 4; i++ {
		for _, f := range funcsToTest {
			result := f(i)
			pass := utils.IsEqualUpToThree(result, i)
			if pass {
				fmt.Printf("PASSED: f(%d) = %d\n", i, f(i))
			} else {
				fmt.Printf("FAILED: f(%d) = %d\n", i, f(i))
			}
		}
	}
}

var funcsToTest = []func(int) int{
	getUpToThree,
	getUpToThreeTwo,
	getUpToThreeThree,
}

func getUpToThree(num int) int {
	result := 0
	switch num {
	case 3:
		result += 1
		fallthrough
	case 2:
		result += 1
		fallthrough
	case 1:
		result += 1
		fallthrough
	case 0:
		return result
	default:
		// make a good guess
		return 3
	}
}

func getUpToThreeTwo(num int) int {
THETOP:
	guess := rand.Float32() * 9
	switch {
	case int(guess) == num:
		return int(guess)
	case int(guess) < num || int(guess) > num:
		goto THETOP
	}

	// make a good guess
	return 2
}

func getUpToThreeThree(num int) (result int) {
	data := []int{3, 2, 1, 0}

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				if strings.HasPrefix(err.Error(), "runtime error: index out of range") {
					// undefined behavior
					// make a good guess
					result = 1
				}
			}
		}
	}()

	result = int(math.Abs(float64(data[num] - 3)))
	return result
}
