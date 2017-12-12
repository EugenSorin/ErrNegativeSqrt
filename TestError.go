package main

import (
	"fmt"
	"math"
)

const eps = 0.00000001

type ErrNegativeSqrt float64 // Reține chiar valoarea numărului blucucaș

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	var z1, z2 float64 // sunt 0
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z2 = x / 2
	for math.Abs(z1-z2) > eps {
		z1 = z2
		z2 -= (z2*z2 - x) / (2 * z2)
	}
	return z2, nil
}

func main() {
  fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
