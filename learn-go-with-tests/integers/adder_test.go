package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, sum int) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}
	t.Run("2 + 2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertCorrectMessage(t, expected, sum)
	})
	t.Run("3 + 5", func(t *testing.T) {
		sum := Add(3, 5)
		expected := 8

		assertCorrectMessage(t, expected, sum)
	})
}
