package buildquery

func (sb *SelectBuilder) IsNull(nulls ItemGroup) *SelectBuilder {

	currentIsNull := sb.SelectSQL.isNull
	combinedNull := combineGroups(currentIsNull, nulls)
	sb.SelectSQL.isNull = combinedNull

	return sb
}
