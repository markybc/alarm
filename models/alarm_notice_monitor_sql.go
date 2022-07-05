package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type AlarmNoticeMonitorSql struct {
	Id        int64
	Cron      string
	Sql       string
	AliasName string
	Status    int
	DbId      int64
	Note      string
	Sn        string
	ShowSql   int
}

func AddMonitorSql(m AlarmNoticeMonitorSql) AlarmNoticeMonitorSql {

	o := orm.NewOrm()
	var db AlarmDbConnection
	_ = o.QueryTable("alarm_db_connection").Filter("id", m.DbId).One(&db)
	m.AliasName = db.AliasName
	m.DbId = db.Id

	id, err := o.Insert(&m)
	if err != nil {
		logs.Error(err)
	}
	m.Id = id
	return m
}

func GetAlarmNoticeMonitorSqlList(status interface{}) []AlarmNoticeMonitorSql {
	var sqlList []AlarmNoticeMonitorSql
	o := orm.NewOrm()
	qs := o.QueryTable("alarm_notice_monitor_sql")
	if status != nil {
		qs.Filter("status", status).All(&sqlList)
	} else {
		qs.All(&sqlList)
	}
	return sqlList
}

func GetAlarmNoticeMonitorSqlById(id int64, o orm.Ormer) AlarmNoticeMonitorSql {
	var sql AlarmNoticeMonitorSql
	if o == nil {
		o = orm.NewOrm()
	}
	qs := o.QueryTable("alarm_notice_monitor_sql")
	qs.Filter("id", id).All(&sql)
	return sql
}
