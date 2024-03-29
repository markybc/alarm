package routers

import (
	"alarm/controllers"
	"github.com/astaxie/beego"
)

func init() {
	alarm := beego.NewNamespace("alarm-api",
		beego.NSInclude(&controllers.MainController{}),
		beego.NSInclude(&controllers.TestController{}),
		beego.NSInclude(&controllers.AlarmNoticeMsgController{}),
		beego.NSInclude(&controllers.AlarmNoticeTaskController{}),
		beego.NSInclude(&controllers.AlarmNoticeMonitorSqlController{}),
		beego.NSInclude(&controllers.AlarmNoticeTopicController{}),
		beego.NSInclude(&controllers.AlarmNoticeListenerController{}),
		beego.NSInclude(&controllers.AlarmNoticeConnectionController{}),
	)
	beego.AddNamespace(alarm)
}
