package task

import (
	"earthquake-crawler/internal/config"
	"earthquake-crawler/internal/model"
	"earthquake-crawler/internal/notifier/email"
	"fmt"

	"github.com/sirupsen/logrus"
)

func SendJPQuakeEmail(data *model.JapanEarthquakeDetail) {
	if !config.Cfg.Email.Enable || !config.Cfg.Email.EmailReceive.EmailReceiveJPQuake.Receive {
		return
	}
	logrus.Info("[Notifier-邮件]正在尝试向已设置的邮件接收者发送日本地震信息")
	content, err := email.RenderJapanEarthquakeEmailTemplate(data)
	if err != nil {
		logrus.Errorf("[Notifier-邮件]在渲染日本地震信息的过程中发生错误: %v", err)
		return
	}
	subject := fmt.Sprintf("[日本地震情报更新]%v发生地震", data.Center)
	err = email.SendEmail(config.Cfg.Email.EmailReceive.ReceiverUsers, subject, *content)
	if err != nil {
		logrus.Errorf("[Notifier-邮件]在发送日本地震信息邮件的过程中发生错误: %v", err)
	}
	logrus.Info("[Notifier-邮件]邮件发送成功")
}

var lastEEWOccurTime string = ""

func SendJPEEWEmail(data *model.JapanEEWData) {
	if !config.Cfg.Email.Enable || !config.Cfg.Email.EmailReceive.EmailReceiveJPEEW.Receive {
		return
	}
	if !config.Cfg.Email.EmailReceive.EmailReceiveJPEEW.ReceiveTrain && data.IsTraining {
		return
	}
	if config.Cfg.Email.EmailReceive.EmailReceiveJPEEW.ReceiveAlertOnly && data.AlertFlg != "警報" {
		return
	}
	// 只发送第一报和最终报
	if lastEEWOccurTime == data.OriginTime && !data.IsFinal {
		return
	}
	lastEEWOccurTime = data.OriginTime
	logrus.Info("[Notifier-邮件]正在尝试向已设置的邮件接收者发送日本EEW信息")
	content, err := email.RenderJapanEEWEmailTemplate(data)
	if err != nil {
		logrus.Errorf("[Notifier-邮件]在渲染日本EEW信息的过程中发生错误: %v", err)
		return
	}
	subject := fmt.Sprintf("[日本EEW]%v发生%v级地震", data.RegionName, *data.Magnitude)
	err = email.SendEmail(config.Cfg.Email.EmailReceive.ReceiverUsers, subject, *content)
	if err != nil {
		logrus.Errorf("[Notifier-邮件]在发送日本EEW信息邮件的过程中发生错误: %v", err)
		return
	}
	logrus.Info("[Notifier-邮件]邮件发送成功")
}
