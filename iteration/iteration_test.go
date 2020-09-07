package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("A", 3)
	expected := "AAA"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("B", 6)
	fmt.Println(repeated)
	// Output: BBBBBB
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("C", 5)
	}
}
