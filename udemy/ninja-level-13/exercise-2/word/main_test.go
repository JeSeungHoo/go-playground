package word

import (
	"fmt"
	"testing"

	"gitgub.com/JeSeungHoo/go-playground/udemy/ninja-level-13/exercise-2/quote"
)

func TestUseCount(t *testing.T) {
	got := UseCount("one two three two three three")
	for k, v := range got {
		var ex int
		switch k {
		case "one":
			ex = 1
		case "two":
			ex = 2
		case "three":
			ex = 3
		}
		if v != ex {
			t.Errorf("Expected %v, got %v", ex, v)
		}
	}
}
func TestCount(t *testing.T) {
	got := Count("one two three")
	if got != 3 {
		t.Error("Expected 3, got", got)
	}
}
func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
