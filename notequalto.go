package buildquery

func (sb *SelectBuilder) NotEqualTo(notEqualTo AliasItem) *SelectBuilder {

	currentCols := sb.SelectSQL.notEqualTo
	//multiple calls to NotEqualTo will not over-write the column set
	combinedCols := combineSets(currentCols, notEqualTo)
	sb.SelectSQL.notEqualTo = combinedCols

	return sb
}

