package controllers

import (
	"alarm/common"
	"alarm/dto"
	"alarm/job"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

type AlarmNoticeTaskController struct {
	beego.Controller
}

// @router /alarm/task/index [get]
func (c *AlarmNoticeTaskController) AlarmTaskIndex() {

	c.TplName = "alarm/task/index.tpl"
}

// @router /alarm/monitor/start/task [post]
func (c *AlarmNoticeMonitorSqlController) StartMonitorTask() {
	var task dto.AlarmNoticeMonitorJobOnOff
	data := c.Ctx.Input.RequestBody

	//json数据封装到user对象中
	err := json.Unmarshal(data, &task)
	if err != nil {
		logs.Info(task)
	}

	cacheKey := "StartMonitorTask_" + strconv.FormatInt(task.Id, 10)
	result := map[string]interface{}{"code": "200"}
	if !common.CacheMemory.IsExist(cacheKey) {
		common.CacheMemory.Put(cacheKey, cacheKey, 30*time.Second)

		jobId := common.JobPrefix + strconv.FormatInt(task.Id, 10)
		if !common.AlarmCronTab.IsExists(jobId) {
			job.AddAlarmMonitorJob(task.Id)
			logs.Info("start monitor task：", jobId)
		} else {
			logs.Info("monitor task：%v is already start", jobId)
		}
	}

	c.Data["json"] = result
	c.ServeJSON()
}

// @router /alarm/monitor/stop/task [post]
func (c *AlarmNoticeMonitorSqlController) StopMonitorTask() {
	var task dto.AlarmNoticeMonitorJobOnOff
	data := c.Ctx.Input.RequestBody

	//json数据封装到user对象中
	err := json.Unmarshal(data, &task)
	if err != nil {
		logs.Info(task)
	}

	jobId := common.JobPrefix + strconv.FormatInt(task.Id, 10)
	if common.AlarmCronTab.IsExists(jobId) {
		logs.Info("monitor task：%v is already start", jobId)
		common.AlarmCronTab.DelByID(jobId)
		logs.Info("monitor task：%v stop", jobId)
	}
	result := map[string]interface{}{"code": "200"}

	c.Data["json"] = result
	c.ServeJSON()
}

// @router /alarm/list/task [get]
func (c *AlarmNoticeMonitorSqlController) ListMonitorTask() {
	var task dto.AlarmNoticeMonitorJobOnOff
	data := c.Ctx.Input.RequestBody

	//json数据封装到user对象中
	err := json.Unmarshal(data, &task)
	if err != nil {
		logs.Info(task)
	}
	ids := common.AlarmCronTab.IDs()

	result := map[string]interface{}{"code": "0"}

	result["data"] = ids
	c.Data["json"] = result
	c.ServeJSON()
}
