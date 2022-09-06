package generator

import (
	"github.com/jmoiron/sqlx"
	"sqlancerProject/random"
	"strings"
)

func GenerateSelect(Db *sqlx.DB, tableName string) string {
	var n string
	if random.GenerateRandomBool() {
		n = "SELECT * FROM " + tableName + " WHERE " + GenerateExpression(Db, tableName, 0)
	} else {
		var sb strings.Builder //初始化
		columns, _ := random.RandomlyGetNotEmptyColumnsAndTypes(Db, tableName)
		flag := false
		for _, value := range columns {
			if flag {
				sb.WriteString(", ")
			}
			sb.WriteString(value)
			flag = true
		}
		n = "SELECT " + sb.String() + " FROM " + tableName + " WHERE " + GenerateExpression(Db, tableName, 0)
	}

	return n
}
