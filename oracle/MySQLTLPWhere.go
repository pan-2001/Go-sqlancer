package oracle

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
	"sort"
	"sqlancerProject/generator"
	"sqlancerProject/mysql"
	"sqlancerProject/random"
	"strings"
)

func TLPCheck(Db *sqlx.DB, tableName string) bool {

	//原始查询
	var originQueryString, randomColumn string
	flag := random.GenerateRandomBool()
	if flag {
		flag = true
		randomColumn = random.RandomlyGetOneColumn(Db, tableName)
		originQueryString = "SELECT * FROM " + tableName
	} else {
		flag = false
		randomColumn = random.RandomlyGetOneColumn(Db, tableName)
		s := generator.GenerateExpression(Db, tableName, 0)
		originQueryString = "SELECT * FROM " + tableName + " WHERE " + s
	}
	_ = randomColumn
	//originResult := 0
	//fmt.Println(originQueryString, originResult)
	//mysql.ExecuteQuery(Db, originQueryString)

	subString := generator.GenerateExpression(Db, tableName, 0)
	var firstQueryString, secondQueryString, thirdQueryString string
	if flag {
		//第一种情况: is
		firstQueryString = originQueryString + " WHERE " + subString
		//第二种情况: not
		secondQueryString = originQueryString + " WHERE NOT(" + subString + ")"
		//第三种情况: null
		thirdQueryString = originQueryString + " WHERE (" + subString + ")is NULL"
	} else {
		//第一种情况: is
		firstQueryString = originQueryString + " AND " + subString
		//第二种情况: not
		secondQueryString = originQueryString + " AND NOT(" + subString + ")"
		//第三种情况: null
		thirdQueryString = originQueryString + " AND (" + subString + ")is NULL"
	}

	//fmt.Println(firstQueryString, "\n")
	//fmt.Println(secondQueryString, "\n")
	//fmt.Println(thirdQueryString, "\n")

	//将三种情况结合起来
	unionString := UnionAll(firstQueryString, secondQueryString, thirdQueryString)
	//fmt.Println(unionString)
	lenOrigin, resultOrigin := mysql.ExecuteQuery(Db, originQueryString)
	lenUnion, resultUnion := mysql.ExecuteQuery(Db, unionString)

	//比较两个结果是否相等
	if lenOrigin != lenUnion {
		fmt.Println(FormatPrint(originQueryString), "\n\n", FormatPrint(unionString), "\n\n")
		fmt.Println(lenOrigin, lenUnion)
		return false
	}

	if !Compare(resultOrigin, resultUnion) {
		fmt.Println(resultOrigin)
		fmt.Println(resultUnion)

		fmt.Println("\norigin: ", FormatPrint(originQueryString), "\n\nunion: ", FormatPrint(unionString), "\n\n")
		return false
	}

	return true
}

func UnionAll(firstQueryString string, secondQueryString string, thirdQueryString string) string {
	return firstQueryString + " UNION ALL " + secondQueryString + " UNION ALL " + thirdQueryString
}

func Compare(inputOne [][]string, inputTwo [][]string) bool {
	var temp1, temp2 []string
	for _, value := range inputOne {
		var t string
		for _, v := range value {
			t += v
		}
		temp1 = append(temp1, t)
	}
	for _, value := range inputTwo {
		var t string
		for _, v := range value {
			t += v
		}
		temp2 = append(temp2, t)
	}
	sort.Sort(sort.StringSlice(temp1))
	sort.Sort(sort.StringSlice(temp2))
	return reflect.DeepEqual(temp1, temp2)
	//return true
}
func FormatPrint(s string) string {
	s = strings.Replace(s, "\n", "\\n", -1)
	s = strings.Replace(s, "\t", "\\t", -1)
	s = strings.Replace(s, "\r", "\\r", -1)
	return s
}
