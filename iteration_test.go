package iteration

import (
	"fmt"
	"os"
	"testing"
)

func Example() {
	var x float64
	var counter int
	if err := Find(func() error {
		counter++
		fmt.Fprintf(os.Stdout, "Value for iteration %2d is %10.8e\n", counter, x)
		x = 5
		return nil
	}, &x); err != nil {
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

func ExampleFind() {
	var x, y float64
	var counter int
	if err := Find(func() error {
		counter++
		fmt.Fprintf(os.Stdout, "Value for iteration %2d is {%10.8e,%10.8e}\n",
			counter, x, y)
		y = 1 + x
		x = 5
		return nil
	}, &x, &y); err != nil {
		fmt.Fprintf(os.Stdout, "Error: %v", err)
		return
	}

	// Output:
	// Value for iteration  1 is {0.00000000e+00,0.00000000e+00}
	// Value for iteration  2 is {3.09016994e+00,6.18033989e-01}
	// Value for iteration  3 is {4.27050983e+00,2.76393202e+00}
	// Value for iteration  4 is {4.72135955e+00,4.31308230e+00}
	// Value for iteration  5 is {4.89356882e+00,5.18344551e+00}
	// Value for iteration  6 is {4.95934691e+00,5.62232585e+00}
	// Value for iteration  7 is {4.98447190e+00,5.83061632e+00}
	// Value for iteration  8 is {4.99406879e+00,5.92570430e+00}
	// Value for iteration  9 is {4.99773448e+00,5.96795588e+00}
	// Value for iteration 10 is {4.99913465e+00,5.98636007e+00}
	// Value for iteration 11 is {4.99966947e+00,5.99425519e+00}
	// Value for iteration 12 is {4.99987375e+00,5.99760140e+00}
	// Value for iteration 13 is {4.99995178e+00,5.99900579e+00}
	// Value for iteration 14 is {4.99998158e+00,5.99959044e+00}
	// Value for iteration 15 is {4.99999296e+00,5.99983218e+00}
	// Value for iteration 16 is {4.99999731e+00,5.99993155e+00}
	// Value for iteration 17 is {4.99999897e+00,5.99997219e+00}
	// Value for iteration 18 is {4.99999961e+00,5.99998874e+00}
	// Value for iteration 19 is {4.99999985e+00,5.99999546e+00}
}

func TestFunc(t *testing.T) {
	var x float64
	if err := Find(func() error {
		x += 1.0
		return fmt.Errorf("Internal error")
	}, &x); err == nil {
		t.Errorf("cannot found internal error")
		return
	}
}

func TestMaxIter(t *testing.T) {
	var x float64
	if err := Find(func() error {
		x += 1.0
		return nil
	}, &x); err == nil {
		t.Errorf("cannot found max iter error")
		return
	}
}

func TestWrong(t *testing.T) {
	var x float64
	if err := FindWithOption(func() error {
		x += 1
		return nil
	}, Option{
		MaxIteration: 0,
		Ratio:        0.1,
		Precision:    0.1,
	}, &x); err == nil {
		t.Fatalf("not valid max iteration")
	}
	if err := FindWithOption(func() error {
		x += 1
		return nil
	}, Option{
		MaxIteration: 10,
		Ratio:        0,
		Precision:    0.1,
	}, &x); err == nil {
		t.Fatalf("not valid ratio")
	}
	if err := FindWithOption(func() error {
		x += 1
		return nil
	}, Option{
		MaxIteration: 10,
		Ratio:        0.1,
		Precision:    0,
	}, &x); err == nil {
		t.Fatalf("not valid ratio")
	}
	if err := FindWithOption(nil, Option{
		MaxIteration: 10,
		Ratio:        0.1,
		Precision:    0.1,
	}, &x); err == nil {
		t.Fatalf("not valid function")
	}
}
