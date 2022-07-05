package controllers

import (
	"alarm/common"
	"alarm/dto"
	"alarm/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type AlarmNoticeMonitorSqlController struct {
	beego.Controller
}

// @router /alarm/monitor/index [get]
func (c *AlarmNoticeMonitorSqlController) AlarmMonitorIndex() {

	c.TplName = "alarm/monitor/index.tpl"
}

// @router /alarm/monitor/list/sql [get]
func (c *AlarmNoticeMonitorSqlController) ListMonitorSqlList() {

	sqlList := models.GetAlarmNoticeMonitorSqlList(nil)

	sqlTaskList := make([]dto.AlarmNoticeMonitorSql, 0)

	for _, m := range sqlList {

		jobId := common.JobPrefix + strconv.FormatInt(m.Id, 10)
		task := dto.AlarmNoticeMonitorSql{JobIsExits: common.AlarmCronTab.IsExists(jobId), TaskId: jobId}
		common.StructAssign(&task, &m)
		sqlTaskList = append(sqlTaskList, task)
	}

	result := map[string]interface{}{"code": "0"}
	result["data"] = sqlTaskList
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /alarm/monitor/add [post]
func (c *AlarmNoticeMonitorSqlController) AddMonitorTask() {
	var monitor models.AlarmNoticeMonitorSql
	data := c.Ctx.Input.RequestBody

	//json数据封装到user对象中
	err := json.Unmarshal(data, &monitor)
	if err != nil {
		logs.Info(monitor)
	}

	dto := models.AddMonitorSql(monitor)

	result := map[string]interface{}{"code": "200"}
	result["data"] = dto
	c.Data["json"] = result
	c.ServeJSON()
}
