package buildquery

import (
	"reflect"
	"testing"
)


func TestIsNotNullSingleSet(t *testing.T) {

	query := Select()
	var want = ItemGroup{"1", "2", "3"}
	query.IsNotNull(want)

	got := query.SelectSQL.isNotNull

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("IsNotNull() Single Call = %q, want %q", got, want)
	}
}


func TestIsNotNullMultiSet(t *testing.T) {

	query := Select()
	var set1 = ItemGroup{"1", "2", "3"}
	query.IsNotNull(set1)

	var set2 = ItemGroup{"4", "5", "6"}
	query.IsNotNull(set2)

	want := ItemGroup{"1", "2", "3", "4", "5", "6"}

	got := query.SelectSQL.isNotNull

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("IsNotNull() Multi Call = %q, want %q", got, want)
	}
}
