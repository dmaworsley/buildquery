package buildquery

func (sb *SelectBuilder) IsNotLike(field string, value string) *SelectBuilder {

	likeStruct := likeSelection{}
	likeStruct.field = field
	likeStruct.value = value

	sb.SelectSQL.isNotLike = append(sb.SelectSQL.isNotLike, likeStruct)

	return sb
}
