package buildquery

import (
	"reflect"
	"testing"
)


func TestGroupBySingleSet(t *testing.T) {

	query := Select()
	var want = ItemGroup{"1", "2"}
	query.GroupBy(want)

	got := query.SelectSQL.groupBy

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("GroupBy() Single Call = %v, want %v", got, want)
	}
}


func TestGroupByMultiSet(t *testing.T) {

	query := Select()
	var set1 = ItemGroup{"1", "2"}
	query.GroupBy(set1)

	var set2 = ItemGroup{"3", "4"}
	query.GroupBy(set2)

	want := ItemGroup{"1", "2", "3", "4"}

	got := query.SelectSQL.groupBy

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("GroupBy() Multi Call = %v, want %v", got, want)
	}
}
