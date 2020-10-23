package buildquery

func (sb *SelectBuilder) EqualTo(equalTo AliasItem) *SelectBuilder {

	currentCols := sb.SelectSQL.equalTo
	//multiple calls to EqualTo will not over-write the column set
	combinedCols := combineSets(currentCols, equalTo)
	sb.SelectSQL.equalTo = combinedCols

	return sb
}

