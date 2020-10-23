package buildquery

import (
	"testing"
	"reflect"
)

func TestInSingleSet(t *testing.T) {

	query := Select()
	field := "user_id"
	values := ItemGroup{"1", "2", "3", "4", "5"}

	in := inSelection{}
	in.field = field
	in.values = values

	want := []inSelection{}
	want = append(want, in)

	query.In(field, values)
	got := query.SelectSQL.in

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("In() Single Set = %q, want %q", got, want)
	}
}

func TestInMultiSet(t *testing.T) {

	query := Select()
	field := "user_id"
	values := ItemGroup{"1", "2", "3", "4", "5"}

	in := inSelection{}
	in.field = field
	in.values = values

	want := []inSelection{}
	want = append(want, in)
	want = append(want, in)

	query.In(field, values)
	query.In(field, values)
	got := query.SelectSQL.in

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("In() Multi Set = %q, want %q", got, want)
	}
}





