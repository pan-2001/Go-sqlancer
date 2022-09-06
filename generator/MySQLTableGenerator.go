package generator

import (
	"github.com/jmoiron/sqlx"
	"sqlancerProject/random"
	"sqlancerProject/static"
	"strconv"
	"strings"
)

func GenerateTable(Db *sqlx.DB, tableName string, tableIsEmpty bool) string {
	var sb strings.Builder //初始化
	sb.WriteString("CREATE")
	sb.WriteString(" TABLE")

	static.WithConfig().PlusTableID()
	if random.GenerateRandomBool() {
		sb.WriteString(" IF NOT EXISTS")
	}
	sb.WriteString(" ")
	sb.WriteString(tableName)
	sb.WriteString(strconv.Itoa(static.WithConfig().GetTableID()))
	if random.GenerateRandomBool() && !tableIsEmpty {
		sb.WriteString(" LIKE ")
		sb.WriteString(random.RandomlyGetTable(Db))
		return sb.String()
	} else {
		sb.WriteString("(")
		for i := 0; i < random.GenerateRandomIntRange(1, 4); i++ {
			if i != 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(appendColumn())

		}
		sb.WriteString(")")
		sb.WriteString(" ")
		//TODO appendTableOptions()
		//TODO appendPartitionOptions()

		return sb.String()
	}

}

func appendColumn() string {
	var sb strings.Builder //初始化
	sb.WriteString("column")
	sb.WriteString(strconv.Itoa(static.WithConfig().GetColumnID()))
	static.WithConfig().PlusColumnID()
	sb.WriteString(" ")
	str := random.RandomlySelectOptions("DECIMAL", "INT", "VARCHAR", "FLOAT", "DOUBLE")
	switch str {
	case "DECIMAL":
		sb.WriteString("DECIMAL")
		//sb.WriteString(optionallyAddPrecisionAndScale())
	case "INT":
		sb.WriteString(random.RandomlySelectOptions("TINYINT", "SMALLINT", "MEDIUMINT", "INT", "BIGINT"))
		if random.GenerateRandomBool() {
			sb.WriteString("(")
			sb.WriteString(strconv.Itoa(random.GenerateRandomIntRange(0, 256)))
			sb.WriteString(")")
		}

	case "VARCHAR":
		sb.WriteString(random.RandomlySelectOptions("VARCHAR(500)", "TINYTEXT", "TEXT", "MEDIUMTEXT", "LONGTEXT"))

	case "FLOAT":
		sb.WriteString("FLOAT")
	//sb.WriteString(optionallyAddPrecisionAndScale())
	case "DOUBLE":
		sb.WriteString(random.RandomlySelectOptions("DOUBLE", "FLOAT"))
		//sb.WriteString(optionallyAddPrecisionAndScale())

	}
	if str == "INT" || str == "DOUBLE" || str == "FLOAT" || str == "DECIMAL" {
		if random.GenerateRandomBool() && str == "INT" {
			sb.WriteString(" UNSIGNED")
		}
		usesPQS := false
		if random.GenerateRandomBool() && usesPQS {
			sb.WriteString(" ZEROFILL")
		}
	}
	return sb.String()
}

func optionallyAddPrecisionAndScale() string {
	var sb strings.Builder //初始化
	if random.GenerateRandomBool() {
		sb.WriteString("(")
		n := random.GenerateRandomIntRange(1, 66)
		sb.WriteString(strconv.Itoa(n))
		sb.WriteString(", ")
		nCandidate := random.GenerateRandomIntRange(1, 31)
		if n < nCandidate {
			sb.WriteString(strconv.Itoa(n))
		} else {
			sb.WriteString(strconv.Itoa(nCandidate))
		}
		sb.WriteString(")")
	}
	return sb.String()
}
