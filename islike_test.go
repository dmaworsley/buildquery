package buildquery

import (
	"testing"
	"reflect"
)

func TestIsLikeSingleSet(t *testing.T) {

	query := Select()
	field := "username"
	value := "casino%"

	isLike := likeSelection{}
	isLike.field = field
	isLike.value = value

	want := []likeSelection{}
	want = append(want, isLike)

	query.IsLike(field, value)
	got := query.SelectSQL.isLike

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("IsLike() Single Set = %q, want %q", got, want)
	}
}

func TestIsLikeMultiSet(t *testing.T) {

	query := Select()
	field_1 := "username"
	value_1 := "casino%"

	isLike_1 := likeSelection{}
	isLike_1.field = field_1
	isLike_1.value = value_1

	want := []likeSelection{}
	want = append(want, isLike_1)
	want = append(want, isLike_1)

	query.IsLike(field_1, value_1)
	query.IsLike(field_1, value_1)
	got := query.SelectSQL.isLike

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("IsLike() Single Set = %q, want %q", got, want)
	}
}





