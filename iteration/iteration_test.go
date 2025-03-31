package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	get := Repeat("a", 5)
	expect := "aaaaa"

	if get != expect {
		t.Errorf("Expected %q but got %q", expect, get)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	res := Repeat("a", 5)
	fmt.Println(res)
	// Output: aaaaa
}
