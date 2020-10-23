package buildquery

func (sb *SelectBuilder) IsNotNull(nulls ItemGroup) *SelectBuilder {

	currentIsNotNull := sb.SelectSQL.isNotNull
	combinedNull := combineGroups(currentIsNotNull, nulls)
	sb.SelectSQL.isNotNull = combinedNull

	return sb
}