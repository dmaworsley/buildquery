package buildquery

func (sb *SelectBuilder) LessThan(lessThan AliasItem) *SelectBuilder {

	currentCols := sb.SelectSQL.lessThan
	//multiple calls to EqualTo will not over-write the column set
	combinedCols := combineSets(currentCols, lessThan)
	sb.SelectSQL.lessThan = combinedCols
	return sb
}