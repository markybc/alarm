package dto

type AlarmNoticeMonitorSql struct {
	Id         int64
	Cron       string
	Sql        string
	AliasName  string
	Status     int
	DbId       int64
	Note       string
	Sn         string
	ShowSql    int
	JobIsExits bool
	TaskId     string
}
