package buildquery

func (sb *SelectBuilder) Having(having ItemGroup) *SelectBuilder {

	currentHaving := sb.SelectSQL.having
	combinedHaving:= combineGroups(currentHaving, having)
	sb.SelectSQL.having = combinedHaving

	return sb
}