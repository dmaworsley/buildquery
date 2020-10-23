package buildquery

type SelectBuilder struct {
	SelectSQL
}

type AliasItem map[string]string
type values interface{}
type ItemGroup []string

type joinItem struct {
	joinTable AliasItem
	onClause string
	columns AliasItem
	joinType string
}

type where struct {
	field string
	value string
}

type betweenSelection struct {
	field string
	min string
	max string
}

type likeSelection struct {
	field string
	value string
}

type inSelection struct {
	field string
	values ItemGroup
}

type orSelection struct {
	expression string
	single bool
}

type orderSelection struct {
	asc bool
	field string
}




//Struct to hold out Select
type SelectSQL struct {
	from AliasItem
	columns AliasItem
	join []joinItem
	between []betweenSelection
	equalTo AliasItem
	notEqualTo AliasItem
	lessThan AliasItem
	lessThanEqualTo AliasItem
	greaterThan AliasItem
	greaterThanEqualTo AliasItem
	in []inSelection
	notIn []inSelection
	isNull ItemGroup
	isNotNull ItemGroup
	isLike []likeSelection
	isNotLike []likeSelection
	orGroup []ItemGroup
	groupBy ItemGroup
	orderBy []orderSelection
	having ItemGroup
}
