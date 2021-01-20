package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

/* 字符串转int */
func StringToInt(data string) int {
	num, _ := strconv.Atoi(data)
	return num
}

/* 字符串转int64 */
func StringToInt64(data string) int64 {
	num, _ := strconv.ParseInt(data, 10, 64)
	return num
}

/* 整形64转字符串 */
func Int64ToString(data int64) string {
	return strconv.FormatInt(data, 10)
}

/* 整形转字符串 */
func IntToString(data int) string {
	return strconv.Itoa(data)
}

/* DSN数据库连接串 */
func DSN(user, password, host, port, dbName, parameters string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, port, dbName, parameters)
}

/* md5加密 */
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
