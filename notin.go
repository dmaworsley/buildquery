package buildquery

func (sb *SelectBuilder) NotIn(field string, values ItemGroup) *SelectBuilder {

	notInStruct := inSelection{}
	notInStruct.field = field
	notInStruct.values = values

	sb.SelectSQL.notIn = append(sb.SelectSQL.notIn, notInStruct)

	return sb
}
