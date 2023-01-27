package array_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, want, got int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("expected '%d' but got '%d', %v", want, got, numbers)
		}
	}

	t.Run("sum 5 numbers from array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectMessage(t, want, got, numbers)
	})

	t.Run("sum 5 other numbers from array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 1}

		got := Sum(numbers)
		want := 11

		assertCorrectMessage(t, want, got, numbers)
	})

	t.Run("sum numbers from array of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, want, got, numbers)
	})
}

func TestSumAll(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, want, got []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got)
		}
	}

	t.Run("sum 2 arrays content", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4, 5}
		numbers2 := []int{1, 2, 3}

		got := SumAll(numbers1, numbers2)
		want := []int{15, 6}

		assertCorrectMessage(t, want, got)
	})

	t.Run("sum 1 arrays content", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4, 5}

		got := SumAll(numbers1)
		want := []int{15}

		assertCorrectMessage(t, want, got)
	})

	t.Run("sum 3 arrays content", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4, 5}
		numbers2 := []int{1, 2}
		numbers3 := []int{1, 2, 3}

		got := SumAll(numbers1, numbers2, numbers3)
		want := []int{15, 3, 6}

		assertCorrectMessage(t, want, got)
	})
}

func TestSumAllTails(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, want, got []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got)
		}
	}

	t.Run("sum 2 arrays tails", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4, 5}
		numbers2 := []int{1, 2, 3}

		got := SumAllTails(numbers1, numbers2)
		want := []int{14, 5}

		assertCorrectMessage(t, want, got)
	})

	t.Run("sum 1 arrays tails", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4, 5}

		got := SumAllTails(numbers1)
		want := []int{14}

		assertCorrectMessage(t, want, got)
	})

	t.Run("sum 3 arrays tails", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4, 5}
		numbers2 := []int{1, 2, 3}
		numbers3 := []int{1, 2}

		got := SumAllTails(numbers1, numbers2, numbers3)
		want := []int{14, 5, 2}

		assertCorrectMessage(t, want, got)
	})

	t.Run("sum an array with no tail", func(t *testing.T) {
		numbers := []int{1}

		got := SumAllTails(numbers)
		want := []int{0}

		assertCorrectMessage(t, want, got)
	})

	t.Run("sum an empty array", func(t *testing.T) {
		var numbers []int

		got := SumAllTails(numbers)
		want := []int{0}

		assertCorrectMessage(t, want, got)
	})
}
