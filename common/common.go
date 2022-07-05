package common

import (
	"alarm/utils"
	"github.com/astaxie/beego/cache"
	"reflect"
	"time"
)

const (
	JobPrefix = "Job_"
)

var CacheMemory cache.Cache

var AlarmCronTab *utils.Crontab

func DelJobById(jobId string) {
	if AlarmCronTab.IsExists(jobId) {
		AlarmCronTab.DelByID(jobId)
	}
}

var CstSh, _ = time.LoadLocation("Asia/Shanghai") //上海

//binding type interface 要修改的结构体
//value type interace 有数据的结构体
func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}
