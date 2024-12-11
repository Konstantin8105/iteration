package iteration

import (
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/Konstantin8105/compare"
)

// cpu: Intel(R) Xeon(R) CPU           X5550  @ 2.67GHz
// Benchmark/single-8         	 3743305	       305.4 ns/op	       8 B/op	       1 allocs/op
// Benchmark/two-8            	 1902074	       618.8 ns/op	      16 B/op	       1 allocs/op
//
// cpu: Intel(R) Xeon(R) CPU E3-1240 V2 @ 3.40GHz
// Benchmark/single-6         	 5913253	       198.9 ns/op	       8 B/op	       1 allocs/op
// Benchmark/two-6            	 3717120	       320.6 ns/op	      16 B/op	       1 allocs/op
//
// Benchmark/single-6         	 2873384	       426.1 ns/op	       8 B/op	       1 allocs/op
// Benchmark/two-6            	 1854848	       611.7 ns/op	      16 B/op	       2 allocs/op
func Benchmark(b *testing.B) {
	b.Run("single", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var x float64
			if err := Find(func() error {
				x = 5
				return nil
			}, []*float64{&x}, []*float64{}, []*float64{}); err != nil {
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
			}, []*float64{&x}, []*float64{&y}, []*float64{}); err != nil {
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
	}, []*float64{&x}, []*float64{}, []*float64{}); err != nil {
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
	}, []*float64{&x, &y}, []*float64{}, []*float64{}); err != nil {
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

func Test(t *testing.T) {
	eps := 1e-5
	t.Run("float64", func(t *testing.T) {
		var x, y float64
		if err := Find(func() error {
			y = 1 + x
			x = 5
			return nil
		}, []*float64{&x}, []*float64{&y}, []*float64{}); err != nil {
			t.Fatal(err)
		}
		if e := math.Abs((float64(x) - 5) / 5); eps < e {
			t.Errorf("not valid root")
		}
		if e := math.Abs((float64(y) - 6) / 6); eps < e {
			t.Errorf("not valid root")
		}
	})
	t.Run("F64-1", func(t *testing.T) {
		type F64 float64
		var x, y F64
		if err := Find(func() error {
			y = 1 + x
			x = 5
			return nil
		}, []*F64{&x, &y}, []*float64{}, []*float64{}); err != nil {
			t.Fatal(err)
		}
		if e := math.Abs((float64(x) - 5) / 5); eps < e {
			t.Errorf("not valid root")
		}
		if e := math.Abs((float64(y) - 6) / 6); eps < e {
			t.Errorf("not valid root")
		}
	})
	t.Run("F64-2", func(t *testing.T) {
		type F64 float64
		var x, y F64
		if err := Find(func() error {
			y = 1 + x
			x = 5
			return nil
		}, []*F64{}, []*F64{&x, &y}, []*float64{}); err != nil {
			t.Fatal(err)
		}
		if e := math.Abs((float64(x) - 5) / 5); eps < e {
			t.Errorf("not valid root")
		}
		if e := math.Abs((float64(y) - 6) / 6); eps < e {
			t.Errorf("not valid root")
		}
	})
	t.Run("F64-3", func(t *testing.T) {
		type F64 float64
		var x, y F64
		if err := Find(func() error {
			y = 1 + x
			x = 5
			return nil
		}, []*F64{}, []*float64{}, []*F64{&x, &y}); err != nil {
			t.Fatal(err)
		}
		if e := math.Abs((float64(x) - 5) / 5); eps < e {
			t.Errorf("not valid root")
		}
		if e := math.Abs((float64(y) - 6) / 6); eps < e {
			t.Errorf("not valid root")
		}
	})
	t.Run("float64+F64", func(t *testing.T) {
		type F64 float64
		var x float64
		var y F64
		if err := Find(func() error {
			y = 1 + F64(x)
			x = 5
			return nil
		}, []*float64{&x}, []*F64{&y}, []*float64{}); err != nil {
			t.Fatal(err)
		}
		if e := math.Abs((float64(x) - 5) / 5); eps < e {
			t.Errorf("not valid root")
		}
		if e := math.Abs((float64(y) - 6) / 6); eps < e {
			t.Errorf("not valid root")
		}
	})
	t.Run("max iteration", func(t *testing.T) {
		var x float64
		step := func() error {
			x += 1.0e-5
			return nil
		}
		err := Find(step, []*float64{&x}, []*float64{}, []*float64{})
		compare.Test(t, ".max.iteration",[]byte( fmt.Sprintf("%v", err)))
		if err == nil {
			t.Errorf("cannot found max iter error")
			return
		}
	})
}

func TestFunc(t *testing.T) {
	var x float64
	if err := Find(func() error {
		x += 1.0
		return fmt.Errorf("Internal error")
	}, []*float64{&x}, []*float64{}, []*float64{}); err == nil {
		t.Errorf("cannot found internal error")
		return
	}
}

func TestMaxIter(t *testing.T) {
	var x float64
	if err := Find(func() error {
		x += 1.0
		return nil
	}, []*float64{&x}, []*float64{}, []*float64{}); err == nil {
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
	}, []*float64{&x}, []*float64{}, []*float64{}); err == nil {
		t.Fatalf("not valid max iteration")
	}
	if err := FindWithOption(func() error {
		x += 1
		return nil
	}, Option{
		MaxIteration: 10,
		Ratio:        0,
		Precision:    0.1,
	}, []*float64{&x}, []*float64{}, []*float64{}); err == nil {
		t.Fatalf("not valid ratio")
	}
	if err := FindWithOption(func() error {
		x += 1
		return nil
	}, Option{
		MaxIteration: 10,
		Ratio:        0.1,
		Precision:    0,
	}, []*float64{&x}, []*float64{}, []*float64{}); err == nil {
		t.Fatalf("not valid ratio")
	}
	if err := FindWithOption(nil, Option{
		MaxIteration: 10,
		Ratio:        0.1,
		Precision:    0.1,
	}, []*float64{&x}, []*float64{}, []*float64{}); err == nil {
		t.Fatalf("not valid function")
	}
}

func TestErrorType(t *testing.T) {
	t.Run("some error", func(t *testing.T) {
		var x float64
		err := Find(func() error {
			return fmt.Errorf("some error")
		}, []*float64{&x}, []*float64{}, []*float64{})
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
		}, []*float64{&x}, []*float64{}, []*float64{})
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
		}, []*float64{&x}, []*float64{}, []*float64{})
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
		}, []*float64{&x}, []*float64{}, []*float64{})
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
		}, []*float64{&x}, []*float64{}, []*float64{})
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
		}, []*float64{&x}, []*float64{}, []*float64{})
		if err == nil {
			t.Errorf("have not errors")
		}
		fmt.Fprintf(os.Stdout, "%v\n", err)
	})
}
