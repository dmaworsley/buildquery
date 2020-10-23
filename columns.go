package buildquery

func (sb *SelectBuilder) Columns(columns AliasItem) *SelectBuilder {

	currentCols := sb.SelectSQL.columns
	//multiple calls to Columns will not over-write the column set
	combinedCols := combineSets(currentCols, columns)
	sb.SelectSQL.columns = combinedCols

	return sb
}
