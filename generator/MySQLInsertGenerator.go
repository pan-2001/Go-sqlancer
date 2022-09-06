package generator

import (
	"github.com/jmoiron/sqlx"
	"sqlancerProject/random"
	"strings"
)

func GenerateInsert(Db *sqlx.DB, tableName string) string {
	//fmt.Println("Generate Insert:")
	var sb strings.Builder //初始化

	//添加insert
	sb.WriteString("INSERT")
	if random.GenerateRandomBool() {
		sb.WriteString(" ")
		sb.WriteString(random.RandomlySelectOptions("LOW_PRIORITY", "DELAYED", "HIGH_PRIORITY"))
	}
	if random.GenerateRandomBool() {
		sb.WriteString(" IGNORE")
	}

	//添加into
	sb.WriteString(" INTO ")

	//sb.WriteString(random.RandomlyGetTable(Db))
	sb.WriteString(tableName)

	//TODO:添加列名
	sb.WriteString(" (")
	columns, types := random.RandomlyGetNotEmptyColumnsAndTypes(Db, tableName)
	_ = types
	for i := 0; i < len(columns); i++ {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(columns[i])
	}
	sb.WriteString(") ")
	sb.WriteString("VALUES")
	sb.WriteString(" (")
	for i := 0; i < len(columns); i++ {
		if i != 0 {
			sb.WriteString(", ")
		}

		sb.WriteString(GenerateConstant())
	}
	sb.WriteString(")")
	return sb.String()
}
