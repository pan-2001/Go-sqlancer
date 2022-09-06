package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sqlancerProject/generator"
	"sqlancerProject/mysql"
	"sqlancerProject/oracle"
	"sqlancerProject/random"
	"sqlancerProject/static"
	"strings"
	"time"
)

func main() {
	Init()

	// 当前本地时间

	Time := time.Now().UnixNano()
	fmt.Println("Start!")
	var Db *sqlx.DB = mysql.ConnectMysql()
	_ = Db
	for i := 0; i < 1000000000000000; i++ {
		if !Start(Db) {
			fmt.Println("find a bug!")
			break
		}
		if i%1000 == 0 {
			Time1 := time.Now().UnixNano()
			//fmt.Println(Time1, Time)
			fmt.Println("num:", i, 1e12/(Time1-Time), "queries/s")
			Time = Time1
		}
	}

	//mysql.ExecuteQuery(Db, "SELECT * FROM table2")
	//tableName := "table2"

	//n := generator.GenerateExpression(Db, tableName, 0)
	//n = "SELECT * FROM " + tableName + "WHERE " + n
	//fmt.Println(n)
	//mysql.ExecuteQuery(Db, n)

	//var a int
	//_, _ = fmt.Scan(&a)
	//if a == 1 {
	//	mysql.ExecuteQuery(Db, n)
	//}
	//if a == 2 {
	//	continue
	//}

	//for i := 0; i < 100; i++ {
	//	s := generator.GenerateInsert(Db, "table2")
	//	mysql.ExecuteInsert(Db, s)
	//	//s := generator.GenerateTable(Db, "test", false)
	//	fmt.Println(s)
	//
	//}

	//for i := var Db *sqlx.DB = mysql.ConnectMysql()
	//	_ = Db0; i < 10000; i++ {
	//	fmt.Print(i, " ")
	//	n := generator.GenerateInsert(Db, tableName)
	//	s := n
	//
	//	fmt.Println(s)
	//	mysql.ExecuteInsert(Db, n)
	//}

}

func Init() {
	random.InitRandom()
	static.Init()
}

func MyPrint(s string) {
	s = strings.Replace(s, "\n", "\\n", -1)
	s = strings.Replace(s, "\t", "\\t", -1)
	s = strings.Replace(s, "\r", "\\r", -1)
	fmt.Println(s)
}

func Start(Db *sqlx.DB) bool {
	tableName := "table1"
	//for i := 0; i < 1000000000000000; i++ {
	//
	//	//n := generator.GenerateSelect(Db, tableName)
	//	//MyPrint(n)
	//	//mysql.ExecuteQuery(Db, n)
	//
	//	if !oracle.TLPCheck(Db, tableName) {
	//		fmt.Println("num:", i, "\n")
	//		break
	//	}
	//	if i%1000 == 0 {
	//		fmt.Println("num:", i)
	//	}
	//}
	var str string
	//switch random.RandomlySelectOptions("SELECT") {
	switch random.RandomlySelectOptions("INSERT", "DELETE", "SELECT") {
	case "INSERT":
		//fmt.Println("INSERT")
		for i := 0; i < 5; i++ {
			str = generator.GenerateInsert(Db, tableName)
		}

		//MyPrint(str)
		mysql.ExecuteInsert(Db, str)
	case "DELETE":
		//fmt.Println("DELETE")
		str = generator.GenerateDelete(Db, tableName)
		//MyPrint(str)
		affectedRows := mysql.ExecuteDelete(Db, str)
		randomInt := random.GenerateRandomIntRange(0, affectedRows)
		//fmt.Println(randomInt)
		for i := 0; i < randomInt; i++ {
			str = generator.GenerateInsert(Db, tableName)
		}

	case "SELECT":
		//fmt.Println("SELECT")
		f := oracle.TLPCheck(Db, tableName)
		if !f {
			return false
		}
		//str = generator.GenerateSelect(Db, tableName)
		//MyPrint(str)
		//n, _ := mysql.ExecuteQuery(Db, str)
		//fmt.Println("result lines: ", n)
	}

	//fmt.Println()
	return true
}
