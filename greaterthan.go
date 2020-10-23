package buildquery

func (sb *SelectBuilder) GreaterThan(greaterThan AliasItem) *SelectBuilder {

	currentCols := sb.SelectSQL.greaterThan
	//multiple calls to EqualTo will not over-write the column set
	combinedCols := combineSets(currentCols, greaterThan)
	sb.SelectSQL.greaterThan = combinedCols
	return sb
}