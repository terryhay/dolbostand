package float_numbers

import (
	"fmt"
	"testing"
)

func TestFloatNumbers(t *testing.T) {
	var f1 float64 = 0.2
	fmt.Printf("%E\n", f1)

	var f2 float32 = 2.
	var f3 float32 = 10.
	fmt.Printf("%E\n", f2/f3)

	fmt.Printf("%v\n", f1 == float64(f2/f3))
}
