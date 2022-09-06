package mysql

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// 获得所有表名
func GetTableNames(Db *sqlx.DB) []string {
	rows, err := Db.Query("SHOW TABLES;")
	var tables []string
	var table string
	if err != nil {
		fmt.Printf("query failed, error:[%v]", err.Error())
		return tables
	}

	for rows.Next() {
		//定义变量接收查询数据
		err := rows.Scan(&table)
		if err != nil {
			fmt.Printf("get data failed, error:[%v]\n", err.Error())
		}
		//fmt.Println(table)
		tables = append(tables, table)
	}
	return tables
}

type Column struct {
	Field   string //列名
	Type    string //类型
	Null    string //是否为空
	Key     string //键
	Default string //
	Extra   string //
}

// 输入表名获得此表所有列的属性
func GetColumns(Db *sqlx.DB, tableName string) []Column {
	//fmt.Println("table name:", tableName)
	sqlInput := "DESCRIBE " + tableName
	rows, err := Db.Query(sqlInput)
	var Field, Type, Null, Key, Extra, DefaultOrNull string
	var Default sql.NullString
	var res []Column
	if err != nil {
		fmt.Printf("query failed, error:[%v]", err.Error())
		return res
	}

	for rows.Next() {
		//定义变量接收查询数据
		err := rows.Scan(&Field, &Type, &Null, &Key, &Default, &Extra)
		if err != nil {
			fmt.Printf("get data failed, error:[%v]\n", err.Error())
		}
		//fmt.Println(Field, Type, Null, Key, Default, Extra)
		if Default.Valid {
			DefaultOrNull = Default.String
		} else {
			DefaultOrNull = ""
		}
		t := Column{Field: Field, Type: Type, Null: Null, Key: Key, Default: DefaultOrNull, Extra: Extra}
		//fmt.Println("field:", t.Field, "\ntype:", t.Type, "\nnull:", t.Null, "\nkey:", t.Key, "\ndefault:", t.Default, "\nextra:", t.Extra)
		//fmt.Println(t.Field, ",", t.Type, ",", t.Null, ",", t.Key, ",", t.Default, ",", t.Extra)

		res = append(res, t)
	}

	return res
}

// 输入表名获得此表所有列名
func GetColumnNames(Db *sqlx.DB, tableName string) []string {
	columns := GetColumns(Db, tableName)
	var names []string
	for i := 0; i < len(columns); i++ {
		names = append(names, columns[i].Field)
	}
	return names
}

// 输入表名获得此表所有列名和类型
func GetColumnNamesAndTypes(Db *sqlx.DB, tableName string) ([]string, []string) {
	columns := GetColumns(Db, tableName)
	var names, types []string
	for i := 0; i < len(columns); i++ {
		names = append(names, columns[i].Field)
		types = append(types, columns[i].Type)
	}

	return names, types
}

//判断表是否为空（列的数量是否为0）
func TableIsEmpty(Db *sqlx.DB, tableName string) bool {
	res := GetColumnNames(Db, tableName)
	if len(res) == 0 {
		return true
	} else {
		return false
	}
}

func GetValuesInColumn(Db *sqlx.DB, tableName string, columnName string) []string {
	str := "SELECT " + columnName + " FROM " + tableName
	//fmt.Println(str)
	rows, err := Db.Query(str)
	var values []string
	var value sql.NullString
	if err != nil {
		fmt.Printf("query values failed, error:[%v]", err.Error())
		return values
	}

	for rows.Next() {
		//定义变量接收查询数据
		err := rows.Scan(&value)
		if err != nil {
			fmt.Printf("get data failed, error:[%v]\n", err.Error())
		}
		//fmt.Println(table)
		if value.Valid {
			values = append(values, value.String)
		} else {
			values = append(values, "null")
		}

	}
	return values
}
