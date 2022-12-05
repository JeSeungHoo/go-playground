// We are going to learn about testing next. For this exercise, take a moment and see how much
// you can figure out about testing by reading the testing documentation & also Caleb Doxsey’s
// article on testing. See if you can get a basic example of testing working.

package exercise5

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	got := math.Abs(-123)
	if got != 123 {
		t.Errorf("Abs(-123) = %v; want 123", got)
	}
}
