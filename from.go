package buildquery

func (sb *SelectBuilder) From(tableDefinition AliasItem) *SelectBuilder {

	table := enforceFirstElement(tableDefinition)

	sb.SelectSQL.from = table
	return sb
}
