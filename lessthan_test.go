package buildquery

import (
	"reflect"
	"testing"
)

func TestLessThanSingleCall(t *testing.T) {

	query := Select()

	want := make(AliasItem)
	want["user_id"] = "a.user_id"
	want["username"] = "a.username"
	want["join_date"] = "a.join_date"

	query.LessThan(want)

	got := query.SelectSQL.lessThan

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("LessThan() Single Call = %q, dont want %q", got, want)
	}
}


func TestLessThanMultiCall(t *testing.T) {

	query := Select()

	cols_inital := make(AliasItem)
	cols_inital["user_id"] = "a.user_id"
	cols_inital["username"] = "a.username"
	cols_inital["join_date"] = "a.join_date"

	//add first set cols
	query.LessThan(cols_inital)

	cols_again := make(AliasItem)

	cols_again["salt"] = "a.salt"
	cols_again["password"] = "a.password"

	//add more columns
	query.LessThan(cols_again)

	want  := make(AliasItem)
	want["join_date"] = "a.join_date"
	want["password"] = "a.password"
	want["salt"] = "a.salt"
	want["user_id"] = "a.user_id"
	want["username"] = "a.username"

	got := query.SelectSQL.lessThan

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("LessThan() Multi Call = %q, dont want %q", got, want)
	}
}
