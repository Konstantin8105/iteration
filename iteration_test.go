package iteration

import (
	"fmt"
	"math"
	"os"
	"testing"
)

// cpu: Intel(R) Xeon(R) CPU           X5550  @ 2.67GHz
// Benchmark/single-8         	 3743305	       305.4 ns/op	       8 B/op	       1 allocs/op
// Benchmark/two-8            	 1902074	       618.8 ns/op	      16 B/op	       1 allocs/op
func Benchmark(b *testing.B) {
	b.Run("single", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var x float64
			if err := Find(func() error {
				x = 5
				return nil
			}, &x); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("two", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var x, y float64
			if err := Find(func() error {
				y = 1 + x
				x = 5
				return nil
			}, &x, &y); err != nil {
				b.Fatal(err)
			}
		}
	})
}

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

func TestErrorType(t *testing.T) {
	t.Run("some error", func(t *testing.T) {
		var x float64
		err := Find(func() error {
			return fmt.Errorf("some error")
		}, &x)
		if err == nil {
			t.Errorf("have not errors")
		}
		fmt.Fprintf(os.Stdout, "%v\n", err)
	})
	t.Run("infinite1", func(t *testing.T) {
		var x float64
		err := Find(func() error {
			x = math.Inf(1)
			return nil
		}, &x)
		if err == nil {
			t.Errorf("have not errors")
		}
		fmt.Fprintf(os.Stdout, "%v\n", err)
	})
	t.Run("infinite0", func(t *testing.T) {
		var x float64
		err := Find(func() error {
			x = math.Inf(0)
			return nil
		}, &x)
		if err == nil {
			t.Errorf("have not errors")
		}
		fmt.Fprintf(os.Stdout, "%v\n", err)
	})
	t.Run("infinite-1", func(t *testing.T) {
		var x float64
		err := Find(func() error {
			x = math.Inf(-1)
			return nil
		}, &x)
		if err == nil {
			t.Errorf("have not errors")
		}
		fmt.Fprintf(os.Stdout, "%v\n", err)
	})
	t.Run("NaN", func(t *testing.T) {
		var x float64
		err := Find(func() error {
			x = math.NaN()
			return nil
		}, &x)
		if err == nil {
			t.Errorf("have not errors")
		}
		fmt.Fprintf(os.Stdout, "%v\n", err)
	})
	t.Run("infinite", func(t *testing.T) {
		var x float64
		var b bool
		err := Find(func() error {
			if b {
				x += 1
			} else {
				x -= 1
			}
			b = !b
			return nil
		}, &x)
		if err == nil {
			t.Errorf("have not errors")
		}
		fmt.Fprintf(os.Stdout, "%v\n", err)
	})
}
