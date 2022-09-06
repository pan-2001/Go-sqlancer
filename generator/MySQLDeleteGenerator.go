package generator

import (
	"github.com/jmoiron/sqlx"
	"sqlancerProject/random"
	"strings"
)

func GenerateDelete(Db *sqlx.DB, tableName string) string {
	var sb strings.Builder //初始化
	sb.WriteString("DELETE")
	if random.GenerateRandomBool() {
		sb.WriteString(" LOW_PRIORITY")
	}
	if random.GenerateRandomBool() {
		sb.WriteString(" QUICK")
	}
	if random.GenerateRandomBool() {
		sb.WriteString(" IGNORE")
	}

	//生成表达式
	sb.WriteString(" FROM ")

	sb.WriteString(tableName)

	if random.GenerateRandomBool() {
		sb.WriteString(" WHERE ")
		//TODO 生成表达式
		sb.WriteString(GenerateColumn(Db, "table2"))
		//sb.WriteString("expression")
	}

	return sb.String()
}
