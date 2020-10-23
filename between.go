package buildquery


func (sb *SelectBuilder) Between(field string, min string, max string) *SelectBuilder {

	betweenStruct := betweenSelection{}
	betweenStruct.field = field
	betweenStruct.min = min
	betweenStruct.max = max

	sb.SelectSQL.between = append(sb.SelectSQL.between, betweenStruct)

	return sb
}
