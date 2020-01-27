package iteration

import (
	"fmt"
	"math"
)

// Constants
const (
	// Precision of rott-finding
	Precision float64 = 1e-6

	// MaxIteration is max allowable amount of iteration.
	// Typically for precition=1e-6 need 25 iterations.
	//
	// Example:
	//
	//		Value for iteration  1 is 0.00000000e+00
	//		Value for iteration  2 is 2.50000000e+00
	//		Value for iteration  3 is 3.75000000e+00
	//		Value for iteration  4 is 4.37500000e+00
	//		Value for iteration  5 is 4.68750000e+00
	//		Value for iteration  6 is 4.84375000e+00
	//		Value for iteration  7 is 4.92187500e+00
	//		Value for iteration  8 is 4.96093750e+00
	//		Value for iteration  9 is 4.98046875e+00
	//		Value for iteration 10 is 4.99023438e+00
	//		Value for iteration 11 is 4.99511719e+00
	//		Value for iteration 12 is 4.99755859e+00
	//		Value for iteration 13 is 4.99877930e+00
	//		Value for iteration 14 is 4.99938965e+00
	//		Value for iteration 15 is 4.99969482e+00
	//		Value for iteration 16 is 4.99984741e+00
	//		Value for iteration 17 is 4.99992371e+00
	//		Value for iteration 18 is 4.99996185e+00
	//		Value for iteration 19 is 4.99998093e+00
	//		Value for iteration 20 is 4.99999046e+00
	//		Value for iteration 21 is 4.99999523e+00
	//		Value for iteration 22 is 4.99999762e+00
	//		Value for iteration 23 is 4.99999881e+00
	//		Value for iteration 24 is 4.99999940e+00
	//
	MaxIteration int = 500
)

// Run iteration by single variable
func Run(x *float64, f func() error) error {
	maxIter, precision := MaxIteration, Precision
	for iter, xLast := 0, *x; ; iter++ {
		if iter >= maxIter {
			return fmt.Errorf("max iter error")
		}
		if err := f(); err != nil {
			return fmt.Errorf("%v", err)
		}
		if math.Abs(*x-xLast) < precision {
			break
		}
		// calculate value for next iteration
		*x = xLast + (*x-xLast)/2.0
		// store last iteration value
		xLast = *x
	}
	return nil
}
