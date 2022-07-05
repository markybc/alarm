package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/go-basic/uuid"
	"strings"
)

type AlarmNoticeTopic struct {
	Id     int64
	Sn     string
	Remark string
}

func GetAlarmNoticeTopicBySn(sn string, o orm.Ormer) AlarmNoticeTopic {

	var msg AlarmNoticeTopic

	if o == nil {
		o = orm.NewOrm()
	}
	_ = o.QueryTable("alarm_notice_topic").Filter("sn", sn).One(&msg)

	return msg
}

func GetAllTopic(o orm.Ormer, sn interface{}) []AlarmNoticeTopic {

	var topics []AlarmNoticeTopic

	if o == nil {
		o = orm.NewOrm()
	}
	if sn != nil && sn != "" {
		_, _ = o.QueryTable("alarm_notice_topic").Filter("sn", sn).All(&topics)

	} else {
		_, _ = o.QueryTable("alarm_notice_topic").OrderBy("-id").All(&topics)
	}
	return topics
}

func SaveTopic(remark string, o orm.Ormer) AlarmNoticeTopic {
	if o == nil {
		o = orm.NewOrm()
	}

	uuid := uuid.New()
	sn := strings.ReplaceAll(uuid, "-", "")
	topic := AlarmNoticeTopic{Sn: sn, Remark: remark}
	logs.Info("saveTopic start : ", topic)
	id, err := o.Insert(&topic)
	if err != nil {
		logs.Info(err)
	}
	topic.Id = id
	logs.Info("saveTopic end : ", topic)
	return topic
}
