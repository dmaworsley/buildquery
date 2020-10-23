package buildquery


func (sb *SelectBuilder) Join(joinTable AliasItem, onClause string, columns AliasItem, joinType string) *SelectBuilder {

	joinStruct := joinItem{}

	//only allowed 1 table in the AliasItem set
	table := enforceFirstElement(joinTable)
	joinStruct.joinTable = table
	joinStruct.onClause = onClause
	joinStruct.columns = columns
	joinStruct.joinType = joinType

	sb.SelectSQL.join = append(sb.SelectSQL.join, joinStruct)

	return sb
}
