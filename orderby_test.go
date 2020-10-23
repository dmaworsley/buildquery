package buildquery

import (
	"reflect"
	"testing"
)

func TestOrderBySingleSet(t *testing.T) {

	query := Select()
	asc := true
	field := "user_id"

	orderBy := orderSelection{}
	orderBy.asc = asc
	orderBy.field = field

	want := []orderSelection{}
	want = append(want, orderBy)

	query.OrderBy(asc, field)
	got := query.SelectSQL.orderBy

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("OrderBy() Single Set = %v, want %v	", got, want)
	}
}

func TestOrderByMultiSet(t *testing.T) {

	query := Select()
	asc := true
	field := "user_id"

	orderBy := orderSelection{}
	orderBy.asc = asc
	orderBy.field = field

	want := []orderSelection{}
	want = append(want, orderBy)
	want = append(want, orderBy)

	query.OrderBy(asc, field)
	query.OrderBy(asc, field)
	got := query.SelectSQL.orderBy

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("OrderBy() Single Set = %v, want %v	", got, want)
	}

}
