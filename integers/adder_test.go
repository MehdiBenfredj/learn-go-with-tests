package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Test integers package", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertEquals(t, sum, expected)
	})
}

func assertEquals(t testing.TB, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("expected: %q, but sum:  %q", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
