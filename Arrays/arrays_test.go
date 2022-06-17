package arrays

import (
	"reflect"

	"testing"
)

func TestSum(t *testing.T) {
	assertExpectedResult := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("testing slice summing with Sum()", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		got := Sum(arr)
		want := 15

		assertExpectedResult(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	assertExpectedResult := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("SumAll with two slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assertExpectedResult(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	assertExpectedResult := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("SumAllTails with two slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertExpectedResult(t, got, want)
	})

	t.Run("SumAllTails with slices of longer length", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{9, 8, 7, 6, 5})
		want := []int{4, 5}

		assertExpectedResult(t, got, want)
	})
}
