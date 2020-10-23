package buildquery

import (
	"reflect"
	"testing"
)


func TestHavingSingleSet(t *testing.T) {

	query := Select()
	var want = ItemGroup{"COUNT(1) > 1"}
	query.Having(want)

	got := query.SelectSQL.having

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("Having() Single Call = %v, want %v", got, want)
	}
}


func TestHavingMultiSet(t *testing.T) {

	query := Select()
	var set1 = ItemGroup{"COUNT(1) > 1"}
	query.Having(set1)

	var set2 = ItemGroup{"user_id = 9000"}
	query.Having(set2)

	want := ItemGroup{"COUNT(1) > 1", "user_id = 9000"}

	got := query.SelectSQL.having

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("Having() Multi Call = %v, want %v", got, want)
	}
}
