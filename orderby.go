package buildquery

func (sb *SelectBuilder) OrderBy(asc bool, field string) *SelectBuilder {

	orderByStruct := orderSelection{}
	orderByStruct.asc = asc
	orderByStruct.field = field

	sb.SelectSQL.orderBy = append(sb.SelectSQL.orderBy, orderByStruct)

	return sb
}