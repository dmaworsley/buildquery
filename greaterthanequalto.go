package buildquery

func (sb *SelectBuilder) GreaterThanEqualTo(greaterThanEqTo AliasItem) *SelectBuilder {

	currentCols := sb.SelectSQL.greaterThanEqualTo
	//multiple calls to EqualTo will not over-write the column set
	combinedCols := combineSets(currentCols, greaterThanEqTo)
	sb.SelectSQL.greaterThanEqualTo = combinedCols
	return sb
}