package myutils

import (
	//"fmt"
	"reflect"
	"strings"
)

func MakeHttpGetParamStr(obj interface{}) string {
	tot := reflect.TypeOf(obj)
	vot := reflect.ValueOf(obj)
	if tot.Kind() == reflect.Ptr {
	    tot = tot.Elem()
	    vot = vot.Elem()
	}
	arrParam := make([]string, vot.NumField())
	for i := 0; i < vot.NumField(); i++ {
		arrParam[i] = tot.Field(i).Name + "=" + vot.Field(i).String()
	}
	return strings.Join(arrParam, "&")
}
