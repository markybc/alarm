package models

import "github.com/astaxie/beego/orm"

// mysql 数据库连接配置
type AlarmDbConnection struct {
	Id        int64
	AliasName string
	Username  string
	Pwd       string
	HostPort  string
	Name      string
	Status    int
}

func GetAlarmDbConnectionList() []AlarmDbConnection {

	var dbList []AlarmDbConnection

	o := orm.NewOrm()

	_, _ = o.QueryTable("alarm_db_connection").All(&dbList)

	return dbList
}
