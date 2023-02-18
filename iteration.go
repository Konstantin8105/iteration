package iteration

import (
	"fmt"
	"math"
)

// Constants
var (
	// Precision of rott-finding
	Precision float64 = 1e-6

	// MaxIteration is max allowable amount of iteration.
	// Typically for precition=1e-6 need 20 iterations.
	//
	// Example:
	//
	// Value for iteration  1 is 0.00000000e+00
	// Value for iteration  2 is 3.09016994e+00
	// Value for iteration  3 is 4.27050983e+00
	// Value for iteration  4 is 4.72135955e+00
	// Value for iteration  5 is 4.89356882e+00
	// Value for iteration  6 is 4.95934691e+00
	// Value for iteration  7 is 4.98447190e+00
	// Value for iteration  8 is 4.99406879e+00
	// Value for iteration  9 is 4.99773448e+00
	// Value for iteration 10 is 4.99913465e+00
	// Value for iteration 11 is 4.99966947e+00
	// Value for iteration 12 is 4.99987375e+00
	// Value for iteration 13 is 4.99995178e+00
	// Value for iteration 14 is 4.99998158e+00
	// Value for iteration 15 is 4.99999296e+00
	// Value for iteration 16 is 4.99999731e+00
	// Value for iteration 17 is 4.99999897e+00
	// Value for iteration 18 is 4.99999961e+00
	//
	MaxIteration int = 500

	// Ratio for choose value on next iteration.
	// Ranges:
	//
	//		negative ratio - not acceptable
	//		0(zero)        - not acceptable
	//		0...1          - acceptable
	//		more 1.0       - not acceptable
	//
	// Recomendation - use "golden ratio" by
	// https://en.wikipedia.org/wiki/Golden_ratio
	//
	Ratio float64 = 2.0 / (1.0 + math.Sqrt(5.0))
)

type ErrorFind struct {
	Type ErrType
	Err  error
}

func (e ErrorFind) Error() string {
	return fmt.Sprintf("%s:%s", e.Type, e.Err)
}

type ErrType int8

const (
	MaximalIteration ErrType = iota
	InternalErr
	NotValidValue
)

func (et ErrType) String() string {
	switch et {
	case MaximalIteration:
		return "max iteration"
	case InternalErr:
		return "internal error"
	case NotValidValue:
		return "not valid value"
	}
	return "undefined"
}

// Run iteration by many variable
func Find(f func() error, xs ...*float64) (err error) {
	maxIter, precision, ratio := MaxIteration, Precision, Ratio
	xLast := make([]float64, len(xs))
	for i := range xs {
		xLast[i] = *xs[i]
	}
	exit := false
	for iter := 0; ; iter++ {
		if iter >= maxIter {
			return ErrorFind{Type: MaximalIteration}
		}
		if err = f(); err != nil {
			return ErrorFind{
				Type: InternalErr,
				Err:  err,
			}
		}
		exit = true
		for i := range xLast {
			if math.IsNaN(*xs[i]) {
				return ErrorFind{
					Type: NotValidValue,
					Err:  fmt.Errorf("Parameter %d is NaN", i),
				}
			}
			if math.IsInf(*xs[i], 0) {
				return ErrorFind{
					Type: NotValidValue,
					Err:  fmt.Errorf("Parameter %d is infinity", i),
				}
			}
			if xLast[i] == 0.0 {
				if precision < math.Abs(*xs[i]) {
					exit = false
				}
			} else {
				if precision < math.Abs((*xs[i]-xLast[i])/xLast[i]) {
					exit = false
				}
			}
		}
		if exit {
			break
		}
		// calculate value for next iteration
		for i := range xLast {
			*xs[i] = xLast[i] + (*xs[i]-xLast[i])*ratio
		}
		// store last iteration value
		for i := range xs {
			xLast[i] = *xs[i]
		}
	}
	return nil
}
