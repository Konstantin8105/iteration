package iteration

import (
	"fmt"
	"os"
	"testing"
)

func Example() {
	var x float64
	var counter int
	if err := Find(&x, func() error {
		counter++
		fmt.Fprintf(os.Stdout, "Value for iteration %2d is %10.8e\n", counter, x)
		x = 5
		return nil
	}); err != nil {
		fmt.Fprintf(os.Stdout, "Error: %v", err)
		return
	}

	// Output:
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
}

func TestFunc(t *testing.T) {
	var x float64
	if err := Find(&x, func() error {
		x += 1.0
		return fmt.Errorf("Internal error")
	}); err == nil {
		t.Errorf("cannot found internal error")
		return
	}
}

func TestMaxIter(t *testing.T) {
	var x float64
	if err := Find(&x, func() error {
		x += 1.0
		return nil
	}); err == nil {
		t.Errorf("cannot found max iter error")
		return
	}
}
