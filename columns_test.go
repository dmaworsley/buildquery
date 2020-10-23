package buildquery

import (
	"reflect"
	"testing"
)

func TestColumnsSingleCall(t *testing.T) {

	query := Select()

	want := make(AliasItem)
	want["user_id"] = "a.user_id"
	want["username"] = "a.username"
	want["join_date"] = "a.join_date"

	query.Columns(want)

	got := query.SelectSQL.columns

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("Columns() Single Call = %q, dont want %q", got, want)
	}
}


func TestColumnsMultiCall(t *testing.T) {
	query := Select()

	cols_inital := make(AliasItem)
	cols_inital["user_id"] = "a.user_id"
	cols_inital["username"] = "a.username"
	cols_inital["join_date"] = "a.join_date"

	//add first set cols
	query.Columns(cols_inital)

	cols_again := make(AliasItem)

	cols_again["salt"] = "a.salt"
	cols_again["password"] = "a.password"

	//add more columns
	query.Columns(cols_again)

	want  := make(AliasItem)
	want["join_date"] = "a.join_date"
	want["password"] = "a.password"
	want["salt"] = "a.salt"
	want["user_id"] = "a.user_id"
	want["username"] = "a.username"

	got := query.SelectSQL.columns

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("Columns() Multi Call = %q, dont want %q", got, want)
	}
}
