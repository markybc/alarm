package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeConnectionController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeConnectionController"],
		beego.ControllerComments{
			Method:           "ListConnectionTask",
			Router:           "/alarm/list/connection",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeListenerController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeListenerController"],
		beego.ControllerComments{
			Method:           "ListenerList",
			Router:           "/alarm/list/listener",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeListenerController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeListenerController"],
		beego.ControllerComments{
			Method:           "AddListener",
			Router:           "/alarm/listener/addListener",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeListenerController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeListenerController"],
		beego.ControllerComments{
			Method:           "AddTopicListener",
			Router:           "/alarm/listener/addTopicListener",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"],
		beego.ControllerComments{
			Method:           "ListMonitorTask",
			Router:           "/alarm/list/task",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"],
		beego.ControllerComments{
			Method:           "AddMonitorTask",
			Router:           "/alarm/monitor/add",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"],
		beego.ControllerComments{
			Method:           "AlarmMonitorIndex",
			Router:           "/alarm/monitor/index",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"],
		beego.ControllerComments{
			Method:           "ListMonitorSqlList",
			Router:           "/alarm/monitor/list/sql",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"],
		beego.ControllerComments{
			Method:           "StartMonitorTask",
			Router:           "/alarm/monitor/start/task",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMonitorSqlController"],
		beego.ControllerComments{
			Method:           "StopMonitorTask",
			Router:           "/alarm/monitor/stop/task",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMsgController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeMsgController"],
		beego.ControllerComments{
			Method:           "SaveAlarmMsg",
			Router:           "/alarm/msg/save",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeTaskController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeTaskController"],
		beego.ControllerComments{
			Method:           "AlarmTaskIndex",
			Router:           "/alarm/task/index",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeTopicController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeTopicController"],
		beego.ControllerComments{
			Method:           "AlarmConnectionIndex",
			Router:           "/alarm/connection/index",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeTopicController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeTopicController"],
		beego.ControllerComments{
			Method:           "AlarmListenerIndex",
			Router:           "/alarm/listener/index",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeTopicController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeTopicController"],
		beego.ControllerComments{
			Method:           "AddTopic",
			Router:           "/alarm/topic/addTopic",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeTopicController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeTopicController"],
		beego.ControllerComments{
			Method:           "GetAllTopicList",
			Router:           "/alarm/topic/getList",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeTopicController"] = append(beego.GlobalControllerRouter["alarm/controllers:AlarmNoticeTopicController"],
		beego.ControllerComments{
			Method:           "AlarmTopicIndex",
			Router:           "/alarm/topic/index",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:MainController"] = append(beego.GlobalControllerRouter["alarm/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/index",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["alarm/controllers:TestController"] = append(beego.GlobalControllerRouter["alarm/controllers:TestController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/test/get",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
