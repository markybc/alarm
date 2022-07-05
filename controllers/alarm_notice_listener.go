package controllers

import (
	"alarm/dto"
	"alarm/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type AlarmNoticeListenerController struct {
	beego.Controller
}

// @router /alarm/listener/index [get]
func (c *AlarmNoticeTopicController) AlarmListenerIndex() {

	c.TplName = "alarm/listener/index.tpl"
}

// @router /alarm/listener/addListener [post]
func (c *AlarmNoticeListenerController) AddListener() {
	var listener dto.AlarmNoticeListener
	data := c.Ctx.Input.RequestBody

	//json数据封装到对象中
	err := json.Unmarshal(data, &listener)

	if err != nil {
		logs.Info(listener)
	}

	dto := models.AddAlarmNoticeListener(nil, listener.AccessToken, listener.Sn, listener.NoticeType, listener.ChatId)

	result := map[string]interface{}{"code": "200"}

	result["data"] = dto
	c.Data["json"] = result
	c.ServeJSON()

}

// @router /alarm/list/listener [get]
func (c *AlarmNoticeListenerController) ListenerList() {

	listenerList := models.GetAllAlarmNoticeListenerList(nil)

	result := map[string]interface{}{"code": "0"}

	result["data"] = listenerList
	c.Data["json"] = result
	c.ServeJSON()

}

// @router /alarm/listener/addTopicListener [post]
func (c *AlarmNoticeListenerController) AddTopicListener() {
	var topicListener dto.AlarmNoticeTopicListener
	data := c.Ctx.Input.RequestBody

	//json数据封装到对象中
	err := json.Unmarshal(data, &topicListener)

	if err != nil {
		logs.Info(topicListener)
	}

	o := orm.NewOrm()

	topics := models.GetAllTopic(o, topicListener.Topic)

	list := models.GetAllTopic(o, topicListener.Topic)

	if len(topics) > 0 {
		dto := list[0]

		listenerList := models.GetAlarmNoticeListenerListByToken(o, topicListener.AccessToken, topicListener.ChatId)
		if len(listenerList) > 0 {
			for _, listener := range listenerList {
				entity := models.AddAlarmNoticeTopicListener(o, dto.Id, listener.Id)
				logs.Info("保存监听关系:", entity)
			}
		}
	}

	result := map[string]interface{}{"code": "200"}
	c.Data["json"] = result
	c.ServeJSON()

}
