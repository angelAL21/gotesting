package suma

import "testing"

func TestAdd(t *testing.T) {
	want := 3
	got := Add(1, 3)
	if got != want {
		t.Logf("Error: se esperaba %d, se obtuvo %d", want, got)
		t.Fail()
	}
}
