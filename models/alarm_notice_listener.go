package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/prometheus/common/log"
)

type AlarmNoticeListener struct {
	Id          int64
	AccessToken string
	NoticeType  string
	Sn          string
	ChatId      int64
}

func GetAccessTokenByAlarmNoticeTopicSn(sn string, o orm.Ormer) []AlarmNoticeListener {
	if o == nil {
		o = orm.NewOrm()
	}
	var noticeListeners []AlarmNoticeListener

	_, _ = o.Raw("SELECT anl.id,anl.access_token,anl.notice_type,anl.sn,anl.chat_id FROM alarm_notice_listener anl "+
		"INNER JOIN alarm_notice_topic_listener antl ON anl.id = antl.listener_id  "+
		"INNER JOIN alarm_notice_topic ant ON ant.id = antl.topic_id where ant.sn = ?", sn).QueryRows(&noticeListeners)

	return noticeListeners
}

func GetAlarmNoticeListenerListByNoticeType(o orm.Ormer, noticeType interface{}) []AlarmNoticeListener {
	if o == nil {
		o = orm.NewOrm()
	}

	var listeners []AlarmNoticeListener
	qs := o.QueryTable("alarm_notice_listener")
	qs.Filter("notice_type", noticeType).All(&listeners)
	return listeners
}

func GetAlarmNoticeListenerListByToken(o orm.Ormer, accessToken interface{}, chatId int64) []AlarmNoticeListener {
	if o == nil {
		o = orm.NewOrm()
	}

	var listeners []AlarmNoticeListener

	qs := o.QueryTable("alarm_notice_listener")
	if chatId == 0 {
		qs.Filter("access_token", accessToken).All(&listeners)
	} else {
		qs.Filter("access_token", accessToken).Filter("chat_id", chatId).All(&listeners)
	}

	return listeners
}

func GetAllAlarmNoticeListenerList(o orm.Ormer) []AlarmNoticeListener {
	if o == nil {
		o = orm.NewOrm()
	}
	var listeners []AlarmNoticeListener
	qs := o.QueryTable("alarm_notice_listener")
	qs.All(&listeners)
	return listeners
}

func AddAlarmNoticeListener(o orm.Ormer, accessToken string, sn string, noticeType string, chatId int64) AlarmNoticeListener {
	if o == nil {
		o = orm.NewOrm()
	}

	listener := AlarmNoticeListener{AccessToken: accessToken, Sn: sn, NoticeType: noticeType, ChatId: chatId}

	id, err := o.Insert(&listener)
	if err != nil {
		log.Error("保存监听者失败", err)
	}
	listener.Id = id
	return listener
}
