# iteration
iteration by single variable

```
package iteration // import "."

const (
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

func Run(x *float64, f func() error) error
    Run iteration by single variable
```
