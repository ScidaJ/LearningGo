package iteration

import (
	"fmt"

	"testing"
)

func TestRepeat(t *testing.T) {
	assertExpectedResult := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("iteration test", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"

		assertExpectedResult(t, got, want)
	})
}

func ExampleRepeat() {
	fmt.Println(Repeat("b"))
	// Output: bbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
