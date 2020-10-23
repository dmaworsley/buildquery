package buildquery

func (sb *SelectBuilder) IsLike(field string, value string) *SelectBuilder {

	likeStruct := likeSelection{}
	likeStruct.field = field
	likeStruct.value = value

	sb.SelectSQL.isLike = append(sb.SelectSQL.isLike, likeStruct)

	return sb
}
