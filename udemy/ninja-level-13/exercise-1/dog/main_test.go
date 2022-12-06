package dog

import (
	"fmt"
	"testing"
)

// go test
// go test -bench .
// go test -cover
// go test -coverprofile c.out
// go tool cover -html=c.out

func TestYears(t *testing.T) {
	got := Years(5)
	if got != 35 {
		t.Error("Expected", 35, "Got", got)
	}
}

func TestYearsTwo(t *testing.T) {
	got := YearsTwo(5)
	if got != 35 {
		t.Error("Expected", 35, "Got", got)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(5)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(5)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	// Output:
	// 35
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(5))
	// Output:
	// 35
}
