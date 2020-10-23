package buildquery

import (
	"reflect"
	"testing"
)


func TestFrom(t *testing.T) {

	want_first := make(AliasItem)
	want_first["t"] = "testing"
	query := Select()
	query.From(want_first)

	got_first := query.SelectSQL.from

	equivilent := reflect.DeepEqual(want_first, got_first)

	if !equivilent {
		t.Errorf("From() = %q, want %q", got_first, want_first)
	}

	//multi element maps could be supplied to FROM() accidentally. We only care about 1st elem only - ever
	dont_want := make(AliasItem)
	dont_want["t"] = "testing"
	dont_want["x"] = "xtreme"

	newQuery := Select()
	newQuery.From(dont_want)

	got_second := newQuery.SelectSQL.from
	equivilent_again := reflect.DeepEqual(want_first, got_second)

	if !equivilent_again {
		t.Errorf("From() = %q, want %q", got_second, want_first)
	}

	//4 out of 10 for this

}
