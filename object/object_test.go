package object

import (
	"reflect"
	"testing"
)

func TestPerro(t *testing.T) {
	want := &Perro{
		Name: "Firulais",
		Age:  1,
		Kind: Kind{
			Name: "electrico",
		},
	}
	got := &Perro{
		Name: "Firulais",
		Age:  1,
		Kind: Kind{
			Name: "electrico",
		},
	}

	// t.Logf("want %p, got %p", want, got)
	//if want != got {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Se esperaba: %v, se obtuvo %v", want, got)
	}
}

//deepEqual ayuda a comparar el contenido de puntero, estructura, arrays, slice, etc.
