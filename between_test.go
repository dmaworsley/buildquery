package buildquery

import (
	"reflect"
	"testing"
)

func TestBetweenSingleSet(t *testing.T) {

	query := Select()
	field := "`date``"
	min := "2020-09-01"
	max := "2020-09-30"

	between := betweenSelection{}
	between.field = field
	between.min = min
	between.max = max

	want := []betweenSelection{}
	want = append(want, between)

	query.Between(field, min, max)
	got := query.SelectSQL.between

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("Between() Single Set = %q, want %q", got, want)
	}
}


func TestBetweenMultiSet(t *testing.T) {

	query := Select()
	field := "`date``"
	min := "2020-09-01"
	max := "2020-09-30"

	between := betweenSelection{}
	between.field = field
	between.min = min
	between.max = max

	want := []betweenSelection{}
	want = append(want, between)
	want = append(want, between)

	query.Between(field, min, max)
	query.Between(field, min, max)
	got := query.SelectSQL.between

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("Between() Multi Set = %q, want %q", got, want)
	}
}
