package buildquery

import (
	"reflect"
)

func combineSets( set1 AliasItem, set2 AliasItem) (returnSet AliasItem) {

	workingSet := make(AliasItem)
	for key, val := range set1 {
		workingSet[key] = val
	}

	for key, val := range set2 {
		workingSet[key] = val
	}

	returnSet = workingSet
	return
}

func combineGroups(group1 ItemGroup, group2 ItemGroup) (returnGroup ItemGroup) {

	var workingGroup ItemGroup

	workingGroup = append(workingGroup, group1...)
	workingGroup = append(workingGroup, group2...)

	returnGroup = workingGroup
	return
}


func enforceFirstElement(set AliasItem) (returnSet AliasItem) {

	workingSet := make(AliasItem)
	cnt := 1

	for key, value := range set {

		if cnt == 1 {
			workingSet[key] = value
			cnt++
		}
	}
	returnSet = workingSet
	return
}

func EscapeField(field string) (escapedField string) {
	escapedField = "`" + field + "`"
	return
}

func aliasItemToSqlColumnString(aliasItems AliasItem) (sqlString string) {

	lenItem := len(aliasItems)
	i := 0

	for key, value := range aliasItems {

		sqlString +=  value + " as " + "`" + key + "`"

		if i != (lenItem-1) {
			sqlString += ", "
		}
		i++
	}
	return
}

func isItemSet(item interface{}) (isSet bool) {

	switch reflect.TypeOf(item).Kind() {
		case reflect.Slice:
			val := reflect.ValueOf(item)
			if val.Len() > 0 {
				isSet = true
			}
		case reflect.Map:
			val := reflect.ValueOf(item)
			if val.Len() > 0 {
				isSet = true
			}
	}
	return
}

func aliasItemToComparisonString(comparableItem AliasItem, operator string) (sqlString string) {

	lenItem := len(comparableItem)
	i :=0

	for key, value := range comparableItem {

		sqlString += key + " " + operator + " " + "'" + value + "'"

		if i != (lenItem-1) {
			sqlString += " AND "
		} else {
			sqlString += " "
		}
		i++
	}
	return
}

func implodeItemGroup(itemGroup ItemGroup, glue string) (gluedString string) {

	lenItem := len(itemGroup)
	i :=0

	for _, val := range itemGroup {

		gluedString += "'" + val + "'"

		if i != (lenItem-1) {
			gluedString += glue + " "
		}
		i++
	}
	return
}

func inSelectionGroupToSqlString(inSelection []inSelection, in bool) (sqlString string) {

	lenItem := len(inSelection)
	i := 0
	for _, val := range inSelection {
		group := implodeItemGroup(val.values, ",")

		sqlString += val.field
		if !in {
			sqlString += " NOT"
		}
		sqlString += " IN(" + group + ")"

		if i != (lenItem-1) {
			sqlString += " AND "
		} else {
			sqlString += " "
		}
		i++
	}
	return
}


func itemGroupToNull(itemGroup ItemGroup, isNull bool) (sqlString string) {

	lenItem := len(itemGroup)
	i :=0
	var opStr string

	if isNull {
		opStr = " IS NULL"
	} else {
		opStr = " IS NOT NULL"
	}

	for _, val := range itemGroup {

		sqlString += val + opStr

		if i != (lenItem-1) {
			sqlString += " AND "
		} else {
			sqlString += " "
		}
		i++
	}
	return
}

func itemGroupSliceToOrString(itemSlice []ItemGroup, isFirstElm bool) (sqlString string) {

	lenItem := len(itemSlice)
	i := 0

	for _, item := range itemSlice {

		switch len(item) {

			case 1:
				for _, val := range item {
					sqlString += "OR " + val
				}

			case 2:
				if isFirstElm {
					sqlString += "("
				} else {
					sqlString += "AND ("
				}

				cnt := 0
				for _, val := range item {
					sqlString += val
					if cnt == 0 {
						sqlString += " OR "
						cnt++
					}
				}
				sqlString += ")"

		}

		if i != (lenItem-1) {
			sqlString += " AND "
		} else {
			sqlString += " "
		}
		i++
	}
	return
}


func orderSelectionToSql(orderBy []orderSelection) (sqlString string) {

	lenItem := len(orderBy)
	i :=0

	sqlString += "ORDER BY "
	for _, val := range orderBy {
		sqlString += val.field
		if val.asc {
			sqlString +=  " ASC"
		} else {
			sqlString +=  " DESC"
		}

		if i != (lenItem-1) {
			sqlString += ", "
		} else {
			sqlString += " "
		}
		i++
	}

	return
}

func itemGroupToGroupBy(itemGroup ItemGroup) (sqlString string) {

	lenItem := len(itemGroup)
	i :=0

	sqlString += "GROUP BY "
	for _, val := range itemGroup {
		sqlString += val
		if i != (lenItem-1) {
			sqlString += ", "
		} else {
			sqlString += " "
		}
		i++
	}
	return
}

func itemGroupToHaving(itemGroup ItemGroup) (sqlString string) {
	lenItem := len(itemGroup)
	i :=0

	sqlString += "HAVING "

	for _, val := range itemGroup {
		sqlString += val
		if i != (lenItem-1) {
			sqlString += ", "
		} else {
			sqlString += " "
		}
		i++
	}
	return
}
