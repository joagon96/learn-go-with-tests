package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, repeated string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeat 5 default times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("repeat char 3 times", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		assertCorrectMessage(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
