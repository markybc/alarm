package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/prometheus/common/log"
)

type AlarmNoticeTopicListener struct {
	Id         int64
	TopicId    int64
	ListenerId int64
}

func AddAlarmNoticeTopicListener(o orm.Ormer, topicId int64, listenerId int64) AlarmNoticeTopicListener {
	if o == nil {
		o = orm.NewOrm()
	}

	antl := new(AlarmNoticeTopicListener)
	qs := o.QueryTable(antl)
	cnt, _ := qs.Filter("TopicId", topicId).Filter("ListenerId", listenerId).Count()
	if cnt != 0 {
		return AlarmNoticeTopicListener{}
	}
	topicListener := AlarmNoticeTopicListener{TopicId: topicId, ListenerId: listenerId}

	id, err := o.Insert(&topicListener)
	if err != nil {
		log.Error("保存监听者失败", err)
	}
	topicListener.Id = id
	return topicListener
}
