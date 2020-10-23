package buildquery

func (sb *SelectBuilder) In(field string, values ItemGroup) *SelectBuilder {

	inStruct := inSelection{}
	inStruct.field = field
	inStruct.values = values

	sb.SelectSQL.in = append(sb.SelectSQL.in, inStruct)

	return sb
}
