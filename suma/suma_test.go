package suma

import (
	"testing"
)

//name of Test
func TestAdd(t *testing.T) {
	want := 5
	t.Logf("want is %d", want)
	got := Add(1, 3)
	t.Logf("got is %d", got)
	if got != want {
		t.Logf("Error: se esperaba %d, se obtuvo %d", want, got)
		//t.Fail will fail when there is a error
		t.Fail()
	}
	t.Log("terminó la prueba add")
}

//name of test
func TestAddAcumulate(t *testing.T) {
	want := 6
	got := AddAcumulate(1, 2, 3)

	if got != want {
		//combination of log and fail.
		t.Errorf("Error: se pesperaba %d, se obtuvo %d", want, got)
	}
}
