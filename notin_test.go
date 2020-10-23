package buildquery

import (
	"testing"
	"reflect"
)

func TestNotInSingleSet(t *testing.T) {

	query := Select()
	field := "user_id"
	values := ItemGroup{"1", "2", "3", "4", "5"}

	notIn := inSelection{}
	notIn.field = field
	notIn.values = values

	want := []inSelection{}
	want = append(want, notIn)

	query.NotIn(field, values)
	got := query.SelectSQL.notIn

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("NotIn() Single Set = %q, want %q", got, want)
	}
}

func TestNotInMultiSet(t *testing.T) {

	query := Select()
	field := "user_id"
	values := ItemGroup{"1", "2", "3", "4", "5"}

	notIn := inSelection{}
	notIn.field = field
	notIn.values = values

	want := []inSelection{}
	want = append(want, notIn)
	want = append(want, notIn)

	query.NotIn(field, values)
	query.NotIn(field, values)
	got := query.SelectSQL.notIn

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("NotIn() Multi Set = %q, want %q", got, want)
	}
}





