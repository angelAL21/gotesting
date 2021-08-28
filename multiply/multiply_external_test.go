package multiply_test

import (
	"testing"

	"github.com/angelAL21/gotesting/multiply"
)

func TestMultiply(t *testing.T) {
	want := 6
	//multiply is our package, Multiply is our func
	got := multiply.Multiply(2, 3)
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

//simulando la importancion de un paquete externo
