package dto

type AlarmNoticeMonitorJobOnOff struct {
	Id int64
}

type AlarmNoticeMonitor struct {
	Sql       string
	Cron      string
	AliasName string
	TopicSn   string
}
