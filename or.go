package buildquery

func (sb *SelectBuilder) Or(orGroup ItemGroup) *SelectBuilder {

	sb.SelectSQL.orGroup = append(sb.SelectSQL.orGroup, orGroup)

	return sb
}
