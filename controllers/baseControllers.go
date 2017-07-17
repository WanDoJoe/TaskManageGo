package controllers

import (
	"github.com/widuu/gojson"
	"fmt"
)

func GetParams(paramsJson string ,key string)string {

	params :=gojson.Json(paramsJson).Get("params").Get(key).Tostring()
	fmt.Println("params="+params)
	return params
}


func SetSysInfo(paramsJson string ){
	//gojson.Json(paramsJson).Get("sysinfo")
	//fmt.Print("SetSysInfo="+sysinfo.Tostring())
}
