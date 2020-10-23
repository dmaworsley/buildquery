package buildquery

func (sb *SelectBuilder) LessThanEqualTo(lessThanEqTo AliasItem) *SelectBuilder {

	currentCols := sb.SelectSQL.lessThanEqualTo
	//multiple calls to EqualTo will not over-write the column set
	combinedCols := combineSets(currentCols, lessThanEqTo)
	sb.SelectSQL.lessThanEqualTo = combinedCols
	return sb
}