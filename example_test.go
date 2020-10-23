package buildquery

import (
	"fmt"
	"testing"
)

func TestDocExample(t *testing.T) {

	//create a new Select
	query := Select()

	//which table in the from clause
	from_table := make(AliasItem)
	//AliasItem = map[string]string
	from_table["u"] = "users"

	query.From(from_table)

	columns := make(AliasItem)
	columns["user_id"] = "`u`." +  EscapeField("user_id")
	columns["username"] = "username"
	columns["join_date"] = "join_date"
	columns["status"] = EscapeField("status")

	query.Columns(columns)

	//lets join something

	join_table := make(AliasItem)
	join_table["ud"] = "user_details"

	join_columns := make(AliasItem)
	join_columns["email"] = "`ud`." +  EscapeField("value")


	onClause := "ud.user_id = u.user_id and ud.`key` = 'email'"
	joinType := "INNER"

	query.Join(join_table, onClause, join_columns, joinType)

	//lets add a less Or Equal than clause to the query
	lessThanOrEqual := make(AliasItem)
	lessThanOrEqual["u.user_id"] = "5"

	query.LessThanEqualTo(lessThanOrEqual)

	//get our sql string
	sql := query.BuildSelect()

	//assuming you include fmt
	fmt.Printf("%v\n", sql)
}
