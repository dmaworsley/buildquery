package buildquery

import (
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {

	//Select returns a SelectBuilder struct - thatzeck
	want := &SelectBuilder{}
	got := Select()

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("Select() = %v, want %v", got, want)
	}

}

