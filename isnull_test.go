package buildquery

import (
	"reflect"
	"testing"
)


func TestIsNullSingleSet(t *testing.T) {

	query := Select()
	var want = ItemGroup{"1", "2", "3"}
	query.IsNull(want)

	got := query.SelectSQL.isNull

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("IsNull() Single Call = %q, want %q", got, want)
	}
}


func TestIsNullMultiSet(t *testing.T) {

	query := Select()
	var set1 = ItemGroup{"1", "2", "3"}
	query.IsNull(set1)

	var set2 = ItemGroup{"4", "5", "6"}
	query.IsNull(set2)

	want := ItemGroup{"1", "2", "3", "4", "5", "6"}

	got := query.SelectSQL.isNull

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("IsNull() Multi Call = %q, want %q", got, want)
	}
}
