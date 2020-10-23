package buildquery

func (sb *SelectBuilder) GroupBy(groupBy ItemGroup) *SelectBuilder {

	currentGroups := sb.SelectSQL.groupBy
	combinedNull := combineGroups(currentGroups, groupBy)
	sb.SelectSQL.groupBy = combinedNull

	return sb
}