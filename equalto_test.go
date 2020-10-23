package buildquery

import (
	"reflect"
	"testing"
)

func TestEqualToSingleCall(t *testing.T) {

	query := Select()

	want := make(AliasItem)
	want["user_id"] = "a.user_id"
	want["username"] = "a.username"
	want["join_date"] = "a.join_date"

	query.EqualTo(want)

	got := query.SelectSQL.equalTo

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("EqualTo() Single Call = %q, dont want %q", got, want)
	}
}


func TestEqualToMultiCall(t *testing.T) {

	query := Select()

	cols_inital := make(AliasItem)
	cols_inital["user_id"] = "a.user_id"
	cols_inital["username"] = "a.username"
	cols_inital["join_date"] = "a.join_date"

	//add first set cols
	query.EqualTo(cols_inital)

	cols_again := make(AliasItem)

	cols_again["salt"] = "a.salt"
	cols_again["password"] = "a.password"

	//add more columns
	query.EqualTo(cols_again)

	want  := make(AliasItem)
	want["join_date"] = "a.join_date"
	want["password"] = "a.password"
	want["salt"] = "a.salt"
	want["user_id"] = "a.user_id"
	want["username"] = "a.username"

	got := query.SelectSQL.equalTo

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("EqualTo() Multi Call = %q, dont want %q", got, want)
	}
}
