package random

import (
	"github.com/jmoiron/sqlx"
	"math/rand"
	"sqlancerProject/mysql"
	"strings"
	"time"
)

func InitRandom() {
	rand.Seed(time.Now().Unix()) //产生Seed
}

func GenerateRandomInt() int32 {
	n := rand.Int31()
	if rand.Intn(2) == 1 {
		n = -n
	}
	return n
}

// 根据范围生成随机整数 [min, max)
func GenerateRandomIntRange(min int, max int) int {
	if min >= max {
		return min
	}
	n := min + rand.Intn(max-min)
	return n
}

func GenerateRandomFloat() float32 {
	//fmt.Println("random float:")
	n := rand.Float32()
	return n
}

func GenerateRandomDouble() float64 {
	//fmt.Println("random double:")
	n := rand.Float64()
	return n
}

//生成一个随机字符，范围：可打印字符和\r \n \t
func GenerateRandomChar() string {
	//fmt.Println("random char:")
	var CHARS = []string{
		"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		":", ";", "<", "=", ">", "?", "@",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"[", "\\", "]", "^", "_", "`",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"{", "|", "}", "~",
		"\n", "\r", "\t"}
	length := len(CHARS)
	char := CHARS[rand.Intn(length)]
	return char
}

//生成一个随机字符，范围：数字
func GenerateRandomNumChar() string {
	//fmt.Println("random char:")
	var CHARS = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	length := len(CHARS)
	char := CHARS[rand.Intn(length)]
	return char
}

// 生成随机字符串
func GenerateRandomString() string {
	//fmt.Println("random string:")
	length := int(GenerateRandomIntRange(1, 256))
	//fmt.Println(length)
	var sb strings.Builder //初始化
	sb.WriteString("")
	for i := 0; i < length; i++ {
		sb.WriteString(GenerateRandomChar())
	}
	return sb.String()
}

// 生成指定长度的随机字符串
func GenerateRandomStringRange(input int) string {
	//fmt.Println("random string:")
	length := int(GenerateRandomIntRange(1, input+1))
	//fmt.Println(length)
	var sb strings.Builder //初始化
	sb.WriteString("")
	for i := 0; i < length; i++ {
		sb.WriteString(GenerateRandomChar())
	}
	return sb.String()
}

// 生成只有数字和\r \n \t结尾的随机字符串
func GenerateRandomNumString() string {
	//fmt.Println("random string:")
	length := int(GenerateRandomIntRange(1, 256))
	//fmt.Println(length)
	var sb strings.Builder //初始化
	sb.WriteString("")
	for i := 0; i < length; i++ {
		sb.WriteString(GenerateRandomNumChar())
	}
	//TODO 添加小数点
	for i := 0; i < GenerateRandomIntRange(0, 5); i++ {
		switch GenerateRandomIntRange(0, 3) {
		case 0:
			sb.WriteString("\n")
		case 1:
			sb.WriteString("\t")
		case 2:
			sb.WriteString("\r")
		}

	}
	return sb.String()
}

// 生成指定长度的只有数字和\r \n \t结尾的随机字符串
func GenerateRandomNumStringRange(input int) string {
	//fmt.Println("random string:")
	length := int(GenerateRandomIntRange(1, input))
	//fmt.Println(length)
	var sb strings.Builder //初始化
	sb.WriteString("")
	for i := 0; i < length; i++ {
		sb.WriteString(GenerateRandomNumChar())
	}
	for i := 0; i < GenerateRandomIntRange(0, 5); i++ {
		switch GenerateRandomIntRange(0, 3) {
		case 0:
			sb.WriteString("\n")
		case 1:
			sb.WriteString("\t")
		case 2:
			sb.WriteString("\r")
		}

	}
	return sb.String()
}

//随机生成一个bool值
func GenerateRandomBool() bool {
	//fmt.Println("random string:")
	n := rand.Intn(2)
	if n == 1 {
		return true
	} else {
		return false
	}
}

// 随机选择输入选项,
// input []string,
// output string,
func RandomlySelectOptions(args ...interface{}) string {
	length := len(args)
	return args[GenerateRandomIntRange(0, length)].(string)
}

//输入string数组，随机选择一个元素返回
func RandomlySelectOptionsByString(input []string) string {
	length := len(input)
	if length == 0 {
		return "null"
	}
	return input[GenerateRandomIntRange(0, length)]
}

//随机返回数据库一个表的名字
func RandomlyGetTable(Db *sqlx.DB) string {
	input := mysql.GetTableNames(Db)
	length := len(input)
	return input[GenerateRandomIntRange(0, length)]
}

//返回一个非空的列名子集
func RandomlyGetNotEmptyColumns(Db *sqlx.DB, tableName string) []string {
	columns := mysql.GetColumnNames(Db, tableName)
	//fmt.Println(columns)
	var notEmptyColumns []string
	for len(notEmptyColumns) == 0 {
		for i := 0; i < len(columns); i++ {
			if GenerateRandomBool() {
				notEmptyColumns = append(notEmptyColumns, columns[i])
			}
		}
	}
	return notEmptyColumns
}

//返回一个非空的列名和类型的子集
func RandomlyGetNotEmptyColumnsAndTypes(Db *sqlx.DB, tableName string) ([]string, []string) {
	columns, types := mysql.GetColumnNamesAndTypes(Db, tableName)
	//fmt.Println(columns)
	var notEmptyColumnNames, notEmptyColumnTypes []string
	for len(notEmptyColumnNames) == 0 {
		for i := 0; i < len(columns); i++ {
			if GenerateRandomBool() {
				notEmptyColumnNames = append(notEmptyColumnNames, columns[i])
				notEmptyColumnTypes = append(notEmptyColumnTypes, types[i])
			}
		}
	}
	return notEmptyColumnNames, notEmptyColumnTypes
}

//随机返回一个列名
func RandomlyGetOneColumn(Db *sqlx.DB, tableName string) string {
	columns := mysql.GetColumnNames(Db, tableName)
	return RandomlySelectOptionsByString(columns)
}

func RandomlyGetValueInColumn(Db *sqlx.DB, tableName string, columnName string) string {
	values := mysql.GetValuesInColumn(Db, tableName, columnName)
	return RandomlySelectOptionsByString(values)
}
