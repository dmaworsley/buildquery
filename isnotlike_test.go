package buildquery

import (
	"testing"
	"reflect"
)

func TestIsNotLikeSingleSet(t *testing.T) {

	query := Select()
	field := "username"
	value := "casino%"

	isLike := likeSelection{}
	isLike.field = field
	isLike.value = value

	want := []likeSelection{}
	want = append(want, isLike)

	query.IsNotLike(field, value)
	got := query.SelectSQL.isNotLike

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("IsNotLike() Single Set = %q, want %q", got, want)
	}
}

func TestIsNotLikeMultiSet(t *testing.T) {

	query := Select()
	field_1 := "username"
	value_1 := "casino%"

	isLike_1 := likeSelection{}
	isLike_1.field = field_1
	isLike_1.value = value_1

	want := []likeSelection{}
	want = append(want, isLike_1)
	want = append(want, isLike_1)

	query.IsNotLike(field_1, value_1)
	query.IsNotLike(field_1, value_1)
	got := query.SelectSQL.isNotLike

	equivilent := reflect.DeepEqual(want, got)

	if !equivilent {
		t.Errorf("IsNotLike() Single Set = %q, want %q", got, want)
	}
}





