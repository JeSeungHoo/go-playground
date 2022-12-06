package mymath

import (
	"fmt"
	"testing"
)

type test struct {
	data   []int
	answer float64
}

func TestCenteredAvg(t *testing.T) {
	tests := []test{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{6, 7, 8, 9, 10}, 8},
		{[]int{236, 6523, 124356, 562, 456, 3456}, 2749.25},
		{[]int{15, 96, 2, 95, 9, 91}, 52.5},
	}
	for _, v := range tests {
		got := CenteredAvg(v.data)
		ex := v.answer
		if got != ex {
			t.Errorf("Expected %v, Got %v", ex, got)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 3, 6, 7, 9, 10, 54}))
	// Output:
	// 7
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{12, 234, 523, 46, 1243, 67, 454, 12323, 412566, 43})
	}
}
