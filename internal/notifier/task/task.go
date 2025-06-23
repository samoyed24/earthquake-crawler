package task

import (
	"earthquake-crawler/internal/config"
	"earthquake-crawler/internal/model"
	"earthquake-crawler/internal/notifier/email"
	"earthquake-crawler/internal/notifier/telegram"
	"fmt"

	"github.com/sirupsen/logrus"
)

var lastEEWOccurTime string = ""

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
	subject := fmt.Sprintf("[日本EEW(%v)]%v发生%v级地震, 预估最大震度%v", data.AlertFlg, data.RegionName, *data.Magnitude, data.CalcIntensity)
	err = email.SendEmail(config.Cfg.Email.EmailReceive.ReceiverUsers, subject, *content)
	if err != nil {
		logrus.Errorf("[Notifier-邮件]在发送日本EEW信息邮件的过程中发生错误: %v", err)
		return
	}
	logrus.Info("[Notifier-邮件]邮件发送成功")
}

func SendJPQuakeTG(data *model.JapanEarthquakeDetail) {
	if !config.Cfg.Telegram.Enable || !config.Cfg.Telegram.Receive.JPQuake.Receive {
		return
	}
	logrus.Info("[Notifier-Telegram]正在尝试向已设置的Telegram信息接收者发送日本地震信息")
	content, err := telegram.RenderJapanEarthquakeTGTemplate(data)
	if err != nil {
		logrus.Errorf("[Notifier-Telegram]在渲染日本地震信息的过程中发生错误: %v", err)
		return
	}
	err = telegram.SendTelegramMessage(content, config.Cfg.Telegram.Receive.ReceiverUsers)
	if err != nil {
		logrus.Errorf("[Notifier-Telegram]在发送日本地震信息Telegram信息的过程中发生错误: %v", err)
		return
	}
	logrus.Info("[Notifier-Telegram]Telegram信息发送成功")
}

func SendJPEEWTG(data *model.JapanEEWData) {
	if !config.Cfg.Telegram.Enable || !config.Cfg.Telegram.Receive.JPEEW.Receive {
		return
	}
	if !config.Cfg.Telegram.Receive.JPEEW.ReceiveTrain && data.IsTraining {
		return
	}
	if config.Cfg.Telegram.Receive.JPEEW.ReceiveAlertOnly && data.AlertFlg != "警報" {
		return
	}
	// 只发送第一报和最终报
	if lastEEWOccurTime == data.OriginTime && !data.IsFinal {
		return
	}
	lastEEWOccurTime = data.OriginTime
	logrus.Info("[Notifier-Telegram]正在尝试向已设置的Telegram信息接收者发送日本EEW信息")
	content, err := telegram.RenderJapanEEWTGTemplate(data)
	if err != nil {
		logrus.Errorf("[Notifier-Telegram]在渲染日本EEW信息的过程中发生错误: %v", err)
		return
	}
	err = telegram.SendTelegramMessage(content, config.Cfg.Telegram.Receive.ReceiverUsers)
	if err != nil {
		logrus.Errorf("[Notifier-Telegram]在发送日本EEW信息Telegram信息的过程中发生错误: %v", err)
		return
	}
	logrus.Info("[Notifier-Telegram]信息发送成功")
}
