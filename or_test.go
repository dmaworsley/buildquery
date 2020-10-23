package buildquery

import (
	"reflect"
	"testing"
)

func TestSingleOrSingleSet(t *testing.T) {

	query := Select()
	orItem_1 := ItemGroup{"user_id < 9000"}

	want :=[]ItemGroup{}
	want = append(want, orItem_1)

	query.Or(orItem_1)
	got := query.SelectSQL.orGroup

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("In() Single Set Single Item = %v, want %v	", got, want)
	}
}

func TestSingleOrMultiSet(t *testing.T) {

	query := Select()
	orItem_1 := ItemGroup{"user_id < 9000"}
	orItem_2 := ItemGroup{"user_id > 20000"}

	want :=[]ItemGroup{}
	want = append(want, orItem_1)
	want = append(want, orItem_2)

	query.Or(orItem_1)
	query.Or(orItem_2)

	got := query.SelectSQL.orGroup

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("In() Single Set Multi Item = %v, want %v	", got, want)
	}
}

func TestDoubleOrSingleSet(t *testing.T) {

	query := Select()
	orItem_1 := ItemGroup{"user_id = 1", "user_id = 5"}

	want :=[]ItemGroup{}
	want = append(want, orItem_1)

	query.Or(orItem_1)
	got := query.SelectSQL.orGroup

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("In() Single Set Single Item = %v, want %v	", got, want)
	}
}

func TestDoubleOrMultiSet(t *testing.T) {

	query := Select()
	orItem_1 := ItemGroup{"user_id = 1", "user_id = 5"}
	orItem_2 := ItemGroup{"user_id = 12", "user_id = 23"}

	want :=[]ItemGroup{}
	want = append(want, orItem_1)
	want = append(want, orItem_2)

	query.Or(orItem_1)
	query.Or(orItem_2)

	got := query.SelectSQL.orGroup

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("In() Single Set Multi Item = %v, want %v	", got, want)
	}
}


