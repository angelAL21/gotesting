package suma

import "testing"

//name of Test
func TestAdd(t *testing.T) {
	want := 4
	t.Logf("want is %d", want)
	got := Add(1, 3)
	t.Logf("got is %d", got)
	if got != want {
		t.Logf("Error: se esperaba %d, se obtuvo %d", want, got)
		t.Fail()
	}
	t.Log("terminó la prueba add")
}

//name of test
func TestAddAcumulate(t *testing.T) {
	want := 6
	//want := 7
	got := AddAcumulate(1, 2, 3)

	if got != want {
		t.Logf("Error: se git pesperaba %d, se obtuvo %d", want, got)
		t.Fail()
	}
}
