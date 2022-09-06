package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

func updateRecord(Db *sqlx.DB) {
	//更新uid=1的username
	result, err := Db.Exec("update userinfo set username = 'anson' where uid = 1")
	if err != nil {
		fmt.Printf("update failed, error:[%v]", err.Error())
		return
	}
	num, _ := result.RowsAffected()
	fmt.Printf("update success, affected rows:[%d]\n", num)
}

//执行DELETE操作
func ExecuteDelete(Db *sqlx.DB, input string) int {
	//删除uid=2的数据
	result, err := Db.Exec(input)
	if err != nil {
		fmt.Printf("delete failed, error:[%v]", err.Error())
		return 0
	}
	num, _ := result.RowsAffected()
	//fmt.Printf("delete success, affected rows:[%d]\n", num)
	return int(num)
}

// 查询,
// input String
// output
func ExecuteQueryReturnMap(Db *sqlx.DB, sql string) (int, []map[string]string) {
	//sql = fmt.Sprintf("SELECT * FROM table2")
	query, err := Db.Query(sql)
	if err != nil {
		fmt.Println("查询数据库失败", err.Error())
		return 0, nil
	}
	// 字段名
	cols, _ := query.Columns()
	//values是每个列的值，这里获取到byte里
	values := make([][]byte, len(cols))
	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans := make([]interface{}, len(cols))
	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}
	//最后得到的map
	var results []map[string]string
	i := 0
	for query.Next() { //循环，让游标往下推
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			return 0, nil
		}

		row := make(map[string]string) //每行数据

		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			if key != "_row" {
				row[key] = string(v)
			}
		}
		results = append(results, row)
		i++
	}
	//fmt.Println(results)
	//fmt.Println(len(results))
	return len(results), results
}

// 查询,
// input String
// output
func ExecuteQuery(Db *sqlx.DB, sql string) (int, [][]string) {
	//sql = fmt.Sprintf("SELECT * FROM table2")
	query, err := Db.Query(sql)
	if err != nil {
		fmt.Println("查询数据库失败", err.Error())
		return 0, nil
	}
	// 字段名
	cols, _ := query.Columns()
	//values是每个列的值，这里获取到byte里
	values := make([][]byte, len(cols))
	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans := make([]interface{}, len(cols))
	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}
	//最后得到的map
	var results [][]string
	i := 0
	for query.Next() { //循环，让游标往下推
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			return 0, nil
		}

		//row := make(map[string]string) //每行数据
		var row []string

		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			if key != "_row" {
				var str string
				str = string(v)
				str = strings.Replace(str, "\n", "\\n", -1)
				str = strings.Replace(str, "\t", "\\t", -1)
				str = strings.Replace(str, "\r", "\\r", -1)

				if str == "" {
					str = "NULL"
				}
				row = append(row, str)
			}
		}
		results = append(results, row)
		i++
	}
	//fmt.Println(results)
	//fmt.Println(len(results))
	return len(results), results
}

//执行Insert 操作
func ExecuteInsert(Db *sqlx.DB, input string) {
	result, err := Db.Exec(input)
	if err != nil {
		fmt.Printf("data insert failed, error:[%v]", err.Error())
		return
	}
	id, _ := result.LastInsertId()
	//fmt.Printf("insert success, last id:[%d]\n", id)
	_ = id
}

//执行Create Table 操作
func ExecuteCreateTable(Db *sqlx.DB, input string) {
	result, err := Db.Exec(input)
	if err != nil {
		fmt.Printf("table create failed, error:[%v]", err.Error())
		return
	}
	id, _ := result.LastInsertId()
	fmt.Printf("create table success, last id:[%d]\n", id)
}
