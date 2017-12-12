package main

import (
	"fmt"
	"math"
)

const eps = 0.00000001

type ErrNegativeSqrt float64 // The source of error is a bad (negative) number

func (e ErrNegativeSqrt) Error() string {
    // The fmt.Sprintf() function use the e.Error() method to represent e as a string;
    // So a call to fmt.Sprint(e) in the next line will send the program into an infinite loop;
    // To avoid this, we must convert e first - in fact e is a number
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	var z1, z2 float64 = 0, x/2
	if x < 0 {
        // Note the conversion from float64 (x) to ErrNegativeSqrt.
		return 0, ErrNegativeSqrt(x)
	}
	for math.Abs(z1-z2) > eps {
		z1 = z2
		z2 -= (z2*z2 - x) / (2*z2)
	}
	return z2, nil
}

func main() {
    fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
