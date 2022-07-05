package controllers

import (
	"alarm/dto"
	"alarm/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AlarmNoticeTopicController struct {
	beego.Controller
}

// @router /alarm/topic/addTopic [post]
func (c *AlarmNoticeTopicController) AddTopic() {
	var topic dto.AlarmNoticeTopic
	data := c.Ctx.Input.RequestBody

	//json数据封装到对象中
	err := json.Unmarshal(data, &topic)

	if err != nil {
		logs.Info(topic)
	}

	result := map[string]interface{}{"code": "200"}

	dto := models.SaveTopic(topic.Remark, nil)
	result["data"] = dto
	c.Data["json"] = result
	c.ServeJSON()

}

// @router /alarm/topic/getList [get]
func (c *AlarmNoticeTopicController) GetAllTopicList() {
	var topic dto.AlarmNoticeTopic
	data := c.Ctx.Input.RequestBody

	//json数据封装到对象中
	err := json.Unmarshal(data, &topic)

	if err != nil {
		logs.Info(topic)
	}

	result := map[string]interface{}{}

	//result := map[string]interface{}{"code": "200"}

	list := models.GetAllTopic(nil, topic.Sn)

	result["code"] = "0"
	result["data"] = list
	c.Data["json"] = result

	c.ServeJSON()
}

// @router /alarm/topic/index [get]
func (c *AlarmNoticeTopicController) AlarmTopicIndex() {

	c.TplName = "alarm/topic/index.tpl"
}
