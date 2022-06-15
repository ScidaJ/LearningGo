package integers

import "testing"

func TestAdder(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("positive addition", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertCorrectMessage(t, sum, expected)
	})

	t.Run("negative addition", func(t *testing.T) {
		sum := Add(-5, -7)
		expected := -12

		assertCorrectMessage(t, sum, expected)
	})

	t.Run("negative/positive addition", func(t *testing.T) {
		sum := Add(-5, 7)
		expected := 2

		assertCorrectMessage(t, sum, expected)
	})
}
