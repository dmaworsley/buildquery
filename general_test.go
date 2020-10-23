package buildquery

import (
	"fmt"
	"testing"
	_"fmt"
)

func TestGeneral(t *testing.T) {

	query := Select()

	from := make(AliasItem)
	from["a"] = "affiliates"

	join_table := make(AliasItem)
	join_table["ad"] = "affiliate_details"
	onClause := "ad.user_id = a.user_id and ad.`key` = 'email'"
	columns_join:= make(AliasItem)
	columns_join["email"] = "`ad`." +  EscapeField("value")


	columns := make(AliasItem)
	columns["user_id"] = "`a`." +  EscapeField("user_id")
	columns["username"] = EscapeField("username")
	columns["join_date"] = EscapeField("join_date")
	columns["status"] = EscapeField("status")


	joinType := "inner"

	join := joinItem{}
	join.joinTable = join_table
	join.onClause = onClause
	join.columns = columns
	join.joinType = joinType

	//In .In(field, values)
	//field := "user_id"
	//values := ItemGroup{"1", "2", "3", "4", "5"}

	//Between .Between(field, min, max)
	//field := "a.user_id"
	//min := "1"
	//max := "100"

	//equalTo .EqualTo(equalTo)
	//equalTo := make(AliasItem)
	//equalTo["a.user_id"] = "1"
	//equalTo["a.username"] = "gofiliate"

	//lessThan .LessThan(lessThan)
	//lessThan := make(AliasItem)
	//lessThan["a.user_id"] = "1000"

	//In .In(field, values) || .NotIn(field, values)
	//field := "user_id"
	//values := ItemGroup{"1", "2", "3", "4", "5"}

	//IsNull .IsNull(fields) || .IsNotNull(fields)
	//fields := ItemGroup{"user_id", "join_date"}

	//Or .Or(orItem)
	//orItem := ItemGroup{"user_id < 9000"}
	//orItem_1 := ItemGroup{"user_id < 9000", "user_id > 20000"}

	//Group .GroupBy(group)
	group := ItemGroup{"1", "2"}

	//order by .OrderBy(asc, field)
	asc := false
	field := "user_id"

	havingGroup := ItemGroup{"COUNT(1) > 1"}

	query.From(from).Columns(columns).Join(join_table, onClause, columns_join, joinType).GroupBy(group).OrderBy(asc, field).Having(havingGroup)
	sql := query.BuildSelect()

	fmt.Printf("Query is: %v\n", sql)


}
