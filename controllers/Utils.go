package controllers

import (
	"strconv"
	"fmt"
	"time"
)


//int 转换成string
func IntToString(intParams int) string{
	return string(intParams)
}

/**
string 转换成int
*/
func StringToInt(strParams string) int{
	intParams,error := strconv.Atoi(strParams)
	if error != nil{
		fmt.Println("字符串转换成整数失败")
	}
	return intParams
}

/**
时间戳
 */
func GetTime() string{
	return string(time.Now().Unix())
}
/**
获取当前时间
 */
func GetDateAndTime() string {

	return ""
}


