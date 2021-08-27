package suma

import "testing"

//name of Test
func TestAdd(t *testing.T) {
	want := 4
	got := Add(1, 3)
	if got != want {
		t.Logf("Error: se esperaba %d, se obtuvo %d", want, got)
		t.Fail()
	}
}

//name of test
func TestAcumulate(t *testing.T) {
	want := 7
	//want := 7
	got := AddAcumulate(1, 2, 3)

	if got != want {
		t.Logf("Error: se esperaba %d, se obtuvo %d", want, got)
		t.Fail()
	}
}
