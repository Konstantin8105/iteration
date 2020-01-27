package iteration

import (
	"fmt"
	"os"
	"testing"
)

func Example() {
	var x float64
	var counter int
	if err := Run(&x, func() error {
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
	// Value for iteration  2 is 2.50000000e+00
	// Value for iteration  3 is 3.75000000e+00
	// Value for iteration  4 is 4.37500000e+00
	// Value for iteration  5 is 4.68750000e+00
	// Value for iteration  6 is 4.84375000e+00
	// Value for iteration  7 is 4.92187500e+00
	// Value for iteration  8 is 4.96093750e+00
	// Value for iteration  9 is 4.98046875e+00
	// Value for iteration 10 is 4.99023438e+00
	// Value for iteration 11 is 4.99511719e+00
	// Value for iteration 12 is 4.99755859e+00
	// Value for iteration 13 is 4.99877930e+00
	// Value for iteration 14 is 4.99938965e+00
	// Value for iteration 15 is 4.99969482e+00
	// Value for iteration 16 is 4.99984741e+00
	// Value for iteration 17 is 4.99992371e+00
	// Value for iteration 18 is 4.99996185e+00
	// Value for iteration 19 is 4.99998093e+00
	// Value for iteration 20 is 4.99999046e+00
	// Value for iteration 21 is 4.99999523e+00
	// Value for iteration 22 is 4.99999762e+00
	// Value for iteration 23 is 4.99999881e+00
	// Value for iteration 24 is 4.99999940e+00
}

func TestFunc(t *testing.T) {
	var x float64
	if err := Run(&x, func() error {
		x += 1.0
		return fmt.Errorf("Internal error")
	}); err == nil {
		t.Errorf("cannot found internal error")
		return
	}
}

func TestMaxIter(t *testing.T) {
	var x float64
	if err := Run(&x, func() error {
		x += 1.0
		return nil
	}); err == nil {
		t.Errorf("cannot found max iter error")
		return
	}
}
