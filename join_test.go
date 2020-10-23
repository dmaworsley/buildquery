package buildquery

import (
	"testing"
	"reflect"
)

func TestJoinSingleTable(t *testing.T) {

	query := Select()
	table := make(AliasItem)
	table["ad"] = "affiliate_details"
	onClause := "ad.user_id = a.user_id and ad.`key` = 'email'"
	columns:= make(AliasItem)
	columns["email"] = "`value`"
	joinType := "inner"


	join := joinItem{}
	join.joinTable = table
	join.onClause = onClause
	join.columns = columns
	join.joinType = joinType

	want := []joinItem{}
	want = append(want, join)

	query.Join(table, onClause, columns, joinType)

	got := query.SelectSQL.join

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("Join() Single Table = %q, want %q", got, want)
	}
}


func TestJoinMultiTables(t *testing.T) {

	query := Select()
	table := make(AliasItem)
	table["ad"] = "affiliate_details"
	onClause := "ad.user_id = a.user_id and ad.`key` = 'email'"
	columns:= make(AliasItem)
	columns["email"] = "`value`"
	joinType := "inner"


	join := joinItem{}
	join.joinTable = table
	join.onClause = onClause
	join.columns = columns
	join.joinType = joinType

	want := []joinItem{}
	want = append(want, join)
	want = append(want, join)

	query.Join(table, onClause, columns, joinType)
	//doesnt matter that its a copy - seperate elements
	query.Join(table, onClause, columns, joinType)

	got := query.SelectSQL.join

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("Join() Multi Table = %q, want %q", got, want)
	}

}
