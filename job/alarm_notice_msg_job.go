package job

import (
	"alarm/common"
	"alarm/models"
	"alarm/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"os"
	"strconv"
	"strings"
)

func InitAlarmNoticeMsgJob() {
	taskId := "SendAlarmMsg"
	taskFunc := func() {
		o := orm.NewOrm()

		noticeTypes := beego.AppConfig.String("notice.type")
		arrays := strings.Split(noticeTypes, ",")

		for _, noticeType := range arrays {
			messageList := models.GetAlarmNoticeListenerListByNoticeType(o, noticeType)
			for _, msgListener := range messageList {
				alarmMsgList := models.GetAlarmNoticeListenerListByAccessToken(msgListener.AccessToken, 0, o)

				var ids = make([]int64, len(alarmMsgList))
				for idx, alarmMsg := range alarmMsgList {
					sendMessage(alarmMsg)
					ids[idx] = alarmMsg.Id
				}

				models.BatchUpdateAlarmNoticeMsgSuccess(ids, o)
			}
		}
	}

	if err := common.AlarmCronTab.AddByFunc(taskId, "0 0/1 * * * ?", taskFunc); err != nil {
		fmt.Printf("error to add crontab task:%s", err)
		os.Exit(-1)
	}

}

func sendMessage(alarmMsg models.AlarmNoticeMsg) {
	if alarmMsg.NoticeType == "DT" {
		sendDTMessage(alarmMsg)
	} else if alarmMsg.NoticeType == "TG" {
		sendTGMessage(alarmMsg)
	}
}

func sendTGMessage(alarmMsg models.AlarmNoticeMsg) {
	msg := alarmMsg.Msg + " 触发时间:" + alarmMsg.CreateTime.In(common.CstSh).Format("2006-01-02 15:04:05")
	url := "https://api.telegram.org/bot" + alarmMsg.AccessToken + "/sendMessage?chat_id=" + strconv.FormatInt(alarmMsg.ChatId, 10) + "&text=" + msg
	logs.Info("发送TG机器人消息：%s", url)
	go utils.Get(url)
}

func sendDTMessage(alarmMsg models.AlarmNoticeMsg) {
	msg := alarmMsg.Msg + " 触发时间:" + alarmMsg.CreateTime.In(common.CstSh).Format("2006-01-02 15:04:05")

	text := map[string]string{
		"content": msg,
	}
	msgMap := map[string]interface{}{
		"msgtype": "text",
		"text":    text,
	}

	msgJson, _ := json.Marshal(msgMap)
	dingTalkMsg := string(msgJson)

	logs.Info("发送钉钉机器人 %s 消息：%s", alarmMsg.AccessToken, dingTalkMsg)
	go utils.PostJson("https://oapi.dingtalk.com/robot/send?access_token="+alarmMsg.AccessToken, dingTalkMsg, "application/json")
}
