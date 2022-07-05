package controllers

import (
	"alarm/models"
	"github.com/astaxie/beego"
)

type AlarmNoticeConnectionController struct {
	beego.Controller
}

// @router /alarm/list/connection [get]
func (c *AlarmNoticeConnectionController) ListConnectionTask() {

	dbList := models.GetAlarmDbConnectionList()

	for index, db := range dbList {
		db.Username = ""
		db.Pwd = ""
		db.HostPort = ""
		dbList[index] = db
	}

	result := map[string]interface{}{"code": "0"}
	result["data"] = dbList
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /alarm/connection/index [get]
func (c *AlarmNoticeTopicController) AlarmConnectionIndex() {

	c.TplName = "alarm/connection/index.tpl"
}
