package generator

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sqlancerProject/mysql"
	"sqlancerProject/random"
	"sqlancerProject/static"
	"strconv"
)

func GenerateExpression(Db *sqlx.DB, tableName string, depth int) string {

	if depth >= static.WithConfig().GetExpressionDepth() {
		if random.GenerateRandomBool() && mysql.TableIsEmpty(Db, tableName) {
			return GenerateColumn(Db, tableName)
		} else {
			return GenerateConstant()
		}
	}

	n := []string{
		"COLUMN", "LITERAL", "UNARY_PREFIX_OPERATION", "UNARY_POSTFIX", "COMPUTABLE_FUNCTION", "BINARY_LOGICAL_OPERATOR",
		"BINARY_COMPARISON_OPERATION", "CAST", "IN_OPERATION", "BINARY_OPERATION", "EXISTS", "BETWEEN_OPERATOR",
	}
	n = []string{"COLUMN", "LITERAL", "BINARY_OPERATION", "EXISTS", "BINARY_LOGICAL_OPERATOR", "CAST",
		"BINARY_COMPARISON_OPERATION", "BETWEEN_OPERATOR", "UNARY_PREFIX_OPERATION", "UNARY_POSTFIX"}
	//n = []string{"UNARY_POSTFIX"}
	switch random.RandomlySelectOptionsByString(n) {
	case "COLUMN":
		//Done
		//fmt.Println("COLUMN")
		return "(" + GenerateColumn(Db, tableName) + ")"
	case "LITERAL":
		//Done
		//fmt.Println("LITERAL")
		return GenerateConstant()
	case "UNARY_PREFIX_OPERATION":
		//Done
		//fmt.Println("UNARY_PREFIX_OPERATION")
		return "(" + GetRandomUnaryPrefixOperator() + GenerateExpression(Db, tableName, depth+1) + ")"
	case "UNARY_POSTFIX":
		//Done
		//fmt.Println("UNARY_POSTFIX")
		return "(" + GenerateExpression(Db, tableName, depth+1) + " " + GetRandomUnaryPostfixOperation() + ")"
	case "COMPUTABLE_FUNCTION":

		fmt.Println("COMPUTABLE_FUNCTION")
	case "BINARY_LOGICAL_OPERATOR":
		//Done
		//fmt.Println("BINARY_LOGICAL_OPERATOR")
		return "(" + GenerateExpression(Db, tableName, depth+1) + GetRandomBinaryLogicalOperator() +
			GenerateExpression(Db, tableName, depth+1) + ")"
	case "BINARY_COMPARISON_OPERATION":
		//Done
		//fmt.Println("BINARY_COMPARISON_OPERATION")
		return "(" + GenerateExpression(Db, tableName, depth+1) + GetRandomBinaryComparisonOperation() +
			GenerateExpression(Db, tableName, depth+1) + ")"
	case "CAST":
		//Done
		//fmt.Println("CAST")
		return "(CAST(" + GenerateExpression(Db, tableName, depth+1) + " as" + GetRandomCastOperation() + "))"
	case "IN_OPERATION":
		fmt.Println("IN_OPERATION")
	case "BINARY_OPERATION":
		//Done ?
		//fmt.Println("BINARY_OPERATION")
		return "(" + GenerateExpression(Db, tableName, depth+1) + GetRandomBinaryOperator() +
			GenerateExpression(Db, tableName, depth+1) + ")"
	case "EXISTS":
		//Done
		//fmt.Println("EXISTS")
		return GetExists(tableName)
	case "BETWEEN_OPERATOR":
		//Done
		//fmt.Println("BETWEEN_OPERATOR")
		return "(" + GenerateExpression(Db, tableName, depth+1) + " BETWEEN " + GenerateExpression(Db, tableName, depth+1) +
			" AND " + GenerateExpression(Db, tableName, depth+1) + ")"
	}
	//fmt.Println(n)
	return "0"
}

func GenerateConstant() string {

	usesPQS := false
	var values []string
	if usesPQS {
		values = []string{"INT", "NULL", "STRING"}
	} else {
		values = []string{"INT", "NULL", "STRING", "DOUBLE"}
	}

	switch random.RandomlySelectOptionsByString(values) {
	case "INT":
		return strconv.Itoa(int(random.GenerateRandomInt()))
	case "DOUBLE":
		return strconv.FormatFloat(random.GenerateRandomDouble(), 'f', -1, 64)
	case "NULL":
		return "null"
	case "STRING":
		return "\"" + random.GenerateRandomNumStringRange(5) + "\""
	}
	return ""
}

//生成列条件变量
func GenerateColumn(Db *sqlx.DB, tableName string) string {
	columnName := random.RandomlyGetOneColumn(Db, tableName)
	rowVal := random.RandomlyGetValueInColumn(Db, tableName, columnName)
	//fmt.Println(rowVal)
	if rowVal == "null" {
		return columnName + " is NULL"
	} else {
		return columnName + " = " + rowVal
	}
}

func GetExists(tableName string) string {
	if random.GenerateRandomBool() {
		return "EXISTS (SELECT 1 FROM " + tableName + ")"
	} else {
		return "EXISTS (SELECT 1 FROM " + tableName + " WHERE FALSE)"
	}
}

func GetRandomBinaryOperator() string {
	str := []string{"AND", "OR", "XOR"}
	var op string
	switch random.RandomlySelectOptionsByString(str) {
	case "AND":
		op = "&"
	case "OR":
		op = "|"
	case "XOR":
		op = "^"
	}

	return op
}

func GetRandomBinaryLogicalOperator() string {
	str := []string{"AND", "OR", "XOR"}
	var op string
	switch random.RandomlySelectOptionsByString(str) {
	case "AND":
		op = " AND "
	case "OR":
		op = " OR "
	case "XOR":
		op = " XOR "
	}

	return op
}

func GetRandomBinaryComparisonOperation() string {
	str := []string{"EQUALS", "NOT_EQUALS", "LESS", "LESS_EQUALS", "GREATER", "GREATER_EQUALS"}
	var op string
	switch random.RandomlySelectOptionsByString(str) {
	case "EQUALS":
		op = " = "
	case "NOT_EQUALS":
		op = " != "
	case "LESS":
		op = " < "
	case "LESS_EQUALS":
		op = " <= "
	case "GREATER":
		op = " > "
	case "GREATER_EQUALS":
		op = " >= "
	}

	return op
}

func GetRandomCastOperation() string {
	str := []string{"SIGNED", "UNSIGNED"}
	var op string
	switch random.RandomlySelectOptionsByString(str) {
	case "SIGNED":
		op = " SIGNED "
	case "UNSIGNED":
		op = " UNSIGNED "
	}

	return op
}

func GetRandomUnaryPrefixOperator() string {
	str := []string{"PLUS", "MINUS", "NOT"}
	var op string
	switch random.RandomlySelectOptionsByString(str) {
	case "PLUS":
		op = "+"
	case "MINUS":
		op = "-"
	case "NOT":
		op = "!"
	}

	return op
}

func GetRandomUnaryPostfixOperation() string {
	str := []string{"IS_NULL", "IS_TRUE", "IS_FALSE"}
	var op string
	switch random.RandomlySelectOptionsByString(str) {
	case "IS_NULL":
		op = "IS NULL"
	case "IS_TRUE":
		op = "IS TRUE"
	case "IS_FALSE":
		op = "IS FALSE"
	}

	return op
}
