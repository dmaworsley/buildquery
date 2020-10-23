package buildquery

import (
	"strings"
)


func (sb *SelectBuilder) BuildSelect() (query string) {

	//start building the query
	selectStatement := "SELECT "

	//step one scan all selectable columns from the from struct and join elements and merge into one
	currentSet := sb.SelectSQL.columns
	if isItemSet(sb.SelectSQL.join) {

		for _, val := range sb.SelectSQL.join {
			currentSet = combineSets(currentSet, val.columns)
		}
	}
	//turn merged columns into sql string
	sqlCols := aliasItemToSqlColumnString(currentSet)

	selectStatement += sqlCols

	//FROM to sql string
	fromSql := aliasItemToSqlColumnString(sb.SelectSQL.from)
	selectStatement += " FROM " + fromSql + " "

	//JOIN portion of the string
	if isItemSet(sb.SelectSQL.join) {

		var joinSql string
		for _, val := range sb.SelectSQL.join {
			joinSql +=  strings.ToUpper(val.joinType) + " JOIN " + aliasItemToSqlColumnString(val.joinTable) + " ON " + val.onClause + " "
		}
		selectStatement += joinSql
	}

	incWhere := false
	var builtUpWhere string

	if isItemSet(sb.SelectSQL.between) {

		var betweenSql string
		if !incWhere {
			incWhere = true
		} else {
			betweenSql += " AND "
		}

		for _, val := range sb.SelectSQL.between {
			betweenSql += val.field + " BETWEEN " + "'" + val.min + "'" + " AND " + "'" + val.max + "' "
		}
		builtUpWhere += betweenSql
	}

	if isItemSet(sb.SelectSQL.equalTo) {
		var equalToSql string
		if !incWhere {
			incWhere = true
		} else {
			equalToSql += " AND "
		}
		equalToSql += aliasItemToComparisonString(sb.SelectSQL.equalTo, "=")
		builtUpWhere += equalToSql
	}

	if isItemSet(sb.SelectSQL.notEqualTo) {
		var notEqualToSql string
		if !incWhere {
			incWhere = true
		} else {
			notEqualToSql += " AND "
		}
		notEqualToSql += aliasItemToComparisonString(sb.SelectSQL.notEqualTo, "=")
		builtUpWhere += notEqualToSql
	}

	if isItemSet(sb.SelectSQL.lessThan) {
		var lessThanSql string
		if !incWhere {
			incWhere = true
		} else {
			lessThanSql += " AND "
		}
		lessThanSql += aliasItemToComparisonString(sb.SelectSQL.lessThan, "<")
		builtUpWhere += lessThanSql
	}

	if isItemSet(sb.SelectSQL.lessThanEqualTo) {
		var lessThanEqualToSql string
		if !incWhere {
			incWhere = true
		} else {
			lessThanEqualToSql += " AND "
		}
		lessThanEqualToSql += aliasItemToComparisonString(sb.SelectSQL.lessThanEqualTo, "<=")
		builtUpWhere += lessThanEqualToSql
	}

	if isItemSet(sb.SelectSQL.greaterThan) {
		var greaterThan string
		if !incWhere {
			incWhere = true
		} else {
			greaterThan += " AND "
		}
		greaterThan += aliasItemToComparisonString(sb.SelectSQL.greaterThan, ">")
		builtUpWhere += greaterThan
	}

	if isItemSet(sb.SelectSQL.greaterThanEqualTo) {
		var greaterThanEqualTo string
		if !incWhere {
			incWhere = true
		} else {
			greaterThanEqualTo += " AND "
		}
		greaterThanEqualTo += aliasItemToComparisonString(sb.SelectSQL.greaterThanEqualTo, ">=")
		builtUpWhere += greaterThanEqualTo
	}

	if isItemSet(sb.SelectSQL.in) {
		var inSql string
		if !incWhere {
			incWhere = true
		} else {
			inSql += " AND "
		}
		inSql += inSelectionGroupToSqlString(sb.SelectSQL.in, true)
		builtUpWhere += inSql
	}

	if isItemSet(sb.SelectSQL.notIn) {
		var notInSql string
		if !incWhere {
			incWhere = true
		} else {
			notInSql += " AND "
		}
		notInSql += inSelectionGroupToSqlString(sb.SelectSQL.notIn, false)
		builtUpWhere += notInSql
	}

	if isItemSet(sb.SelectSQL.isNull) {
		var isNullSql string
		if !incWhere {
			incWhere = true
		} else {
			isNullSql += " AND "
		}
		isNullSql += itemGroupToNull(sb.SelectSQL.isNull, true)
		builtUpWhere += isNullSql
	}

	if isItemSet(sb.SelectSQL.isNotNull) {
		var isNotNullSql string
		if !incWhere {
			incWhere = true
		} else {
			isNotNullSql += " AND "
		}
		isNotNullSql += itemGroupToNull(sb.SelectSQL.isNotNull, false)
		builtUpWhere += isNotNullSql
	}

	if isItemSet(sb.SelectSQL.orGroup) {
		var orSql string
		//only needed in an or case it's believed
		firstElm := false
		if !incWhere {
			incWhere = true
			firstElm = true
		} else {
			orSql += " "
		}
		orSql += itemGroupSliceToOrString(sb.SelectSQL.orGroup, firstElm)
		builtUpWhere += orSql
	}

	if isItemSet(sb.SelectSQL.groupBy) {

		var groupSql string
		groupSql += " "
		groupSql += itemGroupToGroupBy(sb.SelectSQL.groupBy)
		builtUpWhere += groupSql
	}

	if isItemSet(sb.SelectSQL.orderBy) {

		var orderSql string
		orderSql += " "
		orderSql += orderSelectionToSql(sb.SelectSQL.orderBy)
		builtUpWhere += orderSql
	}

	if isItemSet(sb.SelectSQL.having) {

		var havingSql string
		havingSql += " "
		havingSql += itemGroupToHaving(sb.SelectSQL.having)
		builtUpWhere += havingSql
	}

	if incWhere {
		selectStatement += "WHERE " + builtUpWhere
	} else {
		selectStatement += builtUpWhere
	}
	query = selectStatement
	return
}